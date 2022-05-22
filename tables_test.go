package y19

import "math"

type generalTable []struct {
	f64 float64
	f32 float32
	s   string
	b   bool
	i   int
	i64 int64
	i32 int32
	i16 int16
	i8  int8
	u   uint
	u64 uint64
	u32 uint32
	u16 uint16
	u8  uint8
}

func (t generalTable) Length() int {
	return len(t)
}

func (t generalTable) Dimension() int {
	return 14
}

func (t generalTable) Scale(a int) Scale {
	switch a {
	case 0, 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		return AbsoluteScale
	case 2, 3:
		return NominalScale
	default:
		return IndefiniteScale
	}
}

func (t generalTable) Value(o, a int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch a {
	case 0:
		return t[o].f64
	case 1:
		return t[o].f32
	case 2:
		return t[o].s
	case 3:
		return t[o].b
	case 4:
		return t[o].i
	case 5:
		return t[o].i64
	case 6:
		return t[o].i32
	case 7:
		return t[o].i16
	case 8:
		return t[o].i8
	case 9:
		return t[o].u
	case 10:
		return t[o].u64
	case 11:
		return t[o].u32
	case 12:
		return t[o].u16
	case 13:
		return t[o].u8
	default:
		return nil
	}
}

type generalTablePnt []struct {
	f64 *float64
	f32 *float32
	s   *string
	b   *bool
	i   *int
	i64 *int64
	i32 *int32
	i16 *int16
	i8  *int8
	u   *uint
	u64 *uint64
	u32 *uint32
	u16 *uint16
	u8  *uint8
}

func (t generalTablePnt) Length() int {
	return len(t)
}

func (t generalTablePnt) Dimension() int {
	return 14
}

func (t generalTablePnt) Scale(a int) Scale {
	switch a {
	case 0, 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		return AbsoluteScale
	case 2, 3:
		return NominalScale
	default:
		return IndefiniteScale
	}
}

func (t generalTablePnt) Value(o, a int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch a {
	case 0:
		return t[o].f64
	case 1:
		return t[o].f32
	case 2:
		return t[o].s
	case 3:
		return t[o].b
	case 4:
		return t[o].i
	case 5:
		return t[o].i64
	case 6:
		return t[o].i32
	case 7:
		return t[o].i16
	case 8:
		return t[o].i8
	case 9:
		return t[o].u
	case 10:
		return t[o].u64
	case 11:
		return t[o].u32
	case 12:
		return t[o].u16
	case 13:
		return t[o].u8
	default:
		return nil
	}
}

type bitOfGeneralTable []struct {
	a float64
	b string
	c bool
	d int
	e int64
	f uint
	g uint64
}

func (t bitOfGeneralTable) Length() int {
	return len(t)
}

func (t bitOfGeneralTable) Dimension() int {
	return 7
}

func (t bitOfGeneralTable) Scale(a int) Scale {
	switch a {
	case 0, 3, 4, 5, 6:
		return AbsoluteScale
	case 1, 2:
		return NominalScale
	default:
		return IndefiniteScale
	}
}

func (t bitOfGeneralTable) Value(o, a int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch a {
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
	case 5:
		return t[o].f
	case 6:
		return t[o].g
	default:
		return nil
	}
}

type tableA []struct {
	a int
	b string
	c string
	d bool
	e int
}

func (t tableA) Length() int {
	return len(t)
}

func (t tableA) Dimension() int {
	return 5
}

func (t tableA) Scale(a int) Scale {
	switch a {
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

func (t tableA) Value(o, a int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch a {
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

type tableB []struct {
	a int
	b string
	c float64
	d float64
	e int
}

func (t tableB) Length() int {
	return len(t)
}

func (t tableB) Dimension() int {
	return 5
}

func (t tableB) Scale(a int) Scale {
	switch a {
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

func (t tableB) Value(o, a int) interface{} {
	if o < 0 || o >= len(t) {
		return nil
	}

	switch a {
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

func (t tableX) Scale(_ int) Scale {
	return IndefiniteScale
}

func (t tableX) Value(_, _ int) interface{} {
	return nil
}

var (
	genTable = generalTable{
		{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
		{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
		{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		{0.03, 13.03, "d", true, -64 + 3, -64 + 3, -32 + 3, -16 + 3, -8 + 3, 64 + 3, 64 + 3, 32 + 3, 16 + 3, 8 + 3},
		{0.04, 14.04, "e", true, -64 + 4, -64 + 4, -32 + 4, -16 + 4, -8 + 4, 64 + 4, 64 + 4, 32 + 4, 16 + 4, 8 + 4},
		{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
	}

	genTablePnt = generalTablePnt{
		{float64Pnt(0.00), float32Pnt(10.00), stringPnt("a"), boolPnt(true), intPnt(-64 + 0), int64Pnt(-64 + 0), int32Pnt(-32 + 0), int16Pnt(-16 + 0), int8Pnt(-8 + 0), uintPnt(64 + 0), uint64Pnt(64 + 0), uint32Pnt(32 + 0), uint16Pnt(16 + 0), uint8Pnt(8 + 0)},
		{float64Pnt(0.01), float32Pnt(11.01), stringPnt("b"), boolPnt(true), intPnt(-64 + 1), int64Pnt(-64 + 1), int32Pnt(-32 + 1), int16Pnt(-16 + 1), int8Pnt(-8 + 1), uintPnt(64 + 1), uint64Pnt(64 + 1), uint32Pnt(32 + 1), uint16Pnt(16 + 1), uint8Pnt(8 + 1)},
		{float64Pnt(0.02), float32Pnt(12.02), stringPnt("c"), boolPnt(true), intPnt(-64 + 2), int64Pnt(-64 + 2), int32Pnt(-32 + 2), int16Pnt(-16 + 2), int8Pnt(-8 + 2), uintPnt(64 + 2), uint64Pnt(64 + 2), uint32Pnt(32 + 2), uint16Pnt(16 + 2), uint8Pnt(8 + 2)},
		{float64Pnt(0.03), float32Pnt(13.03), stringPnt("d"), boolPnt(true), intPnt(-64 + 3), int64Pnt(-64 + 3), int32Pnt(-32 + 3), int16Pnt(-16 + 3), int8Pnt(-8 + 3), uintPnt(64 + 3), uint64Pnt(64 + 3), uint32Pnt(32 + 3), uint16Pnt(16 + 3), uint8Pnt(8 + 3)},
		{float64Pnt(0.04), float32Pnt(14.04), stringPnt("e"), boolPnt(true), intPnt(-64 + 4), int64Pnt(-64 + 4), int32Pnt(-32 + 4), int16Pnt(-16 + 4), int8Pnt(-8 + 4), uintPnt(64 + 4), uint64Pnt(64 + 4), uint32Pnt(32 + 4), uint16Pnt(16 + 4), uint8Pnt(8 + 4)},
		{float64Pnt(0.00), float32Pnt(10.00), stringPnt("a"), boolPnt(true), intPnt(-64 + 0), int64Pnt(-64 + 0), int32Pnt(-32 + 0), int16Pnt(-16 + 0), int8Pnt(-8 + 0), uintPnt(64 + 0), uint64Pnt(64 + 0), uint32Pnt(32 + 0), uint16Pnt(16 + 0), uint8Pnt(8 + 0)},
	}

	genTableWithNil = generalTablePnt{
		{float64Pnt(00.00), float32Pnt(10.00), stringPnt("a"), boolPnt(true), intPnt(-64 + 0), int64Pnt(-64 + 0), int32Pnt(-32 + 0), nil /*         */, int8Pnt(-8 + 0), uintPnt(64 + 0), uint64Pnt(64 + 0), uint32Pnt(32 + 0), uint16Pnt(16 + 0), uint8Pnt(8 + 0)},
		{float64Pnt(00.01), float32Pnt(nan32), nil /*      */, boolPnt(true), intPnt(-64 + 1), int64Pnt(-64 + 1), int32Pnt(-32 + 1), nil /*         */, int8Pnt(-8 + 1), uintPnt(64 + 1), nil /*         */, uint32Pnt(32 + 1), uint16Pnt(16 + 1), uint8Pnt(8 + 1)},
		{nil /*         */, float32Pnt(12.02), stringPnt("c"), boolPnt(true), intPnt(-64 + 2), int64Pnt(-64 + 2), int32Pnt(-32 + 2), int16Pnt(-16 + 2), int8Pnt(-8 + 2), nil /*       */, uint64Pnt(64 + 2), uint32Pnt(32 + 2), uint16Pnt(16 + 2), uint8Pnt(8 + 2)},
		{float64Pnt(00.03), float32Pnt(13.03), stringPnt("d"), boolPnt(true), intPnt(-64 + 3), nil /*         */, int32Pnt(-32 + 3), int16Pnt(-16 + 3), int8Pnt(-8 + 3), uintPnt(64 + 3), uint64Pnt(64 + 3), uint32Pnt(32 + 3), nil /*         */, uint8Pnt(8 + 3)},
		{float64Pnt(nan64), nil /*         */, stringPnt("e"), nil /*     */, intPnt(-64 + 4), nil /*         */, int32Pnt(-32 + 4), int16Pnt(-16 + 4), int8Pnt(-8 + 4), uintPnt(64 + 4), nil /*         */, uint32Pnt(32 + 4), uint16Pnt(16 + 4), uint8Pnt(8 + 4)},
		{float64Pnt(00.00), float32Pnt(10.00), stringPnt("a"), boolPnt(true), intPnt(-64 + 0), nil /*         */, int32Pnt(-32 + 0), int16Pnt(-16 + 0), int8Pnt(-8 + 0), uintPnt(64 + 0), uint64Pnt(64 + 0), uint32Pnt(32 + 0), uint16Pnt(16 + 0), uint8Pnt(8 + 0)},
	}
	nan32 = float32(math.NaN())
	nan64 = math.NaN()

	bitOfGenTable = bitOfGeneralTable{
		{0.00, "a", true, -64 + 0, -64 + 0, 64 + 0, 64 + 0},
		{0.01, "b", true, -64 + 1, -64 + 1, 64 + 1, 64 + 1},
		{0.02, "c", true, -64 + 2, -64 + 2, 64 + 2, 64 + 2},
		{0.03, "d", true, -64 + 3, -64 + 3, 64 + 3, 64 + 3},
		{0.04, "e", true, -64 + 4, -64 + 4, 64 + 4, 64 + 4},
		{0.00, "a", true, -64 + 0, -64 + 0, 64 + 0, 64 + 0},
	}

	mainTableA = tableA{
		{0, "Name0", "Lastname0", false, 10},
		{1, "Name1", "Lastname1", false, 11},
		{2, "Name2", "Lastname2", false, 12},
		{3, "Name3", "Lastname3", false, 13},
		{4, "Name4", "Lastname4", false, 14},
		{0, "Name0", "Lastname0", false, 10},
	}

	mainTableB = tableB{
		{1000, "Title0", 0.0, 1.0, 50},
		{1001, "Title1", 0.1, 1.1, 51},
		{1002, "Title2", 0.2, 1.2, 52},
		{1003, "Title3", 0.3, 1.3, 53},
		{1004, "Title4", 0.4, 1.4, 54},
		{1000, "Title0", 0.0, 1.0, 50},
	}

	mainTableX = tableX{}

	emptyTable = tableA{}
)

func float64Pnt(f float64) *float64 { return &f }
func float32Pnt(f float32) *float32 { return &f }

func stringPnt(s string) *string { return &s }
func boolPnt(b bool) *bool       { return &b }

func intPnt(i int) *int       { return &i }
func int64Pnt(i int64) *int64 { return &i }
func int32Pnt(i int32) *int32 { return &i }
func int16Pnt(i int16) *int16 { return &i }
func int8Pnt(i int8) *int8    { return &i }

func uintPnt(i uint) *uint       { return &i }
func uint64Pnt(i uint64) *uint64 { return &i }
func uint32Pnt(i uint32) *uint32 { return &i }
func uint16Pnt(i uint16) *uint16 { return &i }
func uint8Pnt(i uint8) *uint8    { return &i }
