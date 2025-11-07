package board

import "chengi/internal/pieces"

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

func (b *Board) GenerateMoves(color pieces.Color) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			p := b.Grid[x][y]
			
			if p.Type == pieces.Blank || p.Color != color {
				continue
			}

			switch p.Type {
			case pieces.Pawn:

			}
		}
	}
}

// func generatePawnMoves(x, y int, ) {
//
// }