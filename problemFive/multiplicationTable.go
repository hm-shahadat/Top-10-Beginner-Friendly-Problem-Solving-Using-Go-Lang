package problemfive

import "fmt"

func MultiplicationTable() {

	fmt.Println("---It's an Multiplication Table---")
	var u1 int
	fmt.Print("Please enter a number what number you have to see the Multiplication Table: ")
	fmt.Scan(&u1)

	var u2 int
	fmt.Print("Please enter a number how many times you want to see the Multiplication: ")
	fmt.Scan(&u2)

	for i := 1; i <= u2; i++ {
		fmt.Println(i * u1)

	}

}
