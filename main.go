package main

import "fmt"

func main() {
	var result = add(2,3)
	println("add",result)
}

func add(a int, b int) int {
	var result int = a + b
	return result
}

func mul(a int, b int) int {
	var result int = a * b;
	return result
}

func sub(a int, b int) int {
	var result = a - b;
	return result
}

func div(a int, b int) int {
	var result = a / b
	return result
}
