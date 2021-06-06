package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pre panic execution")
	value1 := safeDivision(2, 0)
	fmt.Println("Post panic execution, -> ", value1)

	fmt.Println("Pre valid execution")
	value2 := safeDivision(2, 1)
	fmt.Println("Post valid execution, value -> ", value2)
}

/*
This function is supposed to handle the errors
and return a safe value in case of an unusual situation
*/
func safeDivision(x, y float64) float64 {
	var returnValue float64

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occured:", err)
			fmt.Println("Returning safe values")

			returnValue = 0
		}
	}()
	checkForError(y)

	returnValue = x/y

	return returnValue
}

/*
This function is supposed to replicate the mathematical error
`DIVIDE BY ZERO`
and initiates a panic if the denominator is 0
*/
func checkForError(y float64) {
	if y == 0 {
		panic("Divident cannot be 0! Divide by 0 error.")
	}
}
