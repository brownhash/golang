package main

import "fmt"

func main(){
	
	x:=50
	
	if true {
		fmt.Println("True")
	}
	if false {
		fmt.Println("Flase")	
	}
	if !true {
		fmt.Println("Not True")
	}
	if !false {
		fmt.Println("Not False")	
	}
	if x <= 100 {
		fmt.Println("Less than 100")
	} else {
		fmt.Println("More Than 100")
	}
	if x <= 49 {
                fmt.Println("Less than 49")
        } else {
                fmt.Println("More Than 49")
        }
}
