package lc42

import (
	"math"
	"testing"
)

func trap(height []int) int {
	preI := 0
	preJ := 0
	i := 0
	j := len(height)-1
	amount := 0

	for {
		if i < j {
			if height[i] > preI {
				preI = height[i]
				i++
			} else if height[i] == preI {
				i++
			} else {
				if preJ >= preI {
					amount += preI-height[i]
					i++
				} else {
					// do nothing
				}
			}

			if height[j] > preJ {
				preJ = height[j]
				j--
			} else if height[j] == preJ {
				j--
			} else {
				if preI >= preJ {
					amount += preJ-height[j]
					j--
				} else {
					// do nothing
				}
			}
		}

		if i > j {
			break
		}

		if i == j {
			if preJ > height[j] && preI > height[i] {
				amount += int(math.Min(float64(preI), float64(preJ)))-height[j]
			}
			break
		}
	}

	return amount
}

func TestLC42(t *testing.T) {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	println(trap(height))
}
