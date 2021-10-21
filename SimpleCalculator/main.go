package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv" //string cnversion
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduceti primul nr: ")

	numInput, _ := reader.ReadString('\n')

	firstNumber, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		panic("Please insert a number")
	}

	fmt.Print("Introduceti al doilea nr: ")
	numInput, _ = reader.ReadString('\n')

	secondNumber, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		panic("Please insert a number")
	}

	sum := firstNumber + secondNumber

	sum = math.Round(sum*100) / 100

	fmt.Print("Suma numerelor introduse este: ", sum)

}
