package lc10

import (
	"fmt"
	"testing"
)

func isMatch(s string, p string) bool {
	contentHead := readStrSet(s)
	parserHead := generateStateMachine(p)

	return scanning(contentHead.next, parserHead.next)
}

const (
	start = iota
	oneContent
	allContent
	stop
)

type stateType int

type state struct {
	sType   stateType
	content int32
	next    *state
	rotating bool
}

// @input:  parser
// @return: the head of states
func generateStateMachine(parser string) *state {
	head := &state{sType: start, content: 0, next: nil}
	preState := head

	for _, p := range parser {
		if isRotating(p) {
			preState.rotating = true
			continue
		}

		// trying to generate the next state
		currentState := &state{}

		// the type of state and the content of it
		if isAllKinds(p) {
			currentState.sType = allContent
		} else {
			currentState.sType = oneContent
			currentState.content = p
		}

		preState.next = currentState
		preState = currentState
	}

	stopState := &state{sType: stop}
	preState.next = stopState
	return head
}

func isRotating(p int32) bool {
	// *: 42
	return p == 42
}

func isAllKinds(p int32) bool {
	// .: 46
	return p == 46
}

const (
	first = iota
	pace
	last
)

type contentType int

type content struct {
	cType contentType
	value int32
	next  *content
}

func readStrSet(str string) *content {
	// a~z: 97~122
	head := &content{cType: first}
	preContent := head
	for _, c := range str {
		curContent := &content{cType: pace, value: c, next: nil}
		preContent.next = curContent
		preContent = curContent
	}
	stopContent := &content{cType: last}
	preContent.next = stopContent
	return head
}

func scanning(cont *content, parser *state) bool {
	fmt.Printf("%+v  ", cont)
	fmt.Printf("%+v\n", parser)

	if cont.cType == last && parser.sType == stop {
		return true
	}

	if cont.cType != last && parser.sType == stop {
		return false
	}

	if parser.rotating {
		if cont.cType == last && parser.sType != stop {
			return scanning(cont, parser.next)
		}
		if cont.cType == pace && parser.sType == oneContent {
			if cont.value != parser.content {
				return scanning(cont, parser.next)
			}
		}
	} else {
		if cont.cType == last && parser.sType != stop {
			return false
		}
		if cont.cType == pace && parser.sType == oneContent {
			if cont.value != parser.content {
				return false
			}
		}
	}

	if parser.rotating {
		if cont.cType == last {
			return scanning(cont, parser.next)
		}
		return scanning(cont, parser.next) || scanning(cont.next, parser) || scanning(cont.next, parser.next)
	} else {
		return scanning(cont.next, parser.next)
	}
}

func TestLC10(t *testing.T) {
	str := "bbbba"
	h1:=contentList(str)

	parser := ".*a*a"
	h2:=parserPrinter(parser)

	println(scanning(h1.next, h2.next))
}

func contentList(str string) *content {
	head := readStrSet(str)
	curContent := head
	for {
		print(curContent)
		fmt.Printf("%+v\n", curContent)
		if curContent.next == nil {
			break
		}

		curContent = curContent.next
	}
	return head
}

func parserPrinter(parser string) *state {
	head := generateStateMachine(parser)
	curState := head
	for {
		print(curState)
		fmt.Printf(" %+v\n", curState)

		if curState.sType == stop {
			break
		}

		curState = curState.next
	}
	return head
}
