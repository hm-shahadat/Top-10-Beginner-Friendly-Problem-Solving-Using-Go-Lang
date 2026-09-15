package problemThree

import "fmt"

func Factorial() {

	fmt.Println("---It's an Factorial---")

	var input int
	fmt.Println("Please enter a number you want see which number factorial value: ")
	fmt.Scan(&input)
	var output int = 1
	if input == 0 {
		fmt.Println("Please do not use 0 and try again")
		Factorial()
	} else {
		for i := 1; i <= input; i++ {
			output = i * output
		}
		fmt.Println(output)
		Factorial()
	}

}
