package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	size := len(matrix) * len(matrix[0])
	for len(res) < size {
		for i := left; i <= right && len(res) < size; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom && len(res) < size; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		for i := right; i >= left && len(res) < size; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top && len(res) < size; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
