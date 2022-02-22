package y19

type tableA []struct {
	a  float64
	b  float32
	c  string
	d  bool
	i  int
	i6 int64
	i5 int32
	i4 int16
	i3 int8
	u  uint
	u6 uint64
	u5 uint32
	u4 uint16
	u3 uint8
}

func (t tableA) Length() int {
	return len(t)
}

func (t tableA) Dimension() int {
	return 14
}

func (t tableA) Scale(p int) Scale {
	switch p {
	case 0, 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		return AbsoluteScale
	case 2, 3:
		return NominalScale
	default:
		return IndefiniteScale
	}
}

func (t tableA) Value(o, p int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch p {
	case 0:
		return t[o].a
	case 1:
		return t[o].b
	case 2:
		return t[o].c
	case 3:
		return t[o].d
	case 4:
		return t[o].i
	case 5:
		return t[o].i6
	case 6:
		return t[o].i5
	case 7:
		return t[o].i4
	case 8:
		return t[o].i3
	case 9:
		return t[o].u
	case 10:
		return t[o].u6
	case 11:
		return t[o].u5
	case 12:
		return t[o].u4
	case 13:
		return t[o].u3
	default:
		return nil
	}
}

type tableB []struct {
	a int
	b string
	c string
	d bool
	e int
}

func (t tableB) Length() int {
	return len(t)
}

func (t tableB) Dimension() int {
	return 5
}

func (t tableB) Scale(p int) Scale {
	switch p {
	case 0:
		return AbsoluteScale
	case 1:
		return NominalScale
	case 2:
		return NominalScale
	case 3:
		return NominalScale
	case 4:
		return NominalScale
	default:
		return IndefiniteScale
	}
}

func (t tableB) Value(o, p int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch p {
	case 0:
		return t[o].a
	case 1:
		return t[o].b
	case 2:
		return t[o].c
	case 3:
		return t[o].d
	case 4:
		return t[o].e
	default:
		return nil
	}
}

type tableC []struct {
	a int
	b string
	c float64
	d float64
	e int
}

func (t tableC) Length() int {
	return len(t)
}

func (t tableC) Dimension() int {
	return 5
}

func (t tableC) Scale(p int) Scale {
	switch p {
	case 0:
		return AbsoluteScale
	case 1:
		return NominalScale
	case 2:
		return AbsoluteScale
	case 3:
		return AbsoluteScale
	case 4:
		return NominalScale
	default:
		return IndefiniteScale
	}
}

func (t tableC) Value(o, p int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch p {
	case 0:
		return &t[o].a
	case 1:
		return &t[o].b
	case 2:
		return &t[o].c
	case 3:
		return &t[o].d
	case 4:
		return &t[o].e
	default:
		return nil
	}
}

type tableX struct{}

func (t tableX) Length() int {
	return -1
}

func (t tableX) Dimension() int {
	return -1
}

func (t tableX) Scale(p int) Scale {
	return IndefiniteScale
}

func (t tableX) Value(_, _ int) interface{} {
	return nil
}

var (
	mainTableA = tableA{
		{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
		{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
		{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		{0.03, 13.03, "d", true, -64 + 3, -64 + 3, -32 + 3, -16 + 3, -8 + 3, 64 + 3, 64 + 3, 32 + 3, 16 + 3, 8 + 3},
		{0.04, 14.04, "e", true, -64 + 4, -64 + 4, -32 + 4, -16 + 4, -8 + 4, 64 + 4, 64 + 4, 32 + 4, 16 + 4, 8 + 4},
		{0.05, 15.05, "f", true, -64 + 5, -64 + 5, -32 + 5, -16 + 5, -8 + 5, 64 + 5, 64 + 5, 32 + 5, 16 + 5, 8 + 5},
		{0.06, 16.06, "g", true, -64 + 6, -64 + 6, -32 + 6, -16 + 6, -8 + 6, 64 + 6, 64 + 6, 32 + 6, 16 + 6, 8 + 6},
		{0.07, 17.07, "h", true, -64 + 7, -64 + 7, -32 + 7, -16 + 7, -8 + 7, 64 + 7, 64 + 7, 32 + 7, 16 + 7, 8 + 7},
		{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
	}

	mainTableB = tableB{
		{0, "Name0", "Lastname0", false, 10},
		{1, "Name1", "Lastname1", false, 11},
		{2, "Name2", "Lastname2", false, 12},
		{3, "Name3", "Lastname3", false, 13},
		{4, "Name4", "Lastname4", false, 14},
		{5, "Name5", "Lastname5", false, 15},
		{6, "Name6", "Lastname6", false, 16},
		{7, "Name7", "Lastname7", false, 17},
		{0, "Name0", "Lastname0", false, 10},
	}

	mainTableC = tableC{
		{1000, "Title0", 0.0, 1.0, 50},
		{1001, "Title1", 0.1, 1.1, 51},
		{1002, "Title2", 0.2, 1.2, 52},
		{1003, "Title3", 0.3, 1.3, 53},
		{1004, "Title4", 0.4, 1.4, 54},
		{1005, "Title5", 0.5, 1.5, 55},
		{1006, "Title6", 0.6, 1.6, 56},
		{1007, "Title7", 0.7, 1.7, 57},
		{1000, "Title0", 0.0, 1.0, 50},
	}

	mainTableX = tableX{}

	emptyTable = tableA{}
)
