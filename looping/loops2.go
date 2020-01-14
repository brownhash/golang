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
}
