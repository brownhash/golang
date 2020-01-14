package main

import (
	"fmt"
	"os"
)

func main() {
	// using for as for-each
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	// in above loop _ is used as a blank identifier
	// to ignore the index
	for index, arg := range os.Args {
		fmt.Println(arg, index)
	}
}
