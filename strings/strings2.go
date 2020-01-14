package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// initialising a map to count the occurance of a string
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// incrementing the value of a string in the map
	for i := 0; i < 10; i++ {
		input.Scan()
		counts[input.Text()]++
		fmt.Println(counts)
	}

	// listing the duplicate entries
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
