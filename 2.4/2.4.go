package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber float64
	firstNumber = GetAndConvert("Enter first number:")
	fmt.Println(firstNumber) //4

	var secondNumber float64
	secondNumber = GetAndConvert("Enter second number:")
	fmt.Println(secondNumber) //2

	resultSide := pipeWorkCalc(firstNumber, secondNumber)
	fmt.Print("Pipe will fill for " + strconv.FormatFloat(resultSide, 'g', 10, 64))
}

func GetAndConvert(numberTitle string) float64 {
	var strNumber string
	fmt.Print(numberTitle)
	fmt.Scan(&strNumber)
	fmt.Println()
	numberResultI64, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		fmt.Println("Invalid Number")
		return GetAndConvert(numberTitle)
	}
	return numberResultI64
}

func pipeWorkCalc(a float64, b float64) float64 {
	var result1 = (1/a + 1/b)
	var result = 1 / result1

	return result // 1.3333
}
