package main

import "fmt"

func main(){
	switch {
		case false:
			fmt.Println("Should not print")
		case (2==3):
			fmt.Println("Should not print")
		case true:
			fmt.Println("Prints")
		case (2==2):
			fmt.Println("Also true but do not prints")
	}
	
	//using fall through	
	switch {
                case false:
                        fmt.Println("Should not print")
                case (2==3):
                        fmt.Println("Should not print")
                case true:
                        fmt.Println("Prints")
                        fallthrough
                case (2==2):
                        fmt.Println("Prints this time")
			fallthrough
		case (2==5):
			fmt.Println("False but prints")
			fallthrough
		case (5==6):
			fmt.Println("Final statement")
			//fallthrough cannot be applied to the final statement
        }

	//using default case
	switch {
		case false:
			fmt.Println("False case")
		case (2==4):
			fmt.Println("False case")
		default:
			fmt.Println("Default case")	
	}
}
