package problemSeven

import "fmt"

func ReverseNumber() {
	fmt.Println("---It's an Reverse Number---")

	var n int
	fmt.Print("Please enter numbers for addition: ")
	fmt.Scan(&n)
	var res int

	for n > 0 {
		res = n % 10 // 123---res=3// 12---2
		n = n / 10   // 123 --- 12 // 1
		fmt.Print(res)

	}
}
