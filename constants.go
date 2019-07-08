package main

import "fmt"

//way1 to declare constants

const a = 23
const b = 23.54
const c = "Hello World"

//way2 to declare constants

const(
	d = 46
	e = 46.78
	f = "Hi World"
)

func main() {
	fmt.Println(a)
        fmt.Println(b)
        fmt.Println(c)
        fmt.Println(d)
        fmt.Println(e)
        fmt.Println(f)

}
