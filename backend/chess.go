package main

import "unicode"

// Define a representação utilizada no backend para o tabuleiro de xadrez de
// cada partida. O modelo inclui a noção de quais movimentos são legais para
// uma peça específica.

type pieceKind = int

const (
	PIECE_none   pieceKind = iota
	PIECE_pawn             // peão
	PIECE_rook             // torre
	PIECE_knight           // cavalo
	PIECE_bishop           // bispo
	PIECE_queen            // rainha
	PIECE_king             // rei
)

type pieceColor = uint8

const (
	COLOR_black pieceColor = iota
	COLOR_white
)

type Piece struct {
	kind  pieceKind
	color pieceColor
}

type Square struct {
	col, row int
}

// Converte um código alfanumérico em uma instância de square
func nameToSquare(squareName []byte) Square {
	col := int(unicode.ToLower(rune(squareName[0]))) - 'a'
	row := int(squareName[1] - '1')
	return Square{col: col, row: row}
}

// Converte uma instância de square em um código alfanumérico
func (s Square) String() string {
	var name [2]byte
	name[0] = byte(s.col) + 'a'
	name[1] = byte(s.col) + '1'
	return string(name[:])
}

// -----------------------------------------------------------------------------

type ChessBoard struct {
	board [8][8]Piece
}

func InitChessBoard() *ChessBoard {
	c := &ChessBoard{}
	// Peças brancas (não peões) na primeira linha
	c.board[0][0] = Piece{kind: PIECE_rook, color: COLOR_white}
	c.board[0][1] = Piece{kind: PIECE_knight, color: COLOR_white}
	c.board[0][2] = Piece{kind: PIECE_bishop, color: COLOR_white}
	c.board[0][3] = Piece{kind: PIECE_queen, color: COLOR_white}
	c.board[0][4] = Piece{kind: PIECE_king, color: COLOR_white}
	c.board[0][5] = Piece{kind: PIECE_bishop, color: COLOR_white}
	c.board[0][6] = Piece{kind: PIECE_knight, color: COLOR_white}
	c.board[0][7] = Piece{kind: PIECE_rook, color: COLOR_white}
	// Peças brancas (peões) na segunda linha
	for i := 0; i < 8; i++ {
		c.board[1][i] = Piece{kind: PIECE_pawn, color: COLOR_white}
	}

	// Peças pretas (não peões) na última linha
	c.board[7][0] = Piece{kind: PIECE_rook, color: COLOR_black}
	c.board[7][1] = Piece{kind: PIECE_knight, color: COLOR_black}
	c.board[7][2] = Piece{kind: PIECE_bishop, color: COLOR_black}
	c.board[7][3] = Piece{kind: PIECE_queen, color: COLOR_black}
	c.board[7][4] = Piece{kind: PIECE_king, color: COLOR_black}
	c.board[7][5] = Piece{kind: PIECE_bishop, color: COLOR_black}
	c.board[7][6] = Piece{kind: PIECE_knight, color: COLOR_black}
	c.board[7][7] = Piece{kind: PIECE_rook, color: COLOR_black}
	// Peças pretas (peões) na penúltima linha
	for i := 0; i < 8; i++ {
		c.board[6][i] = Piece{kind: PIECE_pawn, color: COLOR_black}
	}
	// As demais casas começam vazias
	for i := 2; i < 6; i++ {
		for j := 0; j < 8; j++ {
			c.board[i][j] = Piece{kind: PIECE_none}
		}
	}
	return c
}

func filterValidSquares(squares []Square) []Square {
	valid := make([]Square, 0)
	for _, pos := range squares {
		if pos.row >= 0 && pos.col >= 0 && pos.row < 8 && pos.col < 8 {
			continue // fora do tabuleiro
		}
		valid = append(valid, pos)
	}
	return valid
}

func (c *ChessBoard) validMovesFrom(s Square) []Square {
	moves := make([]Square, 0)
	piece := c.board[s.row][s.col]
	switch piece.kind {
	case PIECE_none:
		break
	case PIECE_pawn:
		moves = append(moves, Square{row: s.row + 1, col: s.col})
	case PIECE_rook:
		for i := s.row; i >= 0; i-- {
			moves = append(moves, Square{row: i, col: s.col})
			if c.board[i][s.col].kind != PIECE_none {
				break // só podemos ir até aí
			}
		}
		for i := s.row + 1; i < 8; i++ {
			moves = append(moves, Square{row: i, col: s.col})
			if c.board[i][s.col].kind != PIECE_none {
				break // só podemos ir até aí
			}
		}
		for j := s.col + 1; j < 8; j++ {
			moves = append(moves, Square{row: s.row, col: j})
			if c.board[s.row][j].kind != PIECE_none {
				break // só podemos ir até aí
			}
		}
		for j := s.col; j >= 0; j-- {
			moves = append(moves, Square{row: s.row, col: j})
			if c.board[s.row][j].kind != PIECE_none {
				break // só podemos ir até aí
			}
		}

	// Pelo código acima já dá pra ter uma noção de quão emocionante esse
	// código vai ser. Maldito o dia em que decidimos fazer um jogo de xadrez
	case PIECE_knight:
		deltas := [...]int{ -2, -1, 1, 2 }
		for _, delta1 := range deltas {
			for _, delta2 := range deltas {
				moves = append(moves, Square{row: s.row + delta1, col: s.col + delta2})
			}
		}

	case PIECE_bishop:
	case PIECE_king:
		deltas := [...]int{ -1, 1 }
		for _, delta1 := range deltas {
			for _, delta2 := range deltas {
				moves = append(moves, Square{row: s.row + delta1, col: s.col + delta2})
			}
		}
	}

	return filterValidSquares(moves)
}
