package y19

import (
	"reflect"
	"testing"
)

type nodeTestCase struct {
	name         string
	node, sample Data
}

func nodeTesting(t *testing.T, test nodeTestCase, nodeName string) {

	if test.node.Length() != test.sample.Length() {
		t.Fatalf("%s.Length() is not equal to the expected length\n"+
			"\tExpected: %d\n\tGot: %d\n", nodeName, test.sample.Length(), test.node.Length())
	}

	if test.node.Dimension() != test.sample.Dimension() {
		t.Fatalf("%s.Dimension() is not equal to the expected dimension\n"+
			"\tExpected: %d\n\tGot: %d\n", nodeName, test.sample.Dimension(), test.node.Dimension())
	}

	if test.node.Scale(-1) != IndefiniteScale {
		t.Errorf("%s.Scale(-1) must be equal to IndefiniteScale\n\tGot: %v",
			nodeName, test.node.Scale(-1))
	}

	if test.node.Scale(test.node.Dimension()) != IndefiniteScale {
		t.Errorf("%s.Scale(%s.Dimension()) must be equal to IndefiniteScale\n\tGot: %v",
			nodeName, nodeName, test.node.Scale(test.node.Dimension()))
	}

	for a := 0; a < test.sample.Dimension(); a++ {
		if test.node.Scale(a) != test.sample.Scale(a) {
			t.Errorf("%s.Scale(%d) is not equal to the expected scale\n"+
				"\tExpected: %v\n\tGot: %v\n", nodeName, a, test.sample.Scale(a), test.node.Scale(a))
			break
		}
	}

	for i := -3; i < test.sample.Length()+3; i++ {
		for j := -3; j < test.sample.Dimension()+3; j++ {
			if i < 0 || i >= test.sample.Length() || j < 0 || j >= test.sample.Dimension() {
				if test.node.Value(i, j) != nil {
					t.Fatalf("%s.Value(%d, %d) must be equal to nil", nodeName, i, j)
				}
			} else if !reflect.DeepEqual(test.node.Value(i, j), test.sample.Value(i, j)) {
				t.Fatalf("%s.Value(%d, %d) is not equal to the expected value\n"+
					"\tExpected: %v\n\tGot: %v\n", nodeName, i, j, test.sample.Value(i, j), test.node.Value(i, j))
			}
		}
	}
}

func nodeBenchmarking(b *testing.B, node Data) {
	b.Run("Length", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			node.Length()
		}
	})

	b.Run("Dimension", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			node.Dimension()
		}
	})

	b.Run("Scale", func(b *testing.B) {
		var dimension int = node.Dimension()
		for i := 0; i < b.N; i++ {
			node.Scale(i % dimension)
		}
	})

	b.Run("Value", func(b *testing.B) {
		var dimension int = node.Dimension()
		var length int = node.Length()
		for i := 0; i < b.N; i++ {
			node.Value((i/dimension)%length, i%dimension)
		}
	})
}

type areEqualInput struct {
	selectionOne, selectionTwo Data
}

var areEqualTestCases = []struct {
	name           string
	input          areEqualInput
	expectedOutput bool
}{
	{
		name:           "IDENTICAL_TABLES",
		input:          areEqualInput{selectionOne: genTable, selectionTwo: genTable},
		expectedOutput: true,
	},
	{
		name:           "TABLE_AND_IDENTICALPOINTER_TABLE",
		input:          areEqualInput{selectionOne: genTable, selectionTwo: genTablePnt},
		expectedOutput: true,
	},
	{
		name:           "POINTER_TABLE_WITH_ITSELF",
		input:          areEqualInput{selectionOne: genTablePnt, selectionTwo: genTablePnt},
		expectedOutput: true,
	},
	{
		name:           "TABLE_WITH_NIL_WITH_ITSELF",
		input:          areEqualInput{selectionOne: genTableWithNil, selectionTwo: genTableWithNil},
		expectedOutput: true,
	},
	{
		name:           "EMPTY_TABLE_WITH_ITSELF",
		input:          areEqualInput{selectionOne: emptyTable, selectionTwo: emptyTable},
		expectedOutput: true,
	},
	{
		name:           "TABLE_X_WITH_ITSELF",
		input:          areEqualInput{selectionOne: mainTableX, selectionTwo: mainTableX},
		expectedOutput: false,
	},
	{
		name:           "TWO_NIL",
		input:          areEqualInput{selectionOne: nil, selectionTwo: nil},
		expectedOutput: false,
	},
	{
		name:           "NIL_AND_TABLE",
		input:          areEqualInput{selectionOne: nil, selectionTwo: genTable},
		expectedOutput: false,
	},
	{
		name:           "TABLE_AND_NIL",
		input:          areEqualInput{selectionOne: genTable, selectionTwo: nil},
		expectedOutput: false,
	},
	{
		name:           "TABLES_WITH_DIFFERENT_HEADER_LENGTHS",
		input:          areEqualInput{selectionOne: genTable, selectionTwo: mainTableA},
		expectedOutput: false,
	},
	{
		name:           "TABLES_WITH_DIFFERENT_HEADERS",
		input:          areEqualInput{selectionOne: mainTableA, selectionTwo: mainTableB},
		expectedOutput: false,
	},
	{
		name: "DIFFERENT_TABLES_WITH_IDENTICAL_HEADERS",
		input: areEqualInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			selectionTwo: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 /**/},
			},
		},
		expectedOutput: false,
	},
	{
		name: "TABLE_AND_IDENTICAL_MIXED_TABLE",
		input: areEqualInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			selectionTwo: generalTable{
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			},
		},
		expectedOutput: true,
	},
	{
		name: "SIMILAR_TABLES_WITH_DIFFERENT_IDENTICAL_ROWS",
		input: areEqualInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			},
			selectionTwo: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			},
		},
		expectedOutput: false,
	},
}

func TestAreEqual(t *testing.T) {
	for _, test := range areEqualTestCases {
		t.Run(test.name, func(t *testing.T) {
			output := AreEqual(test.input.selectionOne, test.input.selectionTwo)
			if test.expectedOutput != output {
				t.Errorf("the received result differs from the expected one\n\tExpected: %t\n\tGot: %t", test.expectedOutput, output)
			}
		})
	}
}

func BenchmarkAreEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreEqual(genTable, genTable)
	}
}

type areSimilarInput struct {
	selectionOne, selectionTwo Data
	deviation                  float64
}

var areSimilarTests = []struct {
	name           string
	input          areSimilarInput
	expectedOutput bool
}{
	{
		name:           "IDENTICAL_TABLES",
		input:          areSimilarInput{selectionOne: genTable, selectionTwo: genTable},
		expectedOutput: true,
	},
	{
		name:           "TABLE_AND_IDENTICALPOINTER_TABLE",
		input:          areSimilarInput{selectionOne: genTable, selectionTwo: genTablePnt},
		expectedOutput: true,
	},
	{
		name:           "POINTER_TABLE_WITH_ITSELF",
		input:          areSimilarInput{selectionOne: genTablePnt, selectionTwo: genTablePnt},
		expectedOutput: true,
	},
	{
		name:           "TABLE_WITH_NIL_WITH_ITSELF",
		input:          areSimilarInput{selectionOne: genTableWithNil, selectionTwo: genTableWithNil},
		expectedOutput: true,
	},
	{
		name:           "EMPTY_TABLE_WITH_ITSELF",
		input:          areSimilarInput{selectionOne: emptyTable, selectionTwo: emptyTable},
		expectedOutput: true,
	},
	{
		name:           "TABLE_X_WITH_ITSELF",
		input:          areSimilarInput{selectionOne: mainTableX, selectionTwo: mainTableX},
		expectedOutput: false,
	},
	{
		name:           "TWO_NIL",
		input:          areSimilarInput{selectionOne: nil, selectionTwo: nil},
		expectedOutput: false,
	},
	{
		name:           "NIL_AND_TABLE",
		input:          areSimilarInput{selectionOne: nil, selectionTwo: genTable},
		expectedOutput: false,
	},
	{
		name:           "TABLE_AND_NIL",
		input:          areSimilarInput{selectionOne: genTable, selectionTwo: nil},
		expectedOutput: false,
	},
	{
		name:           "TABLES_WITH_DIFFERENT_HEADER_LENGTHS",
		input:          areSimilarInput{selectionOne: genTable, selectionTwo: mainTableA},
		expectedOutput: false,
	},
	{
		name:           "TABLES_WITH_DIFFERENT_HEADERS",
		input:          areSimilarInput{selectionOne: mainTableA, selectionTwo: mainTableB},
		expectedOutput: false,
	},
	{
		name: "DIFFERENT_TABLES_WITH_IDENTICAL_HEADERS",
		input: areSimilarInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			selectionTwo: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "e" /**/, true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
		},
		expectedOutput: false,
	},
	{
		name: "TABLE_AND_IDENTICAL_MIXED_TABLE",
		input: areSimilarInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			selectionTwo: generalTable{
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			},
		},
		expectedOutput: true,
	},
	{
		name: "TABLES_WITH_ACCEPTABLE_DEVIATION",
		input: areSimilarInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			selectionTwo: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{-0.03 /**/, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			deviation: 0.05,
		},
		expectedOutput: true,
	},
	{
		name: "TABLES_WITH_UNACCEPTABLE_DEVIATION",
		input: areSimilarInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			selectionTwo: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.071 /**/, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			deviation: 0.05,
		},
		expectedOutput: false,
	},
	{
		name: "TABLES_WITH_BOUNDARY_DEVIATION",
		input: areSimilarInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			},
			selectionTwo: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 /**/},
			},
			deviation: 2.,
		},
		expectedOutput: true,
	},
	{
		name: "SIMILAR_TABLES_WITH_DIFFERENT_IDENTICAL_ROWS",
		input: areSimilarInput{
			selectionOne: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			},
			selectionTwo: generalTable{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			},
			deviation: 0.05,
		},
		expectedOutput: false,
	},
}

func TestAreSimilar(t *testing.T) {
	for _, test := range areSimilarTests {
		output := AreSimilar(test.input.selectionOne, test.input.selectionTwo, test.input.deviation)
		if test.expectedOutput != output {
			t.Errorf("the received result differs from the expected one\n\tExpected: %t\n\tGot: %t", test.expectedOutput, output)
		}
	}
}

func BenchmarkAreSimilar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreSimilar(genTable, genTable, 0)
	}
}

var combineNodeTestCases = []nodeTestCase{
	{
		name:   "CASE_WITH_MANY_SELECTIONS",
		sample: genTable,
		node: &combineNode{
			header:        []Scale{AbsoluteScale, NominalScale, NominalScale, NominalScale, NominalScale},
			generalLength: 6,
			startPointers: []int{2, 2, 4},
			sources: []Data{
				genTable[:2],
				tableA{},
				genTable[2:4],
				genTable[4:],
			},
		},
	},
	{
		name:   "CASE_WITH_SINGLE_SELECTION",
		sample: genTable,
		node: &combineNode{
			header:        []Scale{AbsoluteScale, NominalScale, NominalScale, NominalScale, NominalScale},
			generalLength: 2,
			startPointers: nil,
			sources:       []Data{genTable},
		},
	},
}

func TestCombineNode(t *testing.T) {
	for _, test := range combineNodeTestCases {
		t.Run(test.name, func(t *testing.T) {
			nodeTesting(t, test, "CombineNode")
		})
	}
}

func BenchmarkCombineNode(b *testing.B) {
	node := Combine(genTable, genTable)
	nodeBenchmarking(b, node)
}

var combineTestCases = []struct {
	name           string
	input          []Data
	expectedOutput Data
}{
	{
		name:  "LOTS_OF_TABLES",
		input: []Data{mainTableA[:2], mainTableA[2:4], mainTableA[4:]},
		expectedOutput: &combineNode{
			header:        []Scale{AbsoluteScale, NominalScale, NominalScale, NominalScale, NominalScale},
			generalLength: len(mainTableA),
			startPointers: []int{2, 4},
			sources:       []Data{mainTableA[:2], mainTableA[2:4], mainTableA[4:]},
		},
	},
	{
		name:  "EMPTY_TABLES",
		input: []Data{emptyTable, emptyTable},
		expectedOutput: &combineNode{
			header:        []Scale{AbsoluteScale, NominalScale, NominalScale, NominalScale, NominalScale},
			generalLength: 0,
			startPointers: []int{0},
			sources:       []Data{emptyTable, emptyTable},
		},
	},
	{
		name:           "TWO_TABLE_X",
		input:          []Data{mainTableX, mainTableX},
		expectedOutput: nil,
	},
	{
		name:           "TABLE_AND_NIL",
		input:          []Data{genTable, nil},
		expectedOutput: nil,
	},

	{
		name:           "TABLES_WITH_DIFFERENT_HEADER_LENGTHS",
		input:          []Data{genTable, mainTableA},
		expectedOutput: nil,
	},
	{
		name:           "TABLES_WITH_DIFFERENT_HEADERS",
		input:          []Data{mainTableA, mainTableB},
		expectedOutput: nil,
	},
	{
		name:  "TABLES_WITH_EMPTY_TABLE",
		input: []Data{mainTableA[:3], emptyTable, mainTableA[3:]},
		expectedOutput: &combineNode{
			header:        []Scale{AbsoluteScale, NominalScale, NominalScale, NominalScale, NominalScale},
			generalLength: len(genTable),
			startPointers: []int{3, 3},
			sources:       []Data{mainTableA[:3], emptyTable, mainTableA[3:]},
		},
	},
}

func TestCombine(t *testing.T) {
	for _, test := range combineTestCases {
		t.Run(test.name, func(t *testing.T) {
			output := Combine(test.input...)
			if !reflect.DeepEqual(test.expectedOutput, output) {
				t.Errorf("the received result differs from the expected one\n\tExpected: %+v\n\tGot: %+v", test.expectedOutput, output)
				return
			}
		})
	}
}

func BenchmarkCombine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combine(genTable, genTable)
	}
}

// var RSNodeTests = nodeTests{
// 	// #0
// 	{
// 		sample: generalTable{
// 			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
// 			{0.03, 13.03, "d", true, -64 + 3, -64 + 3, -32 + 3, -16 + 3, -8 + 3, 64 + 3, 64 + 3, 32 + 3, 16 + 3, 8 + 3},
// 			{0.04, 14.04, "e", true, -64 + 4, -64 + 4, -32 + 4, -16 + 4, -8 + 4, 64 + 4, 64 + 4, 32 + 4, 16 + 4, 8 + 4},
// 			{0.05, 15.05, "f", true, -64 + 5, -64 + 5, -32 + 5, -16 + 5, -8 + 5, 64 + 5, 64 + 5, 32 + 5, 16 + 5, 8 + 5},
// 			{0.07, 17.07, "h", true, -64 + 7, -64 + 7, -32 + 7, -16 + 7, -8 + 7, 64 + 7, 64 + 7, 32 + 7, 16 + 7, 8 + 7},
// 		},
// 		node: &rsNode{
// 			Data:       genTable,
// 			allocation: []int{0, 3, 4, 5, 7},
// 		},
// 	},
// 	// #1
// 	{
// 		sample: emptyTable,
// 		node: &rsNode{
// 			Data:       genTable,
// 			allocation: []int{},
// 		},
// 	},
// }

// func TestRSNode(t *testing.T) {
// 	nodeTesting(t, RSNodeTests, "rsNode")
// }

// func BenchmarkRSNode(b *testing.B) {
// 	node, _ := RandomSelection(genTable, genTable.Length())
// 	nodeBenchmarking(b, node)
// }

// var RandomSelectionTests = []struct {
// 	selection Data
// 	n         int
// 	isError   bool
// }{
// 	// #0-#5
// 	{selection: genTable, n: genTable.Length() / 2, isError: false},
// 	{selection: mainTableA, n: mainTableA.Length() / 2, isError: false},
// 	{selection: mainTableB, n: mainTableB.Length() / 2, isError: false},
// 	{selection: genTable, n: genTable.Length()/2 + 1, isError: false},
// 	{selection: mainTableA, n: mainTableA.Length()/2 + 1, isError: false},
// 	{selection: mainTableB, n: mainTableB.Length()/2 + 1, isError: false},

// 	// #6-#9
// 	{selection: emptyTable, n: 5, isError: false},
// 	{selection: emptyTable, n: 0, isError: false},
// 	{selection: mainTableX, n: 5, isError: true},
// 	{selection: mainTableX, n: 0, isError: true},

// 	// #10
// 	{selection: nil, n: 5, isError: true},

// 	// #11-#14
// 	{selection: genTable, n: 0, isError: false},
// 	{selection: genTable, n: genTable.Length(), isError: false},
// 	{selection: genTable, n: -1, isError: false},
// 	{selection: genTable, n: genTable.Length() + 1, isError: false},
// }

// func TestRandomSelection(t *testing.T) {
// 	for index, test := range RandomSelectionTests {

// 		partOne, partTwo := RandomSelection(test.selection, test.n)

// 		if !test.isError && (partOne == nil || partTwo == nil) ||
// 			test.isError && (partOne != nil || partTwo != nil) {

// 			t.Errorf("[test %d] error in the test", index)
// 			continue
// 		}

// 		if test.isError {
// 			continue
// 		}

// 		var lengthOne, lengthTwo int
// 		switch {
// 		case test.n <= 0:
// 			lengthOne = 0
// 			lengthTwo = test.selection.Length()

// 		case test.n >= test.selection.Length():
// 			lengthOne = test.selection.Length()
// 			lengthTwo = 0

// 		default:
// 			lengthOne = test.n
// 			lengthTwo = test.selection.Length() - test.n
// 		}

// 		if partOne.Length() != lengthOne || partTwo.Length() != lengthTwo {
// 			t.Errorf("[test %d] the lengths of the resulting selections differ from the expected ones\n"+
// 				"\tthe first selection: expected %d, got %d\n"+
// 				"\tthe second selection: expected %d, got %d",
// 				index, lengthOne, partOne.Length(), lengthTwo, partTwo.Length())
// 			continue
// 		}

// 		joinedSelections := Combine(partOne, partTwo)

// 		if joinedSelections == nil {
// 			t.Errorf("[test %d] failed combine selections for verification", index)
// 			continue
// 		}

// 		if !AreEqual(test.selection, joinedSelections) {
// 			t.Errorf("[test %d] the objects of the sample selection are distributed incorrectly", index)
// 		}
// 	}
// }

// func BenchmarkRandomSelection(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		RandomSelection(genTable, i%genTable.Length())
// 	}
// }

// var EANodeTestCases = nodeTestCase{
// 	{
// 		name: "",
// 		sample: bitOfGeneralTable{
// 			{0.00, "a", true, -64 + 0, -64 + 0, 64 + 0, 64 + 0},
// 			{0.01, "b", true, -64 + 1, -64 + 1, 64 + 1, 64 + 1},
// 			{0.02, "c", true, -64 + 2, -64 + 2, 64 + 2, 64 + 2},
// 			{0.03, "d", true, -64 + 3, -64 + 3, 64 + 3, 64 + 3},
// 			{0.04, "e", true, -64 + 4, -64 + 4, 64 + 4, 64 + 4},
// 			{0.00, "a", true, -64 + 0, -64 + 0, 64 + 0, 64 + 0},
// 		},
// 		node: &eaNode{
// 			Data:       genTable,
// 			allocation: []int{0, 2, 3, 4, 5, 9, 10},
// 		},
// 	},
// }

// func TestEANode(t *testing.T) {
// 	nodeTesting(t, EANodeTestCases, "eaNode")
// }

// func BenchmarkEANode(b *testing.B) {
// 	node := ExtractAttributes(genTable, 0, 2, 3, 4, 5, 9, 10)
// 	nodeBenchmarking(b, node)
// }

var IncludeParamsTestCases = []struct {
	name                  string
	input, expectedOutput Data
	attributeNumbers      []int
}{
	{
		name:             "OK",
		input:            genTable,
		attributeNumbers: []int{0, 2, 3, 4, 5, 9, 10},
		expectedOutput:   bitOfGenTable,
	},
	{
		name:             "Empty array of attribute numbers",
		input:            genTable,
		attributeNumbers: []int{},
		expectedOutput:   nil,
	},
	{
		name:             "Nil array of attribute numbers",
		input:            genTable,
		attributeNumbers: nil,
		expectedOutput:   nil,
	},
	{
		name:             "Array contains an incorrect number (#1)",
		input:            genTable,
		attributeNumbers: []int{0, -1, 1, 2},
		expectedOutput:   nil,
	},
	{
		name:             "Array contains an incorrect number (#2)",
		input:            genTable,
		attributeNumbers: []int{0, 999, 1, 2},
		expectedOutput:   nil,
	},
}

func TestExtractAttributes(t *testing.T) {
	for _, test := range IncludeParamsTestCases {
		t.Run(test.name, func(t *testing.T) {
			result := ExtractAttributes(test.input, test.attributeNumbers...)

			if test.expectedOutput == nil && result != nil || test.expectedOutput != nil && result == nil {
				t.Errorf("the result does not match the expected")
				return
			}

			if test.expectedOutput != nil && !AreEqual(test.expectedOutput, result) {
				t.Errorf("the attributes of the original selection are distributed incorrectly")
			}
		})
	}
}

func BenchmarkExtractAttributes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExtractAttributes(genTable, 0, 2, 3, 4, 5, 9, 10)
	}
}
