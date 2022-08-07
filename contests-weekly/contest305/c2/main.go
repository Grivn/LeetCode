package main

import "fmt"

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}
	re := []int{4, 5}
	ret := reachableNodes(7, edges, re)
	fmt.Println(ret)
}
func reachableNodes(n int, edges [][]int, restricted []int) int {
	g := make(map[int]map[int]bool)

	for _, edge := range edges {
		sub1, ok := g[edge[0]]
		if !ok {
			g[edge[0]] = make(map[int]bool)
			sub1 = g[edge[0]]
		}
		sub1[edge[1]] = true

		sub2, ok := g[edge[1]]
		if !ok {
			g[edge[1]] = make(map[int]bool)
			sub2 = g[edge[1]]
		}
		sub2[edge[0]] = true
	}

	r := make(map[int]bool)
	for _, v := range restricted {
		r[v] = true
	}

	visited := make(map[int]bool)
	worker(0, g, visited, r)
	return len(visited)
}

func worker(here int, g map[int]map[int]bool, visited map[int]bool, r map[int]bool) {
	if r[here] {
		return
	}

	//fmt.Println("visit:", here)
	visited[here] = true
	sub, ok := g[here]
	if !ok {
		return
	}

	for v := range sub {
		if v == here {
			continue
		}
		if visited[v] {
			continue
		}
		worker(v, g, visited, r)
	}
}
