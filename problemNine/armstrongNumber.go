package problemNine

import "fmt"

func ArmstrongNumber() {
	fmt.Println("--- It's an Armstrong Number Checker ---")

	var userInput int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&userInput)

	// Keep the original number safe
	original := userInput

	// Step 1: Count digits
	temp := userInput
	count := 0

	for temp > 0 {
		count++
		temp = temp / 10
	}

	// Step 2: Separate digits and calculate power
	temp = userInput
	total := 0

	for temp > 0 {
		digit := temp % 10

		// Calculate digit^count
		result := 1
		powerCount := count

		for powerCount > 0 {

			result = result * digit
			powerCount--
		}

		total = total + result

		temp = temp / 10
	}

	// Step 3: Compare
	if total == original {
		fmt.Println("Armstrong Number")
	} else {
		fmt.Println("Not an Armstrong Number")
	}
}
