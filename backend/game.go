package main

import "github.com/corentings/chess/v2"

type GameState struct {
	game     *chess.Game
	players  map[string]string
	turn     string
	madeMove bool
	lastMove ClientMove
}

func InitGameState(whitePlayer, blackPlayer string) *GameState {
	return &GameState{
		game: chess.NewGame(),
		players: map[string]string{
			"w": whitePlayer,
			"b": blackPlayer,
		},
		turn: "w",
		madeMove: false,
	}
}
