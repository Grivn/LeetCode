package main

func main() {

}
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)

	for i := range res {
		if i == 0 {
			res[i] = []int{1}
			continue
		}

		if i == 1 {
			res[i] = []int{1, 1}
			continue
		}

		val := make([]int, i+1)
		for j := range val {
			if j == 0 {
				val[j] = 1
				continue
			}
			if j == len(val)-1 {
				val[j] = 1
				continue
			}
			val[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = append(res[i], val...)
	}
	return res
}
