package main

// Estado associado à sessão de um usuário em particular
type Session struct {
	nome   string
	gameId int
}

func InitSession(nome string) *Session {
	return &Session{
		nome:   nome,
		gameId: -1, // começa sem estar em nenhum jogo
	}
}
