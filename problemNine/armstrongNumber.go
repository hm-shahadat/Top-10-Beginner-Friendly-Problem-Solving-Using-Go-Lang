package problemNine

import "fmt"

func ArmstrongNumber() {
	fmt.Println("--- It's an Armstrong Number Checker ---")

	var userInput int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&userInput)

	original := userInput

	temp := userInput
	count := 0

	for temp > 0 {
		count++
		temp = temp / 10
	}

	temp = userInput
	total := 0

	for temp > 0 {
		digit := temp % 10

		result := 1
		powerCount := count

		for powerCount > 0 {

			result = result * digit
			powerCount--
		}

		total = total + result

		temp = temp / 10
	}

	if total == original {
		fmt.Println("Armstrong Number")
	} else {
		fmt.Println("Not an Armstrong Number")
	}
}
