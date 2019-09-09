package main

import(
	"fmt"
	"os"
)

func main()  {
	args := os.Args
	if args[1] == "-h" || args[1] == "--help"{
		fmt.Println("Using commandline arguments in Golang")
	} else{
		fmt.Println("Not a recognised action")
	}
}