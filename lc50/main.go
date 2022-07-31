package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(2.10000, 3))
}

func myPow(x float64, n int) float64 {

	res := math.Pow(x, float64(n))

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}
