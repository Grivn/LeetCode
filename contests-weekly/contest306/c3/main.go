package main

import (
	"fmt"
	"strconv"
)

func main() {
	//set := worker("IIIDIDDD")
	//fmt.Println(set)
	//fmt.Println(format(set))
	h := newHandler("DDDIII")
	h.processor()
	fmt.Println(h.res)
}
func smallestNumber(pattern string) string {
	h := newHandler(pattern)
	h.processor()
	str := ""
	for _, v := range h.res {
		str += strconv.Itoa(v)
	}
	return str
}

type handler struct {
	used    map[int]bool
	pattern string
	res     []int
}

func newHandler(pattern string) *handler {
	return &handler{
		used:    make(map[int]bool),
		pattern: pattern + "I",
		res:     make([]int, len(pattern)+1),
	}
}

func (h *handler) processor() {
	if len(h.pattern) == 0 {
		return
	}

	if len(h.pattern) == 1 {
		h.res = []int{1}
		return
	}

	h.res[0] = 1
	h.used[1] = true

	for i := range h.pattern {
		if i == 0 {
			continue
		}
		for {
			if h.findValue(i) {
				break
			}
			h.reFormat(i - 1)
		}
	}
}

func (h *handler) reFormat(idx int) {
	for i := idx; i >= 0; i-- {
		if h.pattern[i] == 'I' {
			break
		}

		if h.pattern[i] == 'D' {
			h.res[i] = h.res[i] + 1
			h.used[h.res[i]] = true
		}
	}

	h.used = make(map[int]bool)
	for _, v := range h.res {
		h.used[v] = true
	}
}

func (h *handler) findValue(idx int) bool {
	if h.pattern[idx-1] == 'I' {
		val := h.res[idx-1] + 1
		for {
			if !h.used[val] {
				h.res[idx] = val
				h.used[val] = true
				break
			}
			val++
		}
		return true
	}

	val := h.res[idx-1] - 1
	if val == 0 {
		return false
	}

	for {
		if !h.used[val] {
			h.res[idx] = val
			h.used[val] = true
			break
		}
		val--
		if val == 0 {
			return false
		}
	}
	return true
}
