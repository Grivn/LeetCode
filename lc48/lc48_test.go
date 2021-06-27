package lc48

import (
	"fmt"
	"testing"
)

func rotate(matrix [][]int)  {
	n := len(matrix)
	newMatrix := make([][]int, n)
	for i:=0; i<n; i++ {
		newMatrix[i] = make([]int, n)
		for j:=0; j<n; j++ {
			newMatrix[i][j] = matrix[n-1-j][i]
		}
	}
	copy(matrix, newMatrix)
}

func TestLC48(t *testing.T) {
	matrix := [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}
	rotate(matrix)
	fmt.Println(matrix)
}
