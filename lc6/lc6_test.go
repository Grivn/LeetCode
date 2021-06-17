package lc6

import (
	"math"
	"testing"
)

func convert(s string, numRows int) string {
	return secondCombine(firstCut(s, numRows), numRows)
}

func firstCut(s string, numRows int) []string {
	var set []string
	groupLen := numRows + (numRows - 2)

	if numRows <= 2 {
		groupLen = numRows
	}
	i := 0
	for {
		if i >= len(s) {
			break
		}
		length := int(math.Min(float64(groupLen), float64(len(s)-i)))
		sub := s[i:i+length]
		set = append(set, sub)
		i += groupLen
	}

	return set
}

func secondCombine(subSet []string, numRows int) string {
	str := ""
	groupLen := numRows + (numRows - 2)
	for i:=0; i<numRows; i++ {
		if i == 0 {
			for _, sub := range subSet {
				str += sub[:1]
			}
			continue
		}

		if i == numRows-1 {
			for _, sub := range subSet {
				if len(sub) < numRows {
					continue
				}
				str += sub[i:numRows]
			}
			break
		}

		for _, sub := range subSet {
			if len(sub) >= i+1 {
				str += sub[i:i+1]
			}

			if len(sub) >= groupLen-(i-1) {
				str += sub[groupLen-(i-1)-1:groupLen-(i-1)]
			}
		}
	}
	return str
}

func TestLC6(t *testing.T) {
	println(convert("PAYPALISHIRING", 4))
}
