package problemTwo

import "fmt"

func Fseries() {
	fmt.Println("---It's an fibonacci series---")

	var input int
	fmt.Println("please enter a integer number how many times you want to run Fibonacci Series: ")
	fmt.Scan(&input)

	var a int = 0
	var b int = 1
	fmt.Print(a)
	fmt.Print(b)
	var c int

	for i := 2; i < input; i++ {

		c = a + b
		a = b
		b = c

		fmt.Print(" ", c)

	}

}
