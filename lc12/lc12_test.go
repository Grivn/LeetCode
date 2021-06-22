package lc12

import (
	"fmt"
	"math"
	"testing"
)

func intToRoman(num int) string {
	set := breakInt(num)

	str := ""
	for i:=len(set)-1; i>=0; i-- {
		c := filter(pow10(i), set[i])
		str += c
	}

	return str
}

func breakInt(num int) []int {
	var set []int

	i := 0
	for {
		i++

		multi := pow10(i)

		tmp := num/multi

		val := num - tmp*multi

		set = append(set, val)

		num -= val

		if num == 0 {
			break
		}
	}

	return set
}

func filter(target int, val int) string {
	switch target {
	case 1:
		return generatePhase(val, "I", "V", "X")
	case 10:
		return generatePhase(val/10, "X", "L", "C")
	case 100:
		return generatePhase(val/100, "C", "D", "M")
	case 1000:
		return generatePhase(val/1000, "M", "", "")
	default:
		panic("error")
	}
}

func generatePhase(val int, base string, middle string, upper string) string {
	if val < 4 {
		str := ""
		for i:=0; i<val; i++ {
			str = str + base
		}
		return str
	} else if val == 4 {
		return base+middle
	} else if val > 4 && val < 9 {
		str := middle
		for i:=5; i<val; i++ {
			str = str + base
		}
		return str
	} else if val == 9 {
		return base+upper
	} else {
		return ""
	}
}

func pow10(i int) int {
	return int(math.Pow10(i))
}

func TestLC12(t *testing.T) {
	num := 1994

	fmt.Println(intToRoman(num))
}
