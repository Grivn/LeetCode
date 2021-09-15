package main

import "fmt"

func main() {

	m := 3
	n := 7

	uniquePaths(m, n)
}

type graph [][]int

type pos struct {
	x int
	y int
}

func uniquePaths(m int, n int) int {
	g := generateGraph(m, n)

	for _, raw := range g {
		fmt.Println(raw)
	}

	return 0
}

func generateGraph(m, n int) graph {
	g := make([][]int, m)

	for index := range g {
		raw := make([]int, n)
		g[index] = raw
	}

	g[m-1][n-1] = 1

	return g
}
