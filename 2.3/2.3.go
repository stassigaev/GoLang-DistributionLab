package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber int64
	firstNumber = GetAndConvert("Enter first number:")
	fmt.Println(firstNumber) //2

	var secondNumber int64
	secondNumber = GetAndConvert("Enter second number:")
	fmt.Println(secondNumber) //3

	var thirdNumber int64
	thirdNumber = GetAndConvert("Enter third number:")
	fmt.Println(thirdNumber) //4

	var a, b, c int64 = thirdNumber, secondNumber, firstNumber
	fmt.Println(a, b, c) // 432
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

// func reverseNumber(num int) int {

// 	res := 0
// 	for num>0 {
// 	   remainder := num % 10
// 	   res = (res * 10) + remainder
// 	   num /= 10
// 	}
// 	return res
//  }
