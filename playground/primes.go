package main

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primes(numbers []int) []int {
	result := []int{}
	for i := 0; i < len(numbers); i++ {
		if isPrime(numbers[i]) {
			result = append(result, numbers[i])
		}
	}

	return result
}
