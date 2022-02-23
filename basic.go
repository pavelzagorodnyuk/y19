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

type combineNode struct {
	header        []Scale
	generalLength int   // this is the total number of objects
	startPointers []int // len(startPointers) == len(sources)-1
	sources       []Data
}

func (node *combineNode) Length() int {
	return node.generalLength
}

func (node *combineNode) Dimension() int {
	return len(node.header)
}

func (node *combineNode) Scale(p int) Scale {
	if p < 0 || p >= len(node.header) {
		return IndefiniteScale
	}

	return node.header[p]
}

func (node *combineNode) Value(o, p int) interface{} {

	if o < 0 || o >= node.generalLength {
		return nil
	}

	if len(node.sources) == 1 || o < node.startPointers[0] {
		return node.sources[0].Value(o, p)
	}

	for i := 0; i < len(node.startPointers)-1; i++ {
		if o >= node.startPointers[i] && o < node.startPointers[i+1] {
			return node.sources[i+1].Value(o-node.startPointers[i], p)
		}
	}

	return node.sources[len(node.sources)-1].Value(o-node.startPointers[len(node.startPointers)-1], p)
}

// Combine combines objects from several selections into one, provided that they have the same headers.
func Combine(selections ...Data) Data {

	if len(selections) == 0 || selections[0] == nil {
		return nil
	}

	var dimension = selections[0].Dimension()
	var length = selections[0].Length()

	if dimension < 0 || length < 0 {
		return nil
	}

	node := new(combineNode)
	node.header = make([]Scale, dimension)
	node.generalLength = selections[0].Length()
	node.sources = selections
	if len(selections) > 1 {
		node.startPointers = make([]int, len(selections)-1)
	}

	for i := 0; i < dimension; i++ {
		node.header[i] = selections[0].Scale(i)
	}

	// checking the others selections and writing values in the node
	for i := 1; i < len(selections); i++ {

		if selections[i] == nil {
			return nil
		}

		length := selections[i].Length()
		if dimension != selections[i].Dimension() || length < 0 {
			return nil
		}

		// checking the equivalence of headers
		for j := 0; j < dimension; j++ {
			if node.header[j] != selections[i].Scale(j) {
				return nil
			}
		}

		node.startPointers[i-1] = node.generalLength
		node.generalLength += length
	}

	return node
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
