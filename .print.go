package y19

import (
	"math"
	"math/rand"
	"reflect"
)

func Fprint(out io.Writer, data Data, max int) (int, error) {

	if out == nil {
		return 0, errors.New("out is nil")
	}

	if data == nil {
		return 0, errors.New("data is nil")
	}

	if max < 0 {
		return 0, nil
	}

	fmt.Fprint(out, "\x1b[1mINDEX")
	for i := 0; i < data.Dimension(); i++ {
		scale := data.Scale(i)
		var h string
		switch scale {
		case NominalScale:
			h = "NS"
		case AbsoluteScale:
			h = "AS"
		case IndefiniteScale:
			h = "IS"
		default:
			h = "IS"
		}
		fmt.Fprintf(out, "%12v [%s]", i, h)
	}
	fmt.Fprintln(out, "\x1b[m")

	if max == 0 || max > data.Length() {
		max = data.Length()
	}

	for i := 0; i < max; i++ {
		fmt.Fprintf(out, "\x1b[1m%d\x1b[m\t", i)
		for j := 0; j < data.Dimension(); j++ {
			fmt.Fprintf(out, "%12.4v", data.Value(i, j))
		}
		fmt.Fprintln(out)
	}
	return 0, nil
}

func Print(data Data, max int) (int, error) {
	return Fprint(os.Stdout, data, max)
}
