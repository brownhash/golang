package main

import "fmt"

type person struct{
	f_name string
	s_name string
	age int
}

type is_married struct{
	person
	status bool
}

func main() {
	data := is_married{
		person: person {
			f_name: "Harshit",
			s_name: "Sharma",
			age: 21,
		},
		status: false,
	}
	fmt.Println(data)

        data2 := is_married{
                person: person {
                        f_name: "Harry",
                        s_name: "potter",
                        age: 22,
                },
                status: false,
        }
        fmt.Println(data2)
	fmt.Println(data.f_name,data.s_name,data.age)
	fmt.Println(data2.f_name,data2.s_name,data2.age)
}
