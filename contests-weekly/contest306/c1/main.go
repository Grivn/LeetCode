package main

func main() {

}
func largestLocal(grid [][]int) [][]int {
	var res [][]int
	for i, row := range grid {
		var subRow []int
		for j := range row {
			sub := []int{
				grid[i][j], grid[i][j+1], grid[i][j+2],
				grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2],
				grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2],
			}
			val := largestValue(sub)
			subRow = append(subRow, val)
			if j+3 == len(row) {
				break
			}
		}
		res = append(res, subRow)
		if i+3 == len(grid) {
			break
		}
	}

	return res
}

func largestValue(grid []int) int {
	num := grid[0]
	for _, val := range grid {
		if val > num {
			num = val
		}
	}
	return num
}
