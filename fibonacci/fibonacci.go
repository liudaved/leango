package main

import "fmt"

func fibonacci(length int) int {
	if length < 2 {
		return 1
	}

	return fibonacci(length-1) + fibonacci(length-2)
}

func main() {
	for i := 0; i < 10; i++ {
		nums := fibonacci(i)
		fmt.Println(nums)
	}
}
