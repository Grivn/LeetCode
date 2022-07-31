package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("1111111", "11111111"))
}

func multiply(num1 string, num2 string) string {
	index1 := len(num1) - 1
	index2 := len(num2) - 1

	exceed := 0
	res := ""
	for {
		v1, v2 := 0, 0

		if index1 >= 0 {
			c1 := num1[index1]
			v1, _ = strconv.Atoi(string(c1))
		}

		if index2 >= 0 {
			c2 := num2[index2]
			v2, _ = strconv.Atoi(string(c2))
		}

		bit := v1*v2 + exceed

		h := bit / 10
		l := bit - h*10
		exceed = h

		fmt.Printf("v1 %d, v2 %d, ex %d, l %d, bit %d\n", v1, v2, exceed, l, bit)
		if l == 0 && exceed == 0 {
			break
		}
		index1--
		index2--

		res += strconv.Itoa(l)
	}

	ret := ""

	for i := len(res) - 1; i >= 0; i-- {
		ret += string(res[i])
	}
	return ret
}
