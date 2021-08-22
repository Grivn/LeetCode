package main

import "fmt"

func main() {
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word := "ABCCED"

	flag := exist(board, word)

	println(flag)
}

type comeFrom int

const (
	null = iota
	left
	right
	up
	down
)

type pos struct {
	x int
	y int
}

// col[larger ---> downer], raw[larger ---> righter]
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return len(word) == 0
	}

	for i, raw := range board {
		for j := range raw {
			p := pos{x: i, y: j}
			flag := search(p, p, board, word)
			if flag {
				return true
			}
		}
	}

	return false
}

func search(from pos, p pos, board [][]byte, word string) bool {
	if len(word) == 0 {
		fmt.Printf("from %v, pos %v, word %v, res %v\n", from, p, word, true)
		return true
	}

	if word[0] != read(board, p) {
		fmt.Printf("from %v, pos %v, word %v, res %v\n", from, p, word, false)
		return false
	}

	if len(word) == 1 {
		fmt.Printf("from %v, pos %v, word %v, res %v\n", from, p, word, true)
		return true
	}

	word = word[1:]

	var next []pos
	next = append(next, lIndex(p))
	next = append(next, rIndex(p))
	next = append(next, uIndex(p))
	next = append(next, dIndex(p))

	for _, nextPos := range next {
		if !notEqual(nextPos, p) {
			continue
		}

		if search(p, nextPos, board, word) {
			fmt.Printf("from %v, pos %v, word %v, res %v\n", from, p, word, true)
			return true
		}
	}

	fmt.Printf("from %v, pos %v, word %v, res %v\n", from, p, word, false)
	return false
}

func mSchema(board [][]byte) (int, int) {
	col := len(board)

	raw := len(board[0])

	return col, raw
}

func read(board [][]byte, p pos) byte {
	return board[p.x][p.y]
}

func notEqual(p1, p2 pos) bool {
	return p1.x != p2.x && p1.y != p2.y
}

func uIndex(p pos) pos {
	return pos{x: p.x-1, y: p.y}
}

func dIndex(p pos) pos {
	return pos{x: p.x+1, y: p.y}
}

func lIndex(p pos) pos {
	return pos{x: p.x, y: p.y-1}
}

func rIndex(p pos) pos {
	return pos{x: p.x, y: p.y+1}
}
