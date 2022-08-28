package main

import "fmt"

func main() {
	fmt.Println(minimumRecolors("BWWWBB", 6))
}
func minimumRecolors(blocks string, k int) int {
	min := len(blocks)
	for i := range blocks {
		if len(blocks)+1 == i+k {
			break
		}
		count := 0
		for _, v := range blocks[i : i+k] {
			if v == 'W' {
				count++
			}
		}
		if count < min {
			min = count
		}
	}
	return min
}
