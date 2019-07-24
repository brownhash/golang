package main

import "fmt"

type person struct{
	first_name string
	second_name string
}

func main() {
	person1 := person{
		first_name: "Harshit",
		second_name: "Sharma",
	}
        person2 := person{
                first_name: "Harry",
                second_name: "Potter",
        }

	fmt.Println(person1)
	fmt.Println(person2)
}
