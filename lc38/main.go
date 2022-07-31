package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	ret := "1"
	for i := 2; i <= n; i++ {
		ret = worker(ret)
	}

	return ret
}

func worker(str string) string {

	if len(str) == 0 {
		return ""
	}

	if len(str) == 1 {
		return "1" + str
	}

	count := 1
	ret := ""
	ch := rune(str[0])
	for _, c := range str[1:] {
		if c == ch {
			count++
			continue
		}

		ret += strconv.Itoa(count) + string(ch)

		// reset
		count = 1
		ch = c
	}

	ret += strconv.Itoa(count) + string(ch)
	return ret
}
