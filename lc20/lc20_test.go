package lc20

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	top := &stackValue{}
	for _, char := range s {
		if compare(top, char) {
			top = pop(top)
		} else {
			top = push(top, char)
		}
		fmt.Printf("char: %d, top: %+v\n", char, top)
	}

	return finished(top)
}

type stackValue struct {
	value int32
	next  *stackValue
}

func push(top *stackValue, char int32) *stackValue {
	return &stackValue{value: char, next: top}
}

func compare(top *stackValue, char int32) bool {
	if top == nil {
		return false
	}

	return (char == 41 && top.value == 40) || (char == 93 && top.value == 91) || (char == 125 && top.value == 123)
}

func pop(top *stackValue) *stackValue {
	return top.next
}

func finished(top *stackValue) bool {
	return top.next == nil
}

func TestLC20(t *testing.T) {
	println(isValid("()[]{}"))
}
