package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	runtime.GOMAXPROCS(2)
	test1()
	test2()

	operationDone := make(chan bool) // creating a channel

	go func(){
		fmt.Println(">>>1")
		fmt.Println(">>>1")
		fmt.Println(">>>1")
		fmt.Println(">>>1")
		operationDone <- true // marking a task as done
	}()
	test1()
	test2()

	go func(){
		fmt.Println(">>>2")
		fmt.Println(">>>2")
		fmt.Println(">>>2")
		fmt.Println(">>>2")
		operationDone <- true
	}()
	<-operationDone // channel end

	test1()
	test2()
}

func test1(){
	fmt.Println("out 1")
	fmt.Println(time.Now())
}
func test2(){
	fmt.Println("out 2")
	fmt.Println(time.Now())
}

/*
One of the execution results -

out 1
2019-09-09 15:32:05.365835 +0530 IST m=+0.000224233
out 2
2019-09-09 15:32:05.36599 +0530 IST m=+0.000378887
out 1
2019-09-09 15:32:05.366015 +0530 IST m=+0.000403928
out 2
2019-09-09 15:32:05.366039 +0530 IST m=+0.000428903
>>>2
>>>2
>>>2
>>>2
>>>1
>>>1
>>>1
>>>1
out 1
2019-09-09 15:32:05.366121 +0530 IST m=+0.000510047
out 2
2019-09-09 15:32:05.366134 +0530 IST m=+0.000523278

*/
