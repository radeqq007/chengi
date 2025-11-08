package pieces

type Piece struct {
	Type  PieceType
	Value int
	Color Color
}

type Color uint8

const (
	Black Color = iota
	White
)

type PieceType uint8

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
