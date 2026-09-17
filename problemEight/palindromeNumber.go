package problemEight

import "fmt"

func PalindromeNumber() {
	fmt.Println("---It's an Palindrome Number---")

	var n int
	fmt.Print("Please enter numbers for Palindrome Number check: ")
	fmt.Scan(&n)
	var f int
	f = n
	var res int
	var reverse int

	for n > 0 {
		res = n % 10
		n = n / 10
		reverse = reverse*10 + res
	}
	fmt.Print(reverse)
	if reverse == f {
		fmt.Print("---Palindrome Number")
	} else {
		fmt.Print("---Not Palindrome Number")
	}
}
