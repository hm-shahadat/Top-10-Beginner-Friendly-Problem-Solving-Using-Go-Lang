package main

import (
	"fmt"

	problemfive "bdtask.com/problemSolving/problemFive"
	"bdtask.com/problemSolving/problemFour"
	"bdtask.com/problemSolving/problemOne"
	"bdtask.com/problemSolving/problemThree"
	"bdtask.com/problemSolving/problemTwo"
)

func main() {
	var problemNumber int
	fmt.Println("Please tell me which program you want to run (1 to 10) \n 1.Odd Even \n 2.Fibonacci Series \n 3.Factorial \n 4.Prime Number \n 5.Multiplication Table")
	fmt.Scan(&problemNumber)

	switch problemNumber {

	case 1:
		problemOne.CheckOddEven()

	case 2:
		problemTwo.Fseries()

	case 3:
		problemThree.Factorial()
	case 4:
		problemFour.PrimeNumber()
	case 5:
		problemfive.MultiplicationTable()

	}
}
