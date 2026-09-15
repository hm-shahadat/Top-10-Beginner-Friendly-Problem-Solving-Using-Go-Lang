package problemSix

import "fmt"

func SumofDigits() {
	fmt.Println("---It's an Sum of Digits---")

	var forVal int
	fmt.Print("Please enter numbers for addition: ")
	fmt.Scan(&forVal)
	var final int = 0
	var res int

	for forVal > 0 {
		res = forVal % 10
		final = final + res
		forVal = forVal / 10

	}
	fmt.Println(final)
	SumofDigits()
}
