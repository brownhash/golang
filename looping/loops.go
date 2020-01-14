package main

import "fmt"

func main() {
	var num int

	// for is the only loop in go
	// basic for loop
	for num = 0; num < 5; num++ {
		fmt.Println(num)
	}

	fmt.Println("---")

	// for loop as while
	for num < 10 {
		fmt.Println(num)
		num++
	}

	fmt.Println("---")

	// infinite loop
	for {
		fmt.Println(num)
		if num == 15 {
			break
		}
		num++
	}
}
