package board

import (
	"chengi/internal/pieces"
	"testing"
)

func TestNew(t *testing.T) {
	b := New()

	tests := []struct {
        y, x int
        want pieces.Piece
    }{
        // Black pieces
        {0, 0, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Rook), Type: pieces.Rook}},
        {0, 1, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Knight), Type: pieces.Knight}},
        {0, 2, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Bishop), Type: pieces.Bishop}},
        {0, 3, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Queen), Type: pieces.Queen}},
        {0, 4, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.King), Type: pieces.King}},
        {0, 5, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Bishop), Type: pieces.Bishop}},
        {0, 6, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Knight), Type: pieces.Knight}},
        {0, 7, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Rook), Type: pieces.Rook}},

        // Black pawns        
        {1, 0, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {1, 1, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {1, 2, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {1, 3, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {1, 4, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {1, 5, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {1, 6, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {1, 7, pieces.Piece{Color: pieces.Black, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        
        // White pieces
        {7, 0, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Rook), Type: pieces.Rook}},
        {7, 1, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Knight), Type: pieces.Knight}},
        {7, 2, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Bishop), Type: pieces.Bishop}},
        {7, 3, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Queen), Type: pieces.Queen}},
        {7, 4, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.King), Type: pieces.King}},
        {7, 5, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Bishop), Type: pieces.Bishop}},
        {7, 6, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Knight), Type: pieces.Knight}},
        {7, 7, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Rook), Type: pieces.Rook}},
        
        // White pawns        
        {6, 0, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {6, 1, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {6, 2, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {6, 3, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {6, 4, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {6, 5, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {6, 6, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
        {6, 7, pieces.Piece{Color: pieces.White, Value: pieceValue(pieces.Pawn), Type: pieces.Pawn}},
    }
    
    for _, tt := range tests {
        if b.Grid[tt.y][tt.x] != tt.want {
            t.Errorf("Grid[%d][%d] = %v, want %v", tt.y, tt.x, b.Grid[tt.y][tt.x], tt.want)
        }
    }

    for y := 2; y < 6; y++ {
        for x := range(8) {
            if b.Grid[y][x].Type != pieces.Blank {
                t.Errorf("Grid[%d][%d].Type = %v, want %v", y, x, b.Grid[y][x].Type, pieces.Blank)
            }
        }
    }
}
