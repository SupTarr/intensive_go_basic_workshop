package main

func sumOfEven(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}
