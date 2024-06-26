package main

import (
	"flag"
	"fmt"
	"log/slog"
	"strconv"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println("addInt:", addInt(1, 2))
	fmt.Println("addFloat:", addFloat(1.2, 1.3))
	fmt.Println("add:", add(1, 2))
	fmt.Println("addStringNumber:", addStringNumber("1", "2"))
	fmt.Println("addStringNumber:", addStringNumber("1", "a"))

	flag.Parse()
	args := flag.Args()
	price, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		slog.Info(err.Error())
	}

	interestRate, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		slog.Info(err.Error())
	}

	year, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		slog.Info(err.Error())
	}
	fmt.Println("price:", price)
	fmt.Println("interestRate:", interestRate)
	fmt.Println("year:", year)
	fmt.Println("monthlyInstallmemt:", monthlyInstallmemt(price, interestRate, year))

	var numbers = []int{11, 8, 5, 1, 4, 10}
	fmt.Println("sumOfEven:", sumOfEven(numbers))

	var numbersArr = [4]int{1, 2, 3, 4}
	fmt.Println("reverseArray:", reverseArray(numbersArr))

	var numbersSlice = []int{4, 5, 6, 7, 11, 21, 37}
	fmt.Println("primes:", primes(numbersSlice))
}
