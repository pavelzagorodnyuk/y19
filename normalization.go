package y19

import "math"

type normalizationNode interface {
	normalizeDataAs(Data) Data
	normalizeObjectAs(Object) Object
}

func NormalizeDataAs(selection Data, sample Data) Data {
	if selection == nil || sample == nil {
		return nil
	}
	if node, ok := sample.(normalizationNode); ok {
		return node.normalizeDataAs(selection)
	}
	return nil
}

func NormalizeObjectAs(object Object, sample Data) Object {
	if object == nil || sample == nil {
		return nil
	}
	if node, ok := sample.(normalizationNode); ok {
		return node.normalizeObjectAs(object)
	}
	return nil
}

type zScoreNode struct {
	Data
	header map[int]struct {
		averageValue      float64
		standardDeviation float64
	}
}

func (node *zScoreNode) Value(o int, p int) interface{} {

	if p < 0 || p >= node.Dimension() {
		return nil
	}

	if _, ok := node.header[p]; !ok {
		return node.Data.Value(o, p)
	}

	sourceValue, ok := interfaceToFloat64(node.Value(o, p))
	if !ok {
		return nil
	}

	return (sourceValue - node.header[p].averageValue) / node.header[p].standardDeviation
}

func (node *zScoreNode) normalizeDataAs(selection Data) Data {

	if !areHeadersTheSame(node.Data, selection) {
		return nil
	}

	newNode := new(zScoreNode)
	newNode.Data = selection
	newNode.header = node.header

	return newNode
}

func (node *zScoreNode) normalizeObjectAs(object Object) Object {
	return nil
}

func ZScoreNormalization(selection Data) Data {

	if selection == nil {
		return nil
	}

	var length = selection.Length()
	var dimension = selection.Dimension()

	if length < 0 || dimension < 1 {
		return nil
	}

	var node = new(zScoreNode)
	node.Data = selection
	node.header = make(map[int]struct {
		averageValue      float64
		standardDeviation float64
	})

	for i := 0; i < dimension; i++ {
		if selection.Scale(i) != AbsoluteScale {
			continue
		}

		var sum float64
		var count int
		for j := 0; j < length; j++ {
			value, ok := interfaceToFloat64(selection.Value(j, i))
			if !ok {
				continue
			}

			sum += value
			count++
		}

		node.header[i] = struct {
			averageValue      float64
			standardDeviation float64
		}{
			averageValue: sum / float64(count),
		}

		sum, count = 0, 0
		for j := 0; j < length; j++ {
			number, ok := interfaceToFloat64(selection.Value(j, i))
			if !ok {
				continue
			}

			a := number - node.header[i].averageValue
			sum += a * a
			count++
		}

		record := node.header[i]
		record.standardDeviation = math.Sqrt(sum / float64(count-1))
		node.header[i] = record
	}

	return node
}
