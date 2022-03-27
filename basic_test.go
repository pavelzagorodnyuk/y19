package y19

import (
	"reflect"
	"testing"
)

type nodeTests []struct {
	original, node Data
}

func nodeTesting(t *testing.T, tests nodeTests, nodeName string) {
	for index, test := range tests {

		// checking the node methods
		if test.node.Length() != test.original.Length() {
			t.Fatalf("[test %d] %s.Length() is not equal to the expected length\n"+
				"\tExpected: %d\n\tGot: %d\n", index, nodeName, test.original.Length(), test.node.Length())
		}

		if test.node.Dimension() != test.original.Dimension() {
			t.Fatalf("[test %d] %s.Dimension() is not equal to the expected dimension\n"+
				"\tExpected: %d\n\tGot: %d\n", index, nodeName, test.original.Dimension(), test.node.Dimension())
		}

		if test.node.Scale(-1) != IndefiniteScale {
			t.Errorf("[test %d] %s.Scale(-1) is not equal to IndefiniteScale\n\tGot: %v",
				index, nodeName, test.node.Scale(-1))
		}

		if test.node.Scale(test.node.Dimension()) != IndefiniteScale {
			t.Errorf("[test %d] %s.Scale(%s.Dimension()) is not equal to IndefiniteScale\n\tGot: %v",
				index, nodeName, nodeName, test.node.Scale(test.node.Dimension()))
		}

		for p := 0; p < test.original.Dimension(); p++ {
			if test.node.Scale(p) != test.original.Scale(p) {
				t.Errorf("[test %d] %s.Scale(%d) is not equal to the origin\n"+
					"\tExpected: %v\n\tGot: %v\n", index, nodeName, p, test.original.Scale(p), test.node.Scale(p))
				break
			}
		}

		for i := -3; i < test.original.Length()+3; i++ {
			for j := -3; j < test.original.Dimension()+3; j++ {
				if i < 0 || i >= test.original.Length() || j < 0 || j >= test.original.Dimension() {
					if test.node.Value(i, j) != nil {
						t.Fatalf("[test %d] %s.Value(%d, %d) is not equal to nil", index, nodeName, i, j)
					}
				} else if !reflect.DeepEqual(test.node.Value(i, j), test.original.Value(i, j)) {
					t.Fatalf("[test %d] %s.Value(%d, %d) is not equal to the expected value\n"+
						"\tExpected: %v\n\tGot: %v\n", index, nodeName, i, j, test.original.Value(i, j), test.node.Value(i, j))
				}
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

var AreEqualTests = []struct {
	selectionOne, selectionTwo Data
	isEqual                    bool
}{
	// #0-#4
	{selectionOne: mainTableA, selectionTwo: mainTableA, isEqual: true},
	{selectionOne: mainTableB, selectionTwo: mainTableB, isEqual: true},
	{selectionOne: mainTableC, selectionTwo: mainTableC, isEqual: true},
	{selectionOne: emptyTable, selectionTwo: emptyTable, isEqual: true},
	{selectionOne: mainTableX, selectionTwo: mainTableX, isEqual: false},

	// #5-#7
	{selectionOne: nil, selectionTwo: nil, isEqual: false},
	{selectionOne: nil, selectionTwo: mainTableA, isEqual: false},
	{selectionOne: mainTableA, selectionTwo: nil, isEqual: false},

	// #8-#10
	{selectionOne: mainTableA, selectionTwo: mainTableB, isEqual: false},
	{selectionOne: mainTableA, selectionTwo: mainTableC, isEqual: false},
	{selectionOne: mainTableB, selectionTwo: mainTableC, isEqual: false},

	// #11, #13
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		selectionTwo: tableA{
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
		},
		isEqual: true,
	},
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		selectionTwo: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 /**/},
		},
		isEqual: false,
	},
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
		},
		selectionTwo: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
		},
		isEqual: false,
	},
}

func TestAreEqual(t *testing.T) {
	var result bool
	for index, test := range AreEqualTests {
		result = AreEqual(test.selectionOne, test.selectionTwo)
		if test.isEqual != result {
			t.Errorf("[test %d] error in the test\n\tExpected: %t\n\tGot: %t", index, test.isEqual, result)
		}
	}
}

func BenchmarkAreEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreEqual(mainTableA, mainTableA)
	}
}

var AreSimilarTests = []struct {
	selectionOne, selectionTwo Data
	deviation                  float64
	isSimilar                  bool
}{
	// #0-#4
	{selectionOne: mainTableA, selectionTwo: mainTableA, isSimilar: true},
	{selectionOne: mainTableB, selectionTwo: mainTableB, isSimilar: true},
	{selectionOne: mainTableC, selectionTwo: mainTableC, isSimilar: true},
	{selectionOne: emptyTable, selectionTwo: emptyTable, isSimilar: true},
	{selectionOne: mainTableX, selectionTwo: mainTableX, isSimilar: false},

	// #5-#7
	{selectionOne: nil, selectionTwo: nil, isSimilar: false},
	{selectionOne: nil, selectionTwo: mainTableA, isSimilar: false},
	{selectionOne: mainTableA, selectionTwo: nil, isSimilar: false},

	// #8-#10
	{selectionOne: mainTableA, selectionTwo: mainTableB, isSimilar: false},
	{selectionOne: mainTableA, selectionTwo: mainTableC, isSimilar: false},
	{selectionOne: mainTableB, selectionTwo: mainTableC, isSimilar: false},

	// #11-#17
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		selectionTwo: tableA{
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
		},
		isSimilar: true,
	},
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		selectionTwo: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{-0.03 /**/, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		deviation: 0.05,
		isSimilar: true,
	},
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		selectionTwo: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.071 /**/, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		deviation: 0.05,
		isSimilar: false,
	},
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		selectionTwo: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 /**/},
		},
		deviation: 2.,
		isSimilar: true,
	},
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
		},
		selectionTwo: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 5 /**/},
		},
		deviation: 2.,
		isSimilar: false,
	},
	{
		selectionOne: tableB{
			{1001, "Name1", "Lastname1", true, 11},
			{1002, "Name2", "Lastname2", false, 12},
			{1003, "Name3", "Lastname3", true, 13},
		},
		selectionTwo: tableB{
			{1001, "Name1", "Lastname1", true, 11},
			{1002, "Name2", "Lastname2", false, 12},
			{1003, "Name3", "Noname3" /**/, true, 13},
		},
		isSimilar: false,
	},
	{
		selectionOne: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
		},
		selectionTwo: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
			{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
		},
		deviation: 0.05,
		isSimilar: false,
	},
}

func TestAreSimilar(t *testing.T) {
	for index, test := range AreSimilarTests {
		result := AreSimilar(test.selectionOne, test.selectionTwo, test.deviation)
		if test.isSimilar != result {
			t.Errorf("[test %d] error in the test\n\tExpected: %t\n\tGot: %t", index, test.isSimilar, result)
		}
	}
}

func BenchmarkAreSimilar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreSimilar(mainTableA, mainTableA, 0)
	}
}

var CombineNodeTests = nodeTests{
	// #0
	{
		original: tableB{
			{1000, "Name0", "Lastname0", false, 10},
			{1001, "Name1", "Lastname1", true, 11},
			{1002, "Name2", "Lastname2", false, 12},
			{1003, "Name3", "Lastname3", true, 13},
			{1004, "Name4", "Lastname4", false, 14},
			{1005, "Name5", "Lastname5", true, 15},
		},
		node: &combineNode{
			header:        []Scale{AbsoluteScale, NominalScale, NominalScale, NominalScale, NominalScale},
			generalLength: 6,
			startPointers: []int{2, 2, 4},
			sources: []Data{
				tableB{
					{1000, "Name0", "Lastname0", false, 10},
					{1001, "Name1", "Lastname1", true, 11},
				},
				tableB{},
				tableB{
					{1002, "Name2", "Lastname2", false, 12},
					{1003, "Name3", "Lastname3", true, 13},
				},
				tableB{
					{1004, "Name4", "Lastname4", false, 14},
					{1005, "Name5", "Lastname5", true, 15},
				},
			},
		},
	},
	// #1
	{
		original: tableB{
			{1000, "Name0", "Lastname0", false, 10},
			{1001, "Name1", "Lastname1", true, 11},
		},
		node: &combineNode{
			header:        []Scale{AbsoluteScale, NominalScale, NominalScale, NominalScale, NominalScale},
			generalLength: 2,
			startPointers: nil,
			sources: []Data{
				tableB{
					{1000, "Name0", "Lastname0", false, 10},
					{1001, "Name1", "Lastname1", true, 11},
				},
			},
		},
	},
}

func TestCombineNode(t *testing.T) {
	nodeTesting(t, CombineNodeTests, "CombineNode")
}

func BenchmarkCombineNode(b *testing.B) {
	node := Combine(mainTableA, mainTableA)
	nodeBenchmarking(b, node)
}

var CombineTests = []struct {
	fullTable  Data
	tableParts []Data
}{
	// #0-#2
	{
		fullTable: mainTableA,
		tableParts: []Data{
			tableA{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
				{0.03, 13.03, "d", true, -64 + 3, -64 + 3, -32 + 3, -16 + 3, -8 + 3, 64 + 3, 64 + 3, 32 + 3, 16 + 3, 8 + 3},
				{0.04, 14.04, "e", true, -64 + 4, -64 + 4, -32 + 4, -16 + 4, -8 + 4, 64 + 4, 64 + 4, 32 + 4, 16 + 4, 8 + 4},
			},
			tableA{
				{0.05, 15.05, "f", true, -64 + 5, -64 + 5, -32 + 5, -16 + 5, -8 + 5, 64 + 5, 64 + 5, 32 + 5, 16 + 5, 8 + 5},
				{0.06, 16.06, "g", true, -64 + 6, -64 + 6, -32 + 6, -16 + 6, -8 + 6, 64 + 6, 64 + 6, 32 + 6, 16 + 6, 8 + 6},
				{0.07, 17.07, "h", true, -64 + 7, -64 + 7, -32 + 7, -16 + 7, -8 + 7, 64 + 7, 64 + 7, 32 + 7, 16 + 7, 8 + 7},
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			},
		},
	},
	{
		fullTable: mainTableB,
		tableParts: []Data{
			tableB{
				{0, "Name0", "Lastname0", false, 10},
				{1, "Name1", "Lastname1", false, 11},
				{2, "Name2", "Lastname2", false, 12},
				{3, "Name3", "Lastname3", false, 13},
				{4, "Name4", "Lastname4", false, 14},
			},
			tableB{
				{5, "Name5", "Lastname5", false, 15},
				{6, "Name6", "Lastname6", false, 16},
				{7, "Name7", "Lastname7", false, 17},
				{0, "Name0", "Lastname0", false, 10},
			},
		},
	},
	{
		fullTable: mainTableC,
		tableParts: []Data{
			tableC{
				{1000, "Title0", 0.0, 1.0, 50},
				{1001, "Title1", 0.1, 1.1, 51},
				{1002, "Title2", 0.2, 1.2, 52},
				{1003, "Title3", 0.3, 1.3, 53},
				{1004, "Title4", 0.4, 1.4, 54},
			},
			tableC{
				{1005, "Title5", 0.5, 1.5, 55},
				{1006, "Title6", 0.6, 1.6, 56},
				{1007, "Title7", 0.7, 1.7, 57},
				{1000, "Title0", 0.0, 1.0, 50},
			},
		},
	},

	// #3, #4
	{fullTable: emptyTable, tableParts: []Data{emptyTable, emptyTable}},
	{fullTable: nil, tableParts: []Data{mainTableX, mainTableX}},

	// #5-#7
	{fullTable: nil, tableParts: []Data{mainTableA, mainTableA, nil}},
	{fullTable: nil, tableParts: []Data{mainTableA, nil, mainTableA}},
	{fullTable: nil, tableParts: []Data{nil, mainTableA, mainTableA}},

	// #8-#10
	{fullTable: nil, tableParts: []Data{mainTableA, mainTableB}},
	{fullTable: nil, tableParts: []Data{mainTableA, mainTableC}},
	{fullTable: nil, tableParts: []Data{mainTableB, mainTableC}},

	// #11
	{
		fullTable: mainTableA,
		tableParts: []Data{
			tableA{
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
				{0.01, 11.01, "b", true, -64 + 1, -64 + 1, -32 + 1, -16 + 1, -8 + 1, 64 + 1, 64 + 1, 32 + 1, 16 + 1, 8 + 1},
				{0.02, 12.02, "c", true, -64 + 2, -64 + 2, -32 + 2, -16 + 2, -8 + 2, 64 + 2, 64 + 2, 32 + 2, 16 + 2, 8 + 2},
				{0.03, 13.03, "d", true, -64 + 3, -64 + 3, -32 + 3, -16 + 3, -8 + 3, 64 + 3, 64 + 3, 32 + 3, 16 + 3, 8 + 3},
				{0.04, 14.04, "e", true, -64 + 4, -64 + 4, -32 + 4, -16 + 4, -8 + 4, 64 + 4, 64 + 4, 32 + 4, 16 + 4, 8 + 4},
			},
			emptyTable,
			tableA{
				{0.05, 15.05, "f", true, -64 + 5, -64 + 5, -32 + 5, -16 + 5, -8 + 5, 64 + 5, 64 + 5, 32 + 5, 16 + 5, 8 + 5},
				{0.06, 16.06, "g", true, -64 + 6, -64 + 6, -32 + 6, -16 + 6, -8 + 6, 64 + 6, 64 + 6, 32 + 6, 16 + 6, 8 + 6},
				{0.07, 17.07, "h", true, -64 + 7, -64 + 7, -32 + 7, -16 + 7, -8 + 7, 64 + 7, 64 + 7, 32 + 7, 16 + 7, 8 + 7},
				{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			},
		},
	},
}

func TestCombine(t *testing.T) {
	for index, test := range CombineTests {
		result := Combine(test.tableParts...)

		if test.fullTable == nil && result != nil || test.fullTable != nil && result == nil {
			t.Errorf("[test %d] error in the test", index)
			continue
		}

		if test.fullTable != nil && !AreEqual(test.fullTable, result) {
			t.Errorf("[test %d] combined selections does not equal the expected one", index)
		}
	}
}

func BenchmarkCombine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combine(mainTableA, mainTableA)
	}
}

var RSNodeTests = nodeTests{
	// #0
	{
		original: tableA{
			{0.00, 10.00, "a", true, -64 + 0, -64 + 0, -32 + 0, -16 + 0, -8 + 0, 64 + 0, 64 + 0, 32 + 0, 16 + 0, 8 + 0},
			{0.03, 13.03, "d", true, -64 + 3, -64 + 3, -32 + 3, -16 + 3, -8 + 3, 64 + 3, 64 + 3, 32 + 3, 16 + 3, 8 + 3},
			{0.04, 14.04, "e", true, -64 + 4, -64 + 4, -32 + 4, -16 + 4, -8 + 4, 64 + 4, 64 + 4, 32 + 4, 16 + 4, 8 + 4},
			{0.05, 15.05, "f", true, -64 + 5, -64 + 5, -32 + 5, -16 + 5, -8 + 5, 64 + 5, 64 + 5, 32 + 5, 16 + 5, 8 + 5},
			{0.07, 17.07, "h", true, -64 + 7, -64 + 7, -32 + 7, -16 + 7, -8 + 7, 64 + 7, 64 + 7, 32 + 7, 16 + 7, 8 + 7},
		},
		node: &rsNode{
			Data:       mainTableA,
			allocation: []int{0, 3, 4, 5, 7},
		},
	},
	// #1
	{
		original: emptyTable,
		node: &rsNode{
			Data:       mainTableA,
			allocation: []int{},
		},
	},
}

func TestRSNode(t *testing.T) {
	nodeTesting(t, RSNodeTests, "rsNode")
}

func BenchmarkRSNode(b *testing.B) {
	node, _ := RandomSelection(mainTableA, mainTableA.Length())
	nodeBenchmarking(b, node)
}

var RandomSelectionTests = []struct {
	selection Data
	n         int
	isError   bool
}{
	// #0-#5
	{selection: mainTableA, n: mainTableA.Length() / 2, isError: false},
	{selection: mainTableB, n: mainTableB.Length() / 2, isError: false},
	{selection: mainTableC, n: mainTableC.Length() / 2, isError: false},
	{selection: mainTableA, n: mainTableA.Length()/2 + 1, isError: false},
	{selection: mainTableB, n: mainTableB.Length()/2 + 1, isError: false},
	{selection: mainTableC, n: mainTableC.Length()/2 + 1, isError: false},

	// #6-#9
	{selection: emptyTable, n: 5, isError: false},
	{selection: emptyTable, n: 0, isError: false},
	{selection: mainTableX, n: 5, isError: true},
	{selection: mainTableX, n: 0, isError: true},

	// #10
	{selection: nil, n: 5, isError: true},

	// #11-#14
	{selection: mainTableA, n: 0, isError: false},
	{selection: mainTableA, n: mainTableA.Length(), isError: false},
	{selection: mainTableA, n: -1, isError: false},
	{selection: mainTableA, n: mainTableA.Length() + 1, isError: false},
}

func TestRandomSelection(t *testing.T) {
	for index, test := range RandomSelectionTests {

		partOne, partTwo := RandomSelection(test.selection, test.n)

		if !test.isError && (partOne == nil || partTwo == nil) ||
			test.isError && (partOne != nil || partTwo != nil) {

			t.Errorf("[test %d] error in the test", index)
			continue
		}

		if test.isError {
			continue
		}

		var lengthOne, lengthTwo int
		switch {
		case test.n <= 0:
			lengthOne = 0
			lengthTwo = test.selection.Length()

		case test.n >= test.selection.Length():
			lengthOne = test.selection.Length()
			lengthTwo = 0

		default:
			lengthOne = test.n
			lengthTwo = test.selection.Length() - test.n
		}

		if partOne.Length() != lengthOne || partTwo.Length() != lengthTwo {
			t.Errorf("[test %d] the lengths of the resulting selections differ from the expected ones\n"+
				"\tthe first selection: expected %d, got %d\n"+
				"\tthe second selection: expected %d, got %d",
				index, lengthOne, partOne.Length(), lengthTwo, partTwo.Length())
			continue
		}

		joinedSelections := Combine(partOne, partTwo)

		if joinedSelections == nil {
			t.Errorf("[test %d] failed combine selections for verification", index)
			continue
		}

		if !AreEqual(test.selection, joinedSelections) {
			t.Errorf("[test %d] the objects of the original selection are distributed incorrectly", index)
		}
	}
}

func BenchmarkRandomSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomSelection(mainTableA, i%mainTableA.Length())
	}
}

var EANodeTestCases = nodeTests{
	{
		original: bitOfTableA{
			{0.00, "a", true, -64 + 0, -64 + 0, 64 + 0, 64 + 0},
			{0.01, "b", true, -64 + 1, -64 + 1, 64 + 1, 64 + 1},
			{0.02, "c", true, -64 + 2, -64 + 2, 64 + 2, 64 + 2},
			{0.03, "d", true, -64 + 3, -64 + 3, 64 + 3, 64 + 3},
			{0.04, "e", true, -64 + 4, -64 + 4, 64 + 4, 64 + 4},
			{0.05, "f", true, -64 + 5, -64 + 5, 64 + 5, 64 + 5},
			{0.06, "g", true, -64 + 6, -64 + 6, 64 + 6, 64 + 6},
			{0.07, "h", true, -64 + 7, -64 + 7, 64 + 7, 64 + 7},
			{0.00, "a", true, -64 + 0, -64 + 0, 64 + 0, 64 + 0},
		},
		node: &eaNode{
			Data:       mainTableA,
			allocation: []int{0, 2, 3, 4, 5, 9, 10},
		},
	},
}

func TestEANode(t *testing.T) {
	nodeTesting(t, EANodeTestCases, "eaNode")
}

func BenchmarkEANode(b *testing.B) {
	node := ExtractAttributes(mainTableA, 0, 2, 3, 4, 5, 9, 10)
	nodeBenchmarking(b, node)
}

var IncludeParamsTestCases = []struct {
	name                  string
	input, expectedOutput Data
	attributeNumbers      []int
}{
	{
		name:             "OK",
		input:            mainTableA,
		attributeNumbers: []int{0, 2, 3, 4, 5, 9, 10},
		expectedOutput:   mainBitA,
	},
	{
		name:             "Empty array of attribute numbers",
		input:            mainTableA,
		attributeNumbers: []int{},
		expectedOutput:   nil,
	},
	{
		name:             "Nil array of attribute numbers",
		input:            mainTableA,
		attributeNumbers: nil,
		expectedOutput:   nil,
	},
	{
		name:             "Array contains an incorrect number (#1)",
		input:            mainTableA,
		attributeNumbers: []int{0, -1, 1, 2},
		expectedOutput:   nil,
	},
	{
		name:             "Array contains an incorrect number (#2)",
		input:            mainTableA,
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
		ExtractAttributes(mainTableA, 0, 2, 3, 4, 5, 9, 10)
	}
}
