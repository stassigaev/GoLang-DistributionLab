package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstPipe float64
	firstPipe = GetAndConvert("Enter first pipe:")
	fmt.Println(firstPipe) //4

	var secondPipe float64
	secondPipe = GetAndConvert("Enter second pipe:")
	fmt.Println(secondPipe) //2

	resultSide := pipeWorkCalc(firstPipe, secondPipe)
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
