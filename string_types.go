package main

import "fmt"

func main() {
	s := "Hello World"
	s1 := `Hello 
		World`
	fmt.Println(s)
	fmt.Println(s1)
	
	bs := []byte(s)
	bs1 := []byte(s1)
	
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)
	fmt.Println(bs1)
	fmt.Printf("%T\n",bs1)

	for i:=0; i<len(s); i++{
		fmt.Printf("%#U",s[i])
	}
}
