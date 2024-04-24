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
}
