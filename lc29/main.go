package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
}
func divide(dividend int, divisor int) int {
	dd := int64(dividend)
	dr := int64(divisor)

	q := dd / dr

	if q > math.MaxInt32 {
		return math.MaxInt32
	}

	if q < math.MinInt32 {
		return math.MinInt32
	}

	return int(q)
}
