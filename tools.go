package y19

import "math"

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

	case *float64:
		value = *(source.(*float64))
		ok = true
	case *float32:
		value = float64(*(source.(*float32)))
		ok = true

	case *int:
		value = float64(*(source.(*int)))
		ok = true
	case *int64:
		value = float64(*(source.(*int64)))
		ok = true
	case *int32:
		value = float64(*(source.(*int32)))
		ok = true
	case *int16:
		value = float64(*(source.(*int16)))
		ok = true
	case *int8:
		value = float64(*(source.(*int8)))
		ok = true

	case *uint:
		value = float64(*(source.(*uint)))
		ok = true
	case *uint64:
		value = float64(*(source.(*uint64)))
		ok = true
	case *uint32:
		value = float64(*(source.(*uint32)))
		ok = true
	case *uint16:
		value = float64(*(source.(*uint16)))
		ok = true
	case *uint8:
		value = float64(*(source.(*uint8)))
		ok = true
	}

	return value, ok
}

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
