package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[3] = 45
	fmt.Println(arr)
	fmt.Println(len(arr))

	arr2 := [5]int{1,2,3,4,5}
	fmt.Println(arr2)
	
	for i:=0; i<len(arr); i++ {
		fmt.Printf("Index %d = %d\n",i,arr[i])
	}

	for i,v:=range arr {
		fmt.Println(i,v)
	}
}
