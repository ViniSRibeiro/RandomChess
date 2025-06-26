package main

import (
	"log"
	"net/http"
	"time"

	"github.com/corentings/chess/v2"
)

type HttpFunc = func(w http.ResponseWriter, r *http.Request)

type ClientMove struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Promotion string `json:"promotion"`
}

type ServerMove struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Promotion string `json:"promotion"`
	Turn      string `json:"turn"`
	NextTurn  string `json:"nextTurn"`
}

func fromClientMove(c ClientMove, turn string) ServerMove {
	return ServerMove{
		From:      c.From,
		To:        c.To,
		Promotion: c.Promotion,
		Turn:      turn,
		// NOTA isso aí depende de ser determinístico
		NextTurn: getNextTurn(turn),
	}
}

func (s *Server) partida(gameId int) HttpFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("entrou na partida")
		enableCors(&w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Validamos o pedido
		token := getToken(r)
		if token == "" {
			http.Error(w, jsonMsg("Faltou o campo Authorization"), http.StatusBadRequest)
			return
		}
		log.Println(token)
		_, validToken := s.sessions[token]
		if !validToken {
			http.Error(w, jsonMsg("Token inválido"), http.StatusBadRequest)
			return
		}
		// Atualizamos a conexão para websocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Erro no upgrade para websocket em esperaJogo: %v", err)
			return
		}
		defer conn.Close()
		log.Printf("Começou partida id: %d do jogador: %s \n", gameId, s.sessions[token].nome)
		session := s.sessions[token]
		nome := session.nome
		gameState := s.games[gameId]
		for gameState.game.Outcome() == chess.NoOutcome {
			currentPlayer := gameState.players[gameState.turn]
			if token == currentPlayer {
				log.Printf("Vez do jogador: %v\n", nome)
				// O jogador tem que fazer um movimento aqui
				var move ClientMove
				if err := conn.ReadJSON(&move); err != nil {
					log.Printf("Ocorreu um erro na comunicação de partida com %s\n", nome)
					log.Printf("[ ! ] Erro: %v\n", err)
					return
				}
				gameState.sincMove = false
				log.Printf("Chegou Movimento: %v\n", move)
				uci := move.From + move.To + move.Promotion
				// Assumimos que movimento será válido, mas imprimimos o erro mesmo assim
				err := gameState.game.PushNotationMove(uci, chess.UCINotation{},
					&chess.PushMoveOptions{ForceMainline: true})
				if err != nil {
					log.Printf("Movimento inválido: %s", err)
				}
				gameState.madeMove = true
				gameState.lastMove = move
			} else {
				log.Printf("Jogador %v espera o outro \n", nome)
				for !gameState.madeMove {
					// conn.WriteJSON(map[string]string{
					// 	"mensagem": "Aguardando oponente...",
					// })
					time.Sleep(200 * time.Millisecond)
				}
				// Notificamos o trem do movimento feito
				serverMove := fromClientMove(gameState.lastMove, gameState.turn)
				log.Printf("Manda jogada para %v", nome)
				if err := conn.WriteJSON(serverMove); err != nil {
					log.Printf("Ocorreu um erro na comunicação de partida com %s\n", nome)
					return
				}
				gameState.sincMove = true
			}
			for {
				if gameState.madeMove && gameState.sincMove {
					gameState.madeMove = false // oponente fez um movimento
					log.Println("Vira turno")
					gameState.turn = getNextTurn(gameState.turn)
					break
				}
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}

func getNextTurn(turn string) string {
	if turn == "w" {
		return "b"
	}
	return "w"
}
