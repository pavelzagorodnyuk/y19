package y19

import "testing"

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
