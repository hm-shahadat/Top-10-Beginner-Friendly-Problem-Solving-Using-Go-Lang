package problemSeven

import "fmt"

func ReverseNumber() {
	fmt.Println("---It's an Reverse Number---")

	var n int
	fmt.Print("Please enter numbers for addition: ")
	fmt.Scan(&n)
	var res int

	for n > 0 {
		res = n % 10
		n = n / 10
		fmt.Print(res)

	}
}
