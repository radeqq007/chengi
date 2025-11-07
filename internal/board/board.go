package board

import (
	"chengi/internal/pieces"
)

type Board struct {
	Grid  [8][8]pieces.Piece
}

func New() *Board {
	b := Board{}
	
	// Setup pawns
	for i := 0; i < 8; i++ {
		b.Grid[1][i] = pieces.Piece{Type: pieces.Pawn, Value: pieceValue(pieces.Pawn), Color: pieces.Black}
		b.Grid[6][i] = pieces.Piece{Type: pieces.Pawn, Value: pieceValue(pieces.Pawn), Color: pieces.White}
	}

	setupBackRank := func(row int, color pieces.Color) {
		order := []pieces.PieceType{
			pieces.Rook, pieces.Knight, pieces.Bishop, pieces.Queen,
			pieces.King, pieces.Bishop, pieces.Knight, pieces.Rook,
		}
		for i, t := range order {
			b.Grid[row][i] = pieces.Piece{Type: t, Value: pieceValue(t), Color: color}
		}
	}

	setupBackRank(0, pieces.Black)
	setupBackRank(7, pieces.White)

	return &b
}

type Move struct {
	FromX, FromY int
	ToX, ToY int
	Promotion pieces.PieceType // for pawn promotions
}

func (b *Board) MakeMove(m Move) {
	b.Grid[m.ToX][m.ToY] = b.Grid[m.FromX][m.FromY]
	b.Grid[m.FromX][m.FromY] = pieces.Piece{}
}

func pieceValue(pieceType pieces.PieceType) int {
	return map[pieces.PieceType]int {
		pieces.King: 0,
		pieces.Pawn: 1,
		pieces.Bishop: 3,
		pieces.Knight: 3,
		pieces.Rook: 5,
		pieces.Queen: 9,

	}[pieceType]
}

func (b *Board) GenerateMoves(color pieces.Color) []Move {
	var moves []Move
	for row := range 8 {
		for col := range 8 {
			p := b.Grid[row][col]
			
			if p.Type == pieces.Blank || p.Color != color {
				continue
			}

			switch p.Type {
			case pieces.Pawn:
				moves = append(moves, b.generatePawnMoves(row, col, p)...)
			}
		}
	}
	return moves
}

// TODO: implement en passant
func (b *Board) generatePawnMoves(row, col int, piece pieces.Piece) []Move {
	var moves []Move

	var direction int
	if piece.Color == pieces.Black {
		direction = 1
	} else {
		direction = -1
	}
	
	targetRow := row + direction
	if targetRow >= 0 && targetRow < 8 {
		if !b.isSquareTaken(col, targetRow) {
			moves = append(moves, Move{FromX: col, FromY: row, ToX: col, ToY: targetRow})
		}

		// Taking
		if col < 7 && b.isSquareTaken(col + 1, targetRow) && b.Grid[targetRow][col + 1].Color != piece.Color {
			moves = append(moves, Move{FromX: col, FromY: row, ToX: col + 1, ToY: targetRow})
		}

		if col > 0 && b.isSquareTaken(col - 1, targetRow) && b.Grid[targetRow][col - 1].Color != piece.Color {
			moves = append(moves, Move{FromX: col, FromY: row, ToX: col - 1, ToY: targetRow})
		}
	}

	// First pawn move for black
	if row == 1 && piece.Color == pieces.Black && !b.isSquareTaken(col, row + 1) && !b.isSquareTaken(col, row + 2) {
		moves = append(moves, Move{FromX: col, FromY: row, ToX: col, ToY: row + 2})
	}

	// First pawn move for white
	if row == 6 && piece.Color == pieces.White && !b.isSquareTaken(col, row - 1) && !b.isSquareTaken(col, row - 2) {
		moves = append(moves, Move{FromX: col, FromY: row, ToX: col, ToY: row - 2})
	}
	return moves
}

func (b *Board) isSquareTaken(col, row int) bool {
	return b.Grid[row][col].Type != pieces.Blank
}
