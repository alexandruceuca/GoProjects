package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv" //string cnversion
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

//reading the numbers from the user
//this function can be change for 1 number
func reading(firstNumber, secondNumber float64) (float64, float64) {

	fmt.Print("First number: ")
	numInput, _ := reader.ReadString('\n')

	firstNumber, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		panic("Please insert a number")
	}

	fmt.Print("Second number: ")
	numInput, _ = reader.ReadString('\n')

	secondNumber, err = strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		panic("Please insert a number")
	}
	return firstNumber, secondNumber
}

//readin the opereantion
func readSign(signInput string) string {

	fmt.Print("Choose an operation sign +,-,*,/: ")
	signInput, _ = reader.ReadString('\n')
	signInput = strings.TrimSpace(signInput) //trimming the spaces
	return signInput
}

func main() {

	var firstNumber, secondNumber float64
	firstNumber, secondNumber = reading(firstNumber, secondNumber)

	var operation float64

	var signInput string
	signInput = readSign(signInput)
	//	fmt.Print(signInput)
	switch signInput {
	case "+":
		operation = firstNumber + secondNumber //the operations here can be separete functions

	case "-":
		operation = firstNumber - secondNumber

	case "*":
		operation = firstNumber * secondNumber

	case "/":
		operation = firstNumber / secondNumber
	default:
		panic("Invalid operantion")

	}
	operation = math.Round(operation*100) / 100

	fmt.Print(firstNumber, " ", signInput, " ", secondNumber, "=", operation)

}
