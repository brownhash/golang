package main

import "fmt"

func main() {
	arr1 := []string {"harshit","sharma","harry","potter"}
	fmt.Println(arr1)

        arr2 := []string {"harshit1","sharma1","harry1","potter1"}
	fmt.Println(arr2)
	
	fmt.Println("Multi dimention slice")
	arr3 := [][]string {arr1,arr2}
	fmt.Println(arr3)
}
