package problemFour

import "fmt"

func PrimeNumber() {

	fmt.Println("---It's an checker of prime or not prime---")

	var n int
	fmt.Print("Please enter a number which you want to check prime or not prime: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("It's not a prime number")
		return
	}

	if n == 2 || n == 3 || n == 5 || n == 7 {
		fmt.Println("It's a prime number")
	}

	if n%2 == 0 || n%3 == 0 || n%5 == 0 || n%7 == 0 {
		fmt.Println("It's not a prime number")
	} else {
		fmt.Println("It's a prime number")
	}

	PrimeNumber()

}
