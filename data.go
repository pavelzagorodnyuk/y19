package y19

type Data interface {
	Length() int
	Dimension() int
	Scale(p int) Scale
	Value(o int, p int) interface{}
}

type Object interface {
	Dimension() int
	Scale(p int) Scale
	Value(p int) interface{}
}

type Scale int

func (s Scale) String() string {
	switch s {
	case IndefiniteScale:
		return "IndefiniteScale"
	case NominalScale:
		return "NominalScale"
	case AbsoluteScale:
		return "AbsoluteScale"
	default:
		return "unknown"
	}
}

const (
	IndefiniteScale Scale = iota
	NominalScale
	AbsoluteScale
)
