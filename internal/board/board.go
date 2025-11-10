package board

import (
	"chengi/internal/pieces"
)

type Board struct {
	Grid [8][8]pieces.Piece
}

func New(board ...[8][8]pieces.Piece) *Board {
	b := &Board{}
	if len(board) > 0 {
		b.Grid = board[0]
	} else {
		// Setup pawns
		for i := 0; i < 8; i++ {
			b.Grid[1][i] = pieces.Piece{Type: pieces.Pawn, Value: pieceValue(pieces.Pawn), Color: pieces.Black}
			b.Grid[6][i] = pieces.Piece{Type: pieces.Pawn, Value: pieceValue(pieces.Pawn), Color: pieces.White}
		}

		b.setupBackRank(0, pieces.Black)
		b.setupBackRank(7, pieces.White)
	}

	return b
}

func (b *Board) setupBackRank(row int, color pieces.Color) {
	order := []pieces.PieceType{
		pieces.Rook, pieces.Knight, pieces.Bishop, pieces.Queen,
		pieces.King, pieces.Bishop, pieces.Knight, pieces.Rook,
	}
	for i, t := range order {
		b.Grid[row][i] = pieces.Piece{Type: t, Value: pieceValue(t), Color: color}
	}
}

type Move struct {
	FromCol, FromRow int
	ToCol, ToRow     int
	Promotion        pieces.PieceType // for pawn promotions
}

func (b *Board) MakeMove(m Move) {
	piece := b.Grid[m.FromRow][m.FromCol]
	if m.Promotion != pieces.Blank {
		piece.Type = m.Promotion
		piece.Value = pieceValue(m.Promotion)
	}

	b.Grid[m.ToRow][m.ToCol] = piece
	b.Grid[m.FromRow][m.FromCol] = pieces.Piece{Type: pieces.Blank, Value: int(pieces.Blank)}
}

func pieceValue(pieceType pieces.PieceType) int {
	values := map[pieces.PieceType]int{
		pieces.King:   0,
		pieces.Pawn:   1,
		pieces.Knight: 3,
		pieces.Bishop: 3,
		pieces.Rook:   5,
		pieces.Queen:  9,
	}
	return values[pieceType]
}

func (b *Board) GenerateMoves(color pieces.Color) []Move {
	var moves []Move
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			p := b.Grid[row][col]

			if p.Type == pieces.Blank || p.Color != color {
				continue
			}

			switch p.Type {
			case pieces.Pawn:
				moves = append(moves, b.generatePawnMoves(row, col, p)...)
			case pieces.Knight:
				moves = append(moves, b.generateKnightMoves(row, col, p)...)
			case pieces.Bishop:
				moves = append(moves, b.generateBishopMoves(row, col, p)...)
			case pieces.Rook:
				moves = append(moves, b.generateRookMoves(row, col, p)...)
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
	if b.isInBounds(targetRow) {
		if !b.isSquareTaken(targetRow, col) {
			if targetRow == 0 || targetRow == 7 {
				moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col, ToRow: row + direction, Promotion: pieces.Queen})
				moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col, ToRow: row + direction, Promotion: pieces.Rook})
				moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col, ToRow: row + direction, Promotion: pieces.Bishop})
				moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col, ToRow: row + direction, Promotion: pieces.Knight})
			} else {
				moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col, ToRow: targetRow})
			}
		}

		// Taking (right)
		if col < 7 && b.isSquareTaken(targetRow, col+1) && b.Grid[targetRow][col+1].Color != piece.Color {
			if targetRow == 0 || targetRow == 7 {
				moves = append(moves,
					Move{FromCol: col, FromRow: row, ToCol: col + 1, ToRow: targetRow, Promotion: pieces.Queen},
					Move{FromCol: col, FromRow: row, ToCol: col + 1, ToRow: targetRow, Promotion: pieces.Rook},
					Move{FromCol: col, FromRow: row, ToCol: col + 1, ToRow: targetRow, Promotion: pieces.Bishop},
					Move{FromCol: col, FromRow: row, ToCol: col + 1, ToRow: targetRow, Promotion: pieces.Knight},
				)
			} else {
				moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col + 1, ToRow: targetRow})
			}
		}

		// Taking (left)
		if col > 0 && b.isSquareTaken(targetRow, col-1) && b.Grid[targetRow][col-1].Color != piece.Color {
			if targetRow == 0 || targetRow == 7 {
				moves = append(moves,
					Move{FromCol: col, FromRow: row, ToCol: col - 1, ToRow: targetRow, Promotion: pieces.Queen},
					Move{FromCol: col, FromRow: row, ToCol: col - 1, ToRow: targetRow, Promotion: pieces.Rook},
					Move{FromCol: col, FromRow: row, ToCol: col - 1, ToRow: targetRow, Promotion: pieces.Bishop},
					Move{FromCol: col, FromRow: row, ToCol: col - 1, ToRow: targetRow, Promotion: pieces.Knight},
				)
			} else {
				moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col - 1, ToRow: targetRow})
			}
		}
	}

	// First pawn move for black
	if row == 1 && piece.Color == pieces.Black && !b.isSquareTaken(row+1, col) && !b.isSquareTaken(row+2, col) {
		moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col, ToRow: row + 2})
	}

	// First pawn move for white
	if row == 6 && piece.Color == pieces.White && !b.isSquareTaken(row-1, col) && !b.isSquareTaken(row-2, col) {
		moves = append(moves, Move{FromCol: col, FromRow: row, ToCol: col, ToRow: row - 2})
	}

	return moves
}

func (b *Board) generateKnightMoves(row, col int, piece pieces.Piece) []Move {
	var moves []Move
	var offsets = [8][2]int{
		{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2},
		{1, -2}, {1, 2}, {2, -1}, {2, 1},
	}

	for _, offset := range offsets {
		newRow := row + offset[0]
		newCol := col + offset[1]

		if b.isInBounds(newRow) && b.isInBounds(newCol) {
			target := b.Grid[newRow][newCol]
			if target.Type == pieces.Blank || target.Color != piece.Color {
				moves = append(moves, Move{
					FromRow: row,
					FromCol: col,
					ToRow:   newRow,
					ToCol:   newCol,
				})
			}
		}
	}

	return moves
}
func (b *Board) generateBishopMoves(row, col int, piece pieces.Piece) []Move {
	var moves []Move
	moves = append(moves, b.generateSlidingMoves(row, col, 1, 1, piece)...)
	moves = append(moves, b.generateSlidingMoves(row, col, -1, -1, piece)...)
	moves = append(moves, b.generateSlidingMoves(row, col, 1, -1, piece)...)
	moves = append(moves, b.generateSlidingMoves(row, col, -1, 1, piece)...)

	return moves
}

func (b *Board) generateRookMoves(row, col int, piece pieces.Piece) []Move {
	var moves []Move
	moves = append(moves, b.generateSlidingMoves(row, col, 1, 0, piece)...)
	moves = append(moves, b.generateSlidingMoves(row, col, -1, 0, piece)...)
	moves = append(moves, b.generateSlidingMoves(row, col, 0, 1, piece)...)
	moves = append(moves, b.generateSlidingMoves(row, col, 0, -1, piece)...)

	return moves
}

func (b *Board) generateSlidingMoves(row, col, rowDir, colDir int, piece pieces.Piece) []Move {
	var moves []Move
	newRow := row + rowDir
	newCol := col + colDir
	for b.isInBounds(newRow) && b.isInBounds(newCol) {
		target := b.Grid[newRow][newCol]
		if target.Type == pieces.Blank {
			moves = append(moves, Move{
				FromRow: row,
				FromCol: col,
				ToRow:   newRow,
				ToCol:   newCol,
			})
		} else {
			if target.Color != piece.Color {
				moves = append(moves, Move{
					FromRow: row,
					FromCol: col,
					ToRow:   newRow,
					ToCol:   newCol,
				})
			}
			break
		}
		newRow += rowDir
		newCol += colDir
	}
	return moves
}

func (b *Board) isSquareTaken(row, col int) bool {
	return b.Grid[row][col].Type != pieces.Blank
}

func (b *Board) isInBounds(index int) bool {
	return index >= 0 && index < 8
}
