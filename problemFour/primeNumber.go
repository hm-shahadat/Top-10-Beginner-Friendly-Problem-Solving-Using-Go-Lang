package problemFour

import "fmt"

func PrimeNumber() {
	fmt.Println("---It's an checker of prime or not prime---")

	// var count int
	// fmt.Print("Please enter a number how many times you want to run the program: ")
	// fmt.Scan(&count)

	// for i := 0; i < count; i++ {

	var in int
	fmt.Print("Please enter a number which you want to check prime or not prime: ")
	fmt.Scan(&in)

	if in > 1 && in/in == 0 {
		fmt.Println("It's a prime number")
	} else if in == 1 {
		fmt.Println("It's not a prime number")
	} else {
		fmt.Println("It's not a prime number")
	}
	PrimeNumber()
}

// }
