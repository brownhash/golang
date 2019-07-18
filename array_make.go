package main

import "fmt"

func main() {
	a := make([]int , 10 , 15)
	fmt.Println("Length - ",len(a))
	fmt.Println("Capacity - ",cap(a))

	a = append(a,1,2,3,4,5)
	fmt.Println(a)
}
