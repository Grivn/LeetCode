package main

import "fmt"

func main() {
	fmt.Println(secondsToRemoveOccurrencesV1("0110101"))
}

func secondsToRemoveOccurrences(s string) int {
	seconds := 0
	count1 := 0
	lastIndex := 0
	for i, v := range s {
		if v == '0' {
			continue
		}
		count1++
		lastIndex = i

		tmp := lastIndex - (count1 - 1)
		if tmp > seconds+1 {
			seconds = tmp
		} else {
			seconds++
		}
	}

	if count1 == lastIndex+1 {
		return 0
	}
	return seconds
}

func secondsToRemoveOccurrences2(s string) int {
	if len(s) == 1 {
		return 0
	}

	if s == "01" {
		return 1
	}

	if s == "10" {
		//fmt.Println("test2")
		return 0
	}

	count1 := 0
	lastIndex := 0
	for i, v := range s {
		if v == '0' {
			continue
		}

		count1++
		if i > lastIndex {
			lastIndex = i
		}
	}

	countSerialLen := 0
	for i := lastIndex; i >= 0; i-- {
		if s[lastIndex] == '1' {
			countSerialLen++
		} else {
			break
		}
	}

	if lastIndex == count1-1 {
		return 0
	}

	return lastIndex - (count1 - 1) + countSerialLen
}

func secondsToRemoveOccurrencesV1(s string) int {
	if len(s) == 1 {
		//fmt.Println("test0")
		return 0
	}

	if s == "01" {
		//fmt.Println("test1")
		return 1
	}

	if s == "10" {
		//fmt.Println("test2")
		return 0
	}

	seconds := 0
	str := ""
	for {

		i := 0
		exist := true
		changed := false
		seconds++
		for {
			if len(s)-1 == i {
				str += string(s[i])
				break
			}
			if len(s)-1 < i {
				break
			}
			if s[i] == '0' && s[i+1] == '1' {
				str += "10"
				i = i + 2
				changed = true
				continue
			}

			str += string(s[i])
			i = i + 1
		}

		if !changed {
			seconds--
			break
		}

		for j := range str {
			if len(str)-1 == j+1 {
				break
			}

			if s[j] == '0' && s[j+1] == '1' {
				exist = true
			}
		}
		//fmt.Println(str)
		s = str
		str = ""

		if !exist {
			break
		}
	}

	return seconds
}
