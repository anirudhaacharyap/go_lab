package main
import "fmt"

func isSymmetric(matrix *[][]int, size *int) bool {
	for i := 0; i < *size; i++ {
		for j := 0; j < *size; j++ {
			if (*matrix)[i][j] != (*matrix)[j][i] {
				return false
			}
		}
	}
	return true
}
func main() {
	var size int
	fmt.Print("Enter the size of the square matrix: ")
	fmt.Scan(&size)
	
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	
	fmt.Println("Enter the matrix elements row by row:")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	
	if isSymmetric(&matrix, &size) {
		fmt.Println("The matrix is symmetric.")
	} else {
		fmt.Println("The matrix is not symmetric.")
	}
}
