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


func TestGeneratePawnMoves(t *testing.T) {
	tests := []struct {
		name          string
		setupFunc     func() *Board
		color         pieces.Color
		expectedMoves int
		description   string
	}{
		{
			name: "white pawn single move",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[4][4] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				return New(grid)
			},
			color:         pieces.White,
			expectedMoves: 1,
			description:   "Single forward move available",
		},
		{
			name: "white pawn double move from start",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[6][3] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				return New(grid)
			},
			color:         pieces.White,
			expectedMoves: 2,
			description:   "Single and double move available",
		},
		{
			name: "black pawn double move from start",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[1][5] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				return New(grid)
			},
			color:         pieces.Black,
			expectedMoves: 2,
			description:   "Single and double move available",
		},
		{
			name: "white pawn with captures",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[4][4] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				grid[3][3] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				grid[3][5] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				return New(grid)
			},
			color:         pieces.White,
			expectedMoves: 3,
			description:   "Forward move + 2 captures",
		},
		{
			name: "black pawn with one capture",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[3][4] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				grid[4][5] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				return New(grid)
			},
			color:         pieces.Black,
			expectedMoves: 2,
			description:   "Forward move + 1 capture",
		},
		{
			name: "white pawn blocked",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[4][4] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				grid[3][4] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				return New(grid)
			},
			color:         pieces.White,
			expectedMoves: 0,
			description:   "Blocked by opponent piece",
		},
		{
			name: "white pawn promotion moves",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[1][4] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				return New(grid)
			},
			color:         pieces.White,
			expectedMoves: 4,
			description:   "4 promotion options (Q, R, B, N)",
		},
		{
			name: "white pawn promotion with captures",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[1][4] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				grid[0][3] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				grid[0][5] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				return New(grid)
			},
			color:         pieces.White,
			expectedMoves: 12,
			description:   "4 forward promotions + 4 left capture promotions + 4 right capture promotions",
		},
		{
			name: "black pawn promotion moves",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[6][2] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				return New(grid)
			},
			color:         pieces.Black,
			expectedMoves: 4,
			description:   "4 promotion options (Q, R, B, N)",
		},
		{
			name: "white pawn at edge can't double move if blocked",
			setupFunc: func() *Board {
				grid := [8][8]pieces.Piece{}
				grid[6][0] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.White}
				grid[4][0] = pieces.Piece{Type: pieces.Pawn, Value: 1, Color: pieces.Black}
				return New(grid)
			},
			color:         pieces.White,
			expectedMoves: 1,
			description:   "Can only move one square, double blocked",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.setupFunc()
			moves := b.GenerateMoves(tt.color)
			if len(moves) != tt.expectedMoves {
				t.Errorf("%s: expected %d moves, got %d. Description: %s",
					tt.name, tt.expectedMoves, len(moves), tt.description)
			}
		})
	}
}