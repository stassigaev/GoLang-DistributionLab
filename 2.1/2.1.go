package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber int64
	firstNumber = GetAndConvert("Enter first number:")
	fmt.Println(firstNumber)

	var secondNumber int64
	secondNumber = GetAndConvert("Enter second number:")
	fmt.Println(secondNumber)

	var thirdNumber int64
	thirdNumber = GetAndConvert("Enter third number:")
	fmt.Println(thirdNumber)

	var foursNumber int64
	foursNumber = GetAndConvert("Enter fours number:")
	fmt.Println(foursNumber)

	result := SumNumbers(firstNumber, secondNumber, thirdNumber, foursNumber)
	fmt.Println("Result" + strconv.FormatInt(result, 10))

}

func GetAndConvert(numberTitle string) int64 {
	var strNumber string
	fmt.Print(numberTitle)
	fmt.Scan(&strNumber)
	fmt.Println()
	numberResultI64, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		fmt.Println("Invalid Number")
		return GetAndConvert(numberTitle)
	}
	return numberResultI64
}

func SumNumbers(a int64, b int64, c int64, d int64) int64 {
	var result = a + b + c + d
	return result
}

// func sumNumbers(a int64, b int64, c int64, d int64) int64 {
// 	var sum = a + b + c + d
// 	return sum
// }
