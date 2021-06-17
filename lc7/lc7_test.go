package lc7

import (
	"math"
	"testing"
)

func reverse(x int) int {
	val := combine(cut(legal(x)))
	if positive(x) {
		return legalResult(val)
	} else {
		return legalResult(0-val)
	}
}

func pow10(n int) int {
	return int(math.Pow10(n))
}

func positive(n int) bool {
	return n>0
}

func legal(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func cut(x int) []int {
	var num []int
	i := 0
	for {
		i++

		tmp := x/pow10(i)

		tmp = tmp * pow10(i)

		val := (x - tmp)/pow10(i-1)

		num = append(num, val)

		x = tmp

		if x == 0 {
			break
		}

	}

	return num
}

func combine(num []int) int {
	res := 0

	i := 0
	for {
		tmp := num[i] * pow10(len(num)-1-i)

		res += tmp

		i++

		if i == len(num) {
			break
		}
	}

	return res
}

func legalResult(res int) int {
	if res < math.MinInt32 {
		return 0
	}

	if res > math.MaxInt32 {
		return 0
	}

	return res
}

func TestLC7(t *testing.T) {
	println(reverse(1534236469))
}
