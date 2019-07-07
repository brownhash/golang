package main

import "fmt"

func main() {
	fmt.Println("Lets start using function:\n")
	a:=0
	fmt.Println("Enter your age-->")
	fmt.Scanf("%d",&a)
	age(a)
}

func age(n int) {
	if(n>=18){
		fmt.Println("Adult")
	}else{
		fmt.Println("Not Adult")
	}
}
