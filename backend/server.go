package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

type randomness = uint8

const (
	RD_bitcoin randomness = iota
	RD_standart
)

type Server struct {
	db             *sql.DB
	sessions       map[string]*Session // chaves são tokens
	userTokens     map[string]string   // chaves são nomes de usuário
	waitingForGame []string            // fila de usuários aguardando jogo
	randomness     randomness
	games          []*GameState
}

func initServer() Server {
	return Server{
		db:         initDB(),
		sessions:   make(map[string]*Session),
		userTokens: make(map[string]string),
		randomness: RD_standart,
		games:      make([]*GameState, 0),
	}
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Println("Não foi possível abrir o banco app.db")
		log.Fatal(err)
		return nil
	}
	// Inicializamos as tabelas por meio da execução do arquivo banco.sql
	contents, err := os.ReadFile("./banco.sql")
	if err != nil {
		log.Println("Não possível executar o script banco.sql")
		log.Fatal(err)
		return nil
	}
	sql := string(contents)
	if _, err := db.Exec(sql); err != nil {
		log.Println("Não possível executar o script banco.sql")
		log.Fatal(err)
		return nil
	}
	return db
}

// -----------------------------------------------------------------------------

func (s *Server) random(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, jsonMsg("Metodo não permitido"), http.StatusMethodNotAllowed)
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	var valor float64
	switch s.randomness {
	case RD_bitcoin:
		valor = getBtcData()
	case RD_standart:
		valor = rand.Float64()
	}

	var variacao float64
	valor_antigo := valor
	for {
		time.Sleep(1 * time.Second)
		switch s.randomness {
		case RD_bitcoin:
			valor = getBtcData()
		case RD_standart:
			valor = (rand.Float64() - 0.5) * 10
		}
		variacao = valor - valor_antigo
		valor_antigo = valor
		if err := conn.WriteMessage(websocket.TextMessage, jsonRandom(variacao)); err != nil {
			conn.Close()
		}
	}
}

func (s *Server) esperaJogo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("Pedido de conexão chegou")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, jsonMsg("Metodo não permitido"), http.StatusMethodNotAllowed)
	}
	// Antes de mais nada, validamos o pedido recebido
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
	// Se não há ninguém na fila, esperamos
	if slices.Contains(s.waitingForGame, token) {
		return
	}
	if len(s.waitingForGame) == 0 {
		s.waitingForGame = append(s.waitingForGame, token)
		// Enquanto não aparece outra pessoa, comunicamos que não há ninguém
		for s.sessions[token].gameId < 0 {
			conn.WriteJSON(map[string]string{
				"encontrou": "N",
				"mensagem":  "Aguardando oponente...",
			})
			log.Printf("Aguardando para o token %s\n", token)
			time.Sleep(2 * time.Second)
		}
		conn.WriteJSON(map[string]string{
			"encontrou": "S",
			"partida":   fmt.Sprint(s.sessions[token].gameId),
			"color":     "white", // podia ser sorteado. Que pena!
		})
		return
	}
	log.Printf("fila de espera antes: %v", s.waitingForGame)
	otherToken := s.waitingForGame[0]
	// clear(s.waitingForGame)
	s.waitingForGame = make([]string, 0)
	log.Printf("fila de espera depois: %v", s.waitingForGame)

	gameId := len(s.games)
	s.games = append(s.games, InitGameState(otherToken, token))
	// Registramos uma nova rota para a nova partida
	routeName := fmt.Sprintf("/partida/%d", gameId)
	http.HandleFunc(routeName, s.partida(gameId))

	s.sessions[token].gameId = gameId
	s.sessions[otherToken].gameId = gameId

	log.Printf("sessões: %v", s.sessions)

	conn.WriteJSON(map[string]string{
		"encontrou": "S",
		"partida":   fmt.Sprint(s.sessions[token].gameId),
		"color":     "black", // podia ser sorteado
	})
	log.Printf("Criada a partida %d com os usuários de token %s e %s",
		gameId, s.sessions[token].nome, s.sessions[otherToken].nome)

}

// -----------------------------------------------------------------------------

func getToken(r *http.Request) string {
	contents, hasAuth := r.Header["Authorization"]
	if hasAuth {
		return contents[0]
	}
	contents, hasAuth = r.Header["Sec-Websocket-Protocol"]
	if hasAuth {
		return contents[0]
	}
	return ""
}

func ok(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func error(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.Error(w, jsonMsg("Zorza Cabeça de melão"), http.StatusBadRequest)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-POST, GET, OPTIONS", "*")
}

func jsonMsg(msg string) string {
	data := map[string]string{"mensagem": msg}
	res, _ := json.Marshal(data)
	return string(res)
}

func jsonToken(msg string) []byte {
	data := map[string]string{"token": msg}
	res, _ := json.Marshal(data)
	return res
}

func jsonRandom(num float64) []byte {
	data := map[string]float64{"valor": num}
	res, _ := json.Marshal(data)
	return res
}

func jsonChat(msg string, user string) []byte {
	data := map[string]string{
		"mensagem": msg,
		"usuario":  user,
	}
	res, _ := json.Marshal(data)
	return res
}

func getBtcData() float64 {
	res, err := http.Get("http://cointradermonitor.com/api/pbb/v1/ticker")
	if err != nil {
		log.Fatalf("[!] Erro ao fazer a requisição HTTP: %s", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Erro ao ler o corpo da resposta: %s", err)
	}

	var tickerData ticker
	if err := json.Unmarshal(body, &tickerData); err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %s", err)
	}
	return tickerData.Last
}
