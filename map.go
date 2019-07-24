package main

import "fmt"

func main() {
	//map definition
	m := map[string]int {
		"Harshit" : 21,
		"Harry" : 22,
	}
	
	fmt.Println(m)
	fmt.Println(m["Harshit"])
	fmt.Println("\nDeleting Harshit")
	delete(m , "Harshit")
	fmt.Println(m)
	
	
}
