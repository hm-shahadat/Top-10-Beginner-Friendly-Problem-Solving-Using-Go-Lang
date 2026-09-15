package problemOne

import "fmt"

func CheckOddEven() {

	fmt.Println("---It's a Even Odd checker---")
	var number int
	fmt.Print("Please enter a Number how much time you want to test the : ")
	fmt.Scan(&number)

	if number >= 0 {
		if number%2 == 0 {
			fmt.Println("It's an even number")
		} else {
			fmt.Println("It's an odd number")
		}

		CheckOddEven()
	} else {
		fmt.Println("Thank you for using our system")
	}
}
