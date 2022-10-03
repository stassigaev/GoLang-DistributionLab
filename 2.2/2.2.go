package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var firstNumber float64
	firstNumber = GetAndConvert("Enter first number:")
	fmt.Println(firstNumber)

	var secondNumber float64
	secondNumber = GetAndConvert("Enter second number:")
	fmt.Println(secondNumber)

	resultSide := TriangleSideCalc(firstNumber, secondNumber)
	fmt.Print("Hypotenuse is " + strconv.FormatFloat(resultSide, 'g', 1, 64))

	resultSquare := TriangleSquareCalc(firstNumber, secondNumber)
	fmt.Print(" Square is " + strconv.FormatFloat(resultSquare, 'g', 1, 64))

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
func TriangleSideCalc(a float64, b float64) float64 {
	var result = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return result
}

func TriangleSquareCalc(a float64, b float64) float64 {
	var result = 0.5 * a * b
	return result
}
