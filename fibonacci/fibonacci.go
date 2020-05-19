package main

import "fmt"

// 递归实现
func fibonacciRe(length int) int {
	if length < 0 {
		return -1
	}
	if length == 0 {
		return 0
	}
	if length <= 2 {
		return 1
	}

	return fibonacciRe(length-1) + fibonacciRe(length-2)
}

// 循环实现
func fibonacciFor(length int) int {
	if length < 0 {
		return -1
	}
	if length == 0 {
		return 0
	}
	if length <= 2 {
		return 1
	}
	a, b := 1, 1
	result := 0
	for i := 3; i <= length; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}

// 闭包实现
func fibonacciCall(length int) int {
	if length<0 {
		return -1
	}
	f := f()
	result := 0
	for i := 0; i < length; i++ {
		result = f()
	}
	return result
}

func f() func() int{
	a, b := 0, 1
	return func() int{
		a, b = b, a+b
		return a
	}
}


func main() {
	var nums int
	for i := 0; i < 10; i++ {
		nums = fibonacciRe(i)
		fmt.Println(nums)
		nums = fibonacciFor(i)
		fmt.Println(nums)
		nums = fibonacciCall(i)
		fmt.Println(nums)
	}
}
