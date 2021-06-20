package lc13

import (
	"fmt"
	"testing"
)

func romanToInt(s string) int {
	return combine(readRoman(s))
}

func combine(set []int) int {
	fmt.Println(set)
	res := 0
	for i:=0; i<len(set); i++ {
		if i+1 >= len(set) {
			res += set[i]
			continue
		}

		switch set[i] {
		case 1:
			if set[i+1] == 5 {
				i++
				res += 4
				continue
			}
			if set[i+1] == 10 {
				i++
				res += 9
				continue
			}
			res += set[i]
		case 10:
			if set[i+1] == 50 {
				i++
				res += 40
				continue
			}
			if set[i+1] == 100 {
				i++
				res += 90
				continue
			}
			res += set[i]
		case 100:
			if set[i+1] == 500 {
				i++
				res += 400
				continue
			}
			if set[i+1] == 1000 {
				i++
				res += 900
				continue
			}
			res += set[i]
		default:
			res += set[i]
		}
	}

	return res
}

func convert(target int32) int {
	switch target {
	case 73:
		return 1
	case 86:
		return 5
	case 88:
		return 10
	case 76:
		return 50
	case 67:
		return 100
	case 68:
		return 500
	case 77:
		return 1000
	default:
		return 0
	}
}

func readRoman(str string) []int {
	var set []int
	for _, s := range str {
		set = append(set, convert(s))
	}
	return set
}

func TestLC13 (t *testing.T) {
	str := "III"
	println(romanToInt(str))
}
