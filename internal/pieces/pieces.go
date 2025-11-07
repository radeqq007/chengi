package pieces

type Piece struct {
	Type  PieceType
	Value int
	Color Color
}

type Color int

const (
	Black Color = iota
	White
)

type PieceType int

const (
	Blank PieceType = iota
	Pawn
	Bishop
	Knight
	Rook
	Queen
	King
)

func New(pieceType PieceType, color Color, value int) *Piece {
	return &Piece{
		Type:  pieceType,
		Color: color,
		Value: value,
	}
}
