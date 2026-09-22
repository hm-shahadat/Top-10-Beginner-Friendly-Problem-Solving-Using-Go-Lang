package main

import (
	"fmt"

	"bdtask.com/problemSolving/problemEight"
	"bdtask.com/problemSolving/problemFive"
	"bdtask.com/problemSolving/problemFour"
	"bdtask.com/problemSolving/problemNine"
	"bdtask.com/problemSolving/problemOne"
	"bdtask.com/problemSolving/problemSeven"
	"bdtask.com/problemSolving/problemSix"
	"bdtask.com/problemSolving/problemTen"
	"bdtask.com/problemSolving/problemThree"
	"bdtask.com/problemSolving/problemTwo"
)

func main() {
	var problemNumber int
	fmt.Println("Please tell me which program you want to run (1 to 10) \n 1.Odd Even \n 2.Fibonacci Series \n 3.Factorial \n 4.Prime Number \n 5.Multiplication Table \n 6.Sum of Digits \n 7.Reverse Number \n 8.Palindrome Number \n 9.Armstrong Number \n 10.Pascal's Triangle ")
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
		problemFive.MultiplicationTable()
	case 6:
		problemSix.SumofDigits()
	case 7:
		problemSeven.ReverseNumber()
	case 8:
		problemEight.PalindromeNumber()
	case 9:
		problemNine.ArmstrongNumber()
	case 10:
		problemTen.PascalTriangle()

	}

}
