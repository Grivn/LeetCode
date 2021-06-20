package lc8

import (
	"math"
	"testing"
)

func myAtoi(s string) int {

	set := readString(s)

	val := uint64(0)

	pos := true

	flag := false

	for _, c := range set {
		typ := filter(c)

		if typ == blank {
			if flag {
				break
			}
			continue
		}

		if typ == positive {
			if flag {
				break
			}
			pos = true
			flag = true
		}

		if typ == negative {
			if flag {
				break
			}
			pos = false
			flag = true
		}

		if typ == illegal {
			break
		}

		if typ == number {
			flag = true
			tmp := convertCharToNumber(c)
			val = val*10+uint64(tmp)
		}

		if pos {
			if !checkPos(val) {
				break
			}
		} else {
			if !checkNeg(val) {
				break
			}
		}
	}

	if pos {
		return makePosLegal(val)
	} else {
		return makeNegLegal(val)
	}
}

func readString(s string) []int32 {
	var res []int32
	for _, c := range s {
		res = append(res, c)
	}
	return res
}

const (
	blank = iota
	negative
	positive
	number
	illegal
)

type characterType int

func filter(c int32) characterType {
	switch c {
	case 32:
		return blank
	case 43:
		return positive
	case 45:
		return negative
	default:
		if c >= 48 && c <= 57 {
			return number
		} else {
			return illegal
		}
	}
}

func convertCharToNumber(c int32) int32 {
	return c-48
}

func makePosLegal(num uint64) int {
	if !checkPos(num) {
		return math.MaxInt32
	}

	return int(num)
}

func makeNegLegal(num uint64) int {
	if !checkNeg(num) {
		return math.MinInt32
	}

	return int(0-num)
}

func checkPos(num uint64) bool {
	if num > math.MaxInt32 {
		return false
	}
	return true
}

func checkNeg(num uint64) bool {
	if num > -math.MinInt32 {
		return false
	}
	return true
}

func TestLC8(t *testing.T) {
	str := "18446744073709551617"

	println(myAtoi(str))
}
