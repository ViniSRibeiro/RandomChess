package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all connections
	},
}

var clients = make(map[*websocket.Conn]bool)

type Mensagem struct {
	Msg string `json:"msg"`
}

func (s *Server) chat(w http.ResponseWriter, r *http.Request) {
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
	// Upgrade initial GET request to a websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed ws:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true
	log.Println("New client connected")

	for {
		var msg Mensagem
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Ocorreu um erro na decodificação da mensagem no chat.go\n")
			return
		}

		// Broadcast message to all connected clients
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, jsonChat(msg.Msg, s.sessions[token].nome)); err != nil {
				log.Println("write error ws:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
