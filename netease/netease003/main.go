package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		n, m, heights, query, err := input()

		if err == io.EOF {
			break
		}

		for _, q := range query {
			v, h := max(n, m, heights)

			if float64(q) < v {
				fmt.Printf("%.3f\n", h)
				continue
			}

			remain := float64(q)-v

			remainH := remain/(float64(n)*float64(m))

			fmt.Printf("%.3f\n", remainH+h)
		}
	}
}

func max(n, m int, heights []int) (float64, float64) {
	total := float64(0)
	hMax := 0

	for _, h := range heights {
		hMax = maxInt(hMax, h)
		total += float64(h)/3
	}

	fill := float64(hMax)*float64(n)*float64(m)

	return fill-total, float64(hMax)
}

func maxInt(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func output(n, m int, heights, query []int) {
	fmt.Printf("n %d, m %d\n", n,m)

	fmt.Printf("heights %v\n", heights)

	fmt.Printf("query %v\n", query)
}

func input() (int, int, []int, []int, error) {
	var n, m int

	_, err := fmt.Scan(&n, &m)

	if err != nil {
		return 0, 0, nil, nil, err
	}

	count := n*m
	heights := make([]int, count)
	for i:=0; i<count; i++ {
		_, _ = fmt.Scan(&heights[i])
	}

	_, _ = fmt.Scan(&count)

	query := make([]int, count)

	for i:=0; i<count; i++ {
		_, _ = fmt.Scan(&query[i])
	}

	return n, m, heights, query, nil
}
