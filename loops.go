package main

import "fmt"

func main(){
	
	//loop 1
	i:=0
	for i<=3 {
		fmt.Println(i)
		i++
	}

	//loop 2
	for j:=0; j<=3; j++ {
		fmt.Println(j)
	}

	//loop3
	for {
		fmt.Println("loop")
		break
	}
}
