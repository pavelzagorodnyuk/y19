package y19

import "reflect"

// AreEqual reports whether selection1 and selection2 data selections are equivalent.
func AreEqual(selection1, selection2 Data) bool {

	if selection1 == nil || selection2 == nil {
		return false
	}

	var length = selection1.Length()
	var dimension = selection1.Dimension()

	if length != selection2.Length() || dimension != selection2.Dimension() || length < 0 || dimension < 1 {
		return false
	}

	for i := 0; i < dimension; i++ {
		if selection1.Scale(i) != selection2.Scale(i) {
			return false
		}
	}

	var (
		found  bool
		marks  = make([]bool, length)
		object = make([]interface{}, dimension)
	)

	// checking whether the selections contain the same objects
	for i := 0; i < length; i++ {
		found = false

		// reading an object
		for k := 0; k < dimension; k++ {
			object[k] = selection1.Value(i, k)
		}

		// search for the same object in the selection2
	LOOP:
		for j := 0; j < length && !found; j++ {
			if marks[j] {
				continue
			}

			for k := 0; k < dimension; k++ {
				if !reflect.DeepEqual(object[k], selection2.Value(j, k)) {
					continue LOOP
				}
			}

			found = true
			marks[j] = true
		}

		// if a duplicate is not found, then finish the search
		if !found {
			return false
		}
	}

	return true
}
