package y19

import (
	"math"
	"reflect"
)

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

// AreSimilar reports whether selection1 and selection2 data selections are similar.
func AreSimilar(selection1, selection2 Data, deviation float64) bool {

	if selection1 == nil || selection2 == nil || math.IsNaN(deviation) || deviation < 0 {
		return false
	}

	var length = selection1.Length()
	var dimension = selection1.Dimension()

	if length != selection2.Length() || dimension != selection2.Dimension() || length < 0 || dimension < 1 {
		return false
	}

	// checking similarity of scales
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

	// checking whether the selections contain similar objects
	for i := 0; i < length; i++ {
		found = false

		// reading an object
		for k := 0; k < dimension; k++ {
			object[k] = selection1.Value(i, k)
		}

		// search for a similar object in the selection2
	LOOP:
		for j := 0; j < length && !found; j++ {
			if marks[j] {
				continue
			}

			for k := 0; k < dimension; k++ {

				value1, ok1 := interfaceToFloat64(object[k])
				value2, ok2 := interfaceToFloat64(selection2.Value(j, k))

				if ok1 && ok2 {
					if math.Abs(value2-value1) > deviation {
						continue LOOP
					}
				} else if !reflect.DeepEqual(object[k], selection2.Value(j, k)) {
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

func interfaceToFloat64(source interface{}) (float64, bool) {

	var value float64 = math.NaN()
	var ok bool

	switch source.(type) {
	case float64:
		value = source.(float64)
		ok = true
	case float32:
		value = float64(source.(float32))
		ok = true

	case int:
		value = float64(source.(int))
		ok = true
	case int64:
		value = float64(source.(int64))
		ok = true
	case int32:
		value = float64(source.(int32))
		ok = true
	case int16:
		value = float64(source.(int16))
		ok = true
	case int8:
		value = float64(source.(int8))
		ok = true

	case uint:
		value = float64(source.(uint))
		ok = true
	case uint64:
		value = float64(source.(uint64))
		ok = true
	case uint32:
		value = float64(source.(uint32))
		ok = true
	case uint16:
		value = float64(source.(uint16))
		ok = true
	case uint8:
		value = float64(source.(uint8))
		ok = true
	}

	return value, ok
}
