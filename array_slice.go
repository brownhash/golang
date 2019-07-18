package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5,6}
	
	fmt.Println(a[2:])
	fmt.Println(a[2:5])

	a = append(a, 123, 456)
	fmt.Println("Appended values - ",a)
	
	a = append(a[:2],a[4:]...)
	fmt.Println(a)
}
