package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("All args, array:")
	fmt.Println(os.Args)
	fmt.Println("\nAll args:")
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}

// running the binary -
// $ go build cmd_args.go
// $ ./build_name arg1 arg2
