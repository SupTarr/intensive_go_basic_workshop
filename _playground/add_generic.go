package main

// Generic type T as int or float
func add[T int | float64](a, b T) T {
	return a + b
}
