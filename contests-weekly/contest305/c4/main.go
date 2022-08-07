package main

import (
	"fmt"
	"math"
)

func main() {
	str := "acfgbd"
	k := 2
	fmt.Println(format(str, k))
	fmt.Println(str)
}
func longestIdealString(s string, k int) int {
	return worker1(s, k, 0)
}

func longestIdealStringProcessor(s string, fmt string) int {
	res := make([]int, len(s)+1)

	for i := range s {
		if i == 0 {
			res[0] = 1
			continue
		}

		res[i] = res[i-1]

	}

	return len(res)
}

func format(str string, k int) string {
	s := ""
	for _, c := range str {
		cut := c - rune(k)
		if cut < 'a' {
			s += "$"
		} else {
			s += string(cut)
		}
	}

	return s
}

func worker1(s string, k int, cut int) int {
	loop := 0
	var subs []string
	subs = append(subs, s)
	var nextSubs []string
	checked := make(map[string]bool)
	for {
		loop++
		for _, s := range subs {
			//if checked[s] {
			//	continue
			//}
			//fmt.Println("to check:", s, "len:", len(s))
			if checker(s, k) {
				checked[s] = true
				return len(s)
			}
			for i := range s {
				str := ""
				if len(s)-1 == i {
					str = s[:i]
				} else {
					str = s[:i] + s[i+1:]
				}
				if len(s) == 1 {
					return 0
				}
				nextSubs = append(nextSubs, str)
				//fmt.Println("to check:", str, "len:", len(str))
				checked[str] = true
				if checker(str, k) {
					return len(str)
				}
			}
		}
		subs = []string{}
		subs = append(subs, nextSubs...)
		nextSubs = []string{}
		//fmt.Println("sub:", subs)
		//panic("err")
	}
}

func checkSubs(s string, k int, cut int) int {
	fmt.Println("to check:", s, "cut: ", cut)
	if cut == 0 {
		if checker(s, k) {
			return len(s)
		}
		return 0
	}

	for i := range s {
		str := ""
		if len(s)-1 == i {
			str = s[:i]
		} else {
			str = s[:i] + s[i+1:]
		}
		ret := checkSubs(str, k, cut-1)
		if ret != 0 {
			return ret
		}
	}
	return 0
}

// long to short
func worker(remain string, k int) int {
	fmt.Println("to check:", remain)
	if checker(remain, k) {
		fmt.Println("valid:", remain)
		return len(remain)
	}

	for i := range remain {
		str := ""
		if i == len(remain)-1 {
			str = remain[:i]
		} else {
			str = remain[:i] + remain[i+1:]
		}

		if checker(str, k) {
			fmt.Println("valid:", remain)
			return len(str)
		}
	}

	longest := 0
	for i := range remain {
		str := ""
		if i == len(remain)-1 {
			str = remain[:i]
		} else {
			str = remain[:i] + remain[i+1:]
		}
		tmp := worker(str, k)
		if tmp > longest {
			longest = tmp
		}
	}
	return longest
}

func checker(str string, k int) bool {
	for i, c := range str {
		if i == len(str)-1 {
			continue
		}

		next := str[i+1]
		if math.Abs(float64(next)-float64(c)) > float64(k) {
			return false
		}
	}
	return true
}
