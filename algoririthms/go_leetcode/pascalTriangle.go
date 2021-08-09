package go_leetcode

import "fmt"

// Pascal's Triangle
func main() {
	var n int = 6
	fmt.Println(pascalTriangle(n))
}

func pascalTriangle(numRows int) [][]int {
	a := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		a[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || i == j {
				a[i][j] = 1
			} else {
				a[i][j] = a[i-1][j] + a[i-1][j-1]
			}
		}
	}
	return a
}
