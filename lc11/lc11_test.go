package lc11

import (
	"math"
	"testing"
)

func maxArea(height []int) int {
	i := 0
	j := len(height)-1
	v := 0

	for {
		v = int(math.Max(float64(v), float64(calculateV(i, j, height[i], height[j]))))

		if height[i] > height[j] {
			j--
		} else {
			i++
		}

		if i >= j {
			break
		}
	}


	return v
}

func maxAreaOn2(height []int) int {
	v := 0

	for i:=0; i<len(height); i++ {
		for j:=i+1; j<len(height); j++ {
			v = int(math.Max(float64(v), float64(calculateV(i, j, height[i], height[j]))))
		}
	}

	return v
}

func calculateV(i, j, n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2))) * int(math.Abs(float64(i-j)))
}

func TestLC11(t *testing.T) {
	height := []int{1,8,6,2,5,4,8,3,7}
	println(maxAreaOn2(height))
	println(maxArea(height))
}
