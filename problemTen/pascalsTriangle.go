package problemTen

import "fmt"

func PascalTriangle() {
	fmt.Println("--- Pascal's Triangle ---")

	var n int
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		for space := 0; space < n-i-1; space++ {
			fmt.Print(" ")
		}

		value := 1

		for j := 0; j <= i; j++ {
			fmt.Print(value, " ")

			value = value * (i - j) / (j + 1)
		}

		fmt.Println()
	}
}
