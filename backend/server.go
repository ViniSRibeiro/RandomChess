package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/corentings/chess/v2"
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
	games          []*chess.Game
}

func initServer() Server {
	return Server{
		db:         initDB(),
		sessions:   make(map[string]*Session),
		userTokens: make(map[string]string),
		randomness: RD_bitcoin,
		games:      make([]*chess.Game, 0),
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
	valor_antigo := 0.0
	for {
		time.Sleep(1 * time.Second)
		switch s.randomness {
		case RD_bitcoin:
			valor = getBtcData()
		case RD_standart:
			valor = rand.Float64()
		}
		valor = valor - valor_antigo
		valor_antigo = valor
		if err := conn.WriteMessage(websocket.TextMessage, jsonRandom(valor)); err != nil {
			log.Println("write error:", err)
			conn.Close()
		}
	}
}

func (s *Server) esperaJogo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, jsonMsg("Metodo não permitido"), http.StatusMethodNotAllowed)
	}
	// Antes de mais nada, validamos o pedido recebido
	headers := r.Header
	token := headers.Get("Authorization")
	if token != "" {
		http.Error(w, jsonMsg("Faltou o campo Authorization"), http.StatusBadRequest)
		return
	}
	_, validToken := s.sessions[token]
	if !validToken {
		http.Error(w, jsonMsg("Token inválido"), http.StatusBadRequest)
		return
	}

	// Verificamos se já existe alguém na fila. Se já existir, podemos
	// simplesmente montar o novo jogo entre essas duas pessoas
	// Adicionamos o usuário na fila, caso não haja mais ninguém na fila
	// esperando por um jogo
	if len(s.waitingForGame) == 0 {
		s.waitingForGame = append(s.waitingForGame, token)
		w.WriteHeader(http.StatusOK)
		return
	}
	otherToken := s.waitingForGame[0]
	s.waitingForGame = s.waitingForGame[1:]

	gameId := len(s.games)
	s.games = append(s.games, chess.NewGame())

	s.sessions[token].gameId = gameId
	s.sessions[otherToken].gameId = gameId
	w.WriteHeader(http.StatusOK)
}

// -----------------------------------------------------------------------------

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
