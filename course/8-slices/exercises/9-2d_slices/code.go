package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for row:= 0; row<rows; row++{
		rows_slice := []int{}
		for col:=0; col<cols;col++{
			rows_slice = append(rows_slice,row*col)
		}
		matrix = append(matrix,rows_slice)
	}
	return matrix
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
