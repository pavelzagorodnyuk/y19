package y19

/*
var zScoreNodeTestCases = nodeTests{
	{
		original: genTable,
		node: &zScoreNode{
			Data: genTable,
			header: map[int]struct {
				averageValue      float64
				standardDeviation float64
			}{
				6: {averageValue: 6, standardDeviation: 6},
			},
		},
	},
}


var ZScoreNormalizationTests = []struct {
	input, output Data
	note          ZScoreNote
}{
	{
		input: tableC{
			{1000, "Title0", 0.0, 1.0, 50},
			{1001, "Title1", 0.1, 1.1, 51},
			{1002, "Title2", 0.2, 1.2, 52},
			{1003, "Title3", 0.3, 1.3, 53},
		},
		output: tableC{
			{1000, "Title0", 0.0, 1.0, 50},
			{1001, "Title1", 0.1, 1.1, 51},
			{1002, "Title2", 0.2, 1.2, 52},
			{1003, "Title3", 0.3, 1.3, 53},
		},
		note: ZScoreNote{
			1: struct {
				averageValue      float64
				standardDeviation float64
			}{
				averageValue:      0,
				standardDeviation: 0,
			},
		},
	},
}

func TestZScoreNormalization(t *testing.T) {
	for index, test := range ZScoreNormalizationTests {
		result, _ := ZScoreNormalization(test.input)

		if test.output == nil && result != nil ||
			test.output != nil && result == nil {
			t.Errorf("[test %d] error in the test", index)
			continue
		}

		if test.output != nil && !AreEqual(test.output, result) {
			t.Errorf("[test %d] the data is normalized incorrectly", index)
		}
	}
}

func BenchmarkZScoreNormalization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZScoreNormalization(mainTableA)
	}
}

*/
