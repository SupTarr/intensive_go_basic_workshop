package main

// func reverseArray(numbers [4]int) [4]int {
// 	result := [4]int{}
// 	for i := len(numbers) - 1; i >= 0; i-- {
// 		result[(len(numbers)-1)-i] = numbers[i]
// 	}

// 	return result
// }

func reverseArray(numbers [4]int) [4]int {
	result := [4]int{}
	for i := 0; i < len(numbers); i++ {
		result[(len(numbers)-1)-i] = numbers[i]
	}

	return result
}
