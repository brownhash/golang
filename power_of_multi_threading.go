package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	operationDone := make(chan bool)
	now := time.Now()
	sum := 0
	for i := 0; i <= 10000000; i++ {
		sum = sum + i
	}
	fmt.Println("####", sum)

	go func() {
		sum2 := 0
		for j := 0; j <= 10000000; j++ {
			sum2 = sum2 + j
		}
		fmt.Println(">>>>", sum2)
		operationDone <- true
	}()
	<-operationDone
	end := time.Now()
	fmt.Println(end.Sub(now))
}

/*
Results in go using multi threading -
#### 50000005000000
>>>> 50000005000000
8.384325 ms

Results in Python without multi threading -
50000005000000
50000005000000
2839.966058 ms

Golang is 338x faster than python in this comparision
*/
