package main

func facto(n int) int {
	var result = 1
	for i := 1; i < n; i++ {
		result *= i
	}
	return result
}
