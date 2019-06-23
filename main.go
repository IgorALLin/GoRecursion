package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var number float64;

func isPrime(devider float64) string {
	if devider > number / 2 {
		return "Tris number is prime"
	}

	divisionResult := number / devider;

	if divisionResult == math.Trunc(divisionResult) {
		return "Tris number is not prime"
	}

	return isPrime(devider + 1)
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	input, _ := reader.ReadString('\n')

	//Maybe there is exist better way to delete those symbols from the end of command line string?
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	//Parse input string to float64
	inputFloat, err := strconv.ParseFloat(input, 64)

	if err != nil{
		fmt.Println(err)

		return
	}

	number = inputFloat

	fmt.Println(isPrime(2))
}
