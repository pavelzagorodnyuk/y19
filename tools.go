package y19

func areHeadersTheSame(selections ...Data) bool {

	if len(selections) == 0 {
		return false
	}

	var dimension = selections[0].Dimension()
	for i := 1; i < len(selections); i++ {
		if dimension != selections[i].Dimension() {
			return false
		}
	}

	var header = make([]Scale, dimension)
	for i := 0; i < dimension; i++ {
		header[i] = selections[0].Scale(i)
	}

	for i := 1; i < len(selections); i++ {
		for j := 0; j < dimension; j++ {
			if header[j] != selections[i].Scale(j) {
				return false
			}
		}
	}

	return true
}
