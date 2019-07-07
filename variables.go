package main

import "fmt"

func main() {
	fmt.Println("Printing the variables")

	// shorthand declaration of variables

	a :=10
	b := "golang"
	c := 4.76

	fmt.Printf("%v \n",a)
	fmt.Printf("%v \n",b)
	fmt.Printf("%v \n",c)

	//declare variables

	var m string
	m="Harshit Sharma"
	fmt.Printf("%s\n",m)

	//declaring many at once

	var mess,mess1 string= "Hello","Hi"
	var x,y,z int = 1,2,3
	fmt.Printf("%s\n",mess)
	fmt.Printf("%s\n",mess1)
	fmt.Printf("%d\n",x)
	fmt.Printf("%d\n",y)
	fmt.Printf("%d\n",z)

	//init shorthand

	p,q,r := 1 , true , "hello"
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)

	type hotdog int
	var s hotdog
	s=12
	var integer int
	integer = int(s)

	fmt.Println(integer)
	fmt.Println(s)
}
