package lc9

import (
	"testing"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	set := readSet(x)

	val := reConstruct(set)

	return val == x
}

func readSet(x int) []int {
	var set []int

	for {
		tmp := x/10

		val := x-tmp*10

		set = append(set, val)

		x -= val

		x = x/10

		if x == 0 {
			break
		}
	}
	return set
}

func reConstruct(set []int) int {
	val := 0
	for _, num := range set {
		val = val*10 + num
	}
	return val
}

func TestLC9 (t *testing.T) {
	num := 123
	println(isPalindrome(num))
}
