package main

// Estado associado à sessão de um usuário em particular
type Session struct {
	nomeUsuario string
	gameId      int
}

func InitSession(nome string) Session {
	return Session{
		nomeUsuario: nome,
		gameId:  -1, // começa sem estar em nenhum jogo
	}
}
