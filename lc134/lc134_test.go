package lc134

import "testing"

func canCompleteCircuit(gas []int, cost []int) int {
	for startIndex := range gas {
		totalGas := 0

		totalGas += gas[startIndex]
		totalGas -= cost[startIndex]
		if totalGas < 0 {
			continue
		}

		index := startIndex
		for {
			index++
			if index%len(gas) == startIndex {
				return index%len(gas)
			}

			totalGas += gas[index%len(gas)]
			totalGas -= cost[index%len(gas)]

			if totalGas < 0 {
				break
			}
		}
	}

	return -1
}


func TestLC134(t *testing.T) {
	gas := []int{1,2,3,4,5}

	cost := []int{3,4,5,1,2}

	println(canCompleteCircuit(gas, cost))
}
