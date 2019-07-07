package main
  
import "fmt"
var x uint64
var y int64
var z int8 //store values fron -128 to 127
func main() {
        x = 8193085864
        y = -8193085864
        z = 127
        fmt.Println(x)
        fmt.Printf("%T\n",x)
        fmt.Println(y)
        fmt.Printf("%T\n",y)
        fmt.Println(z)
        fmt.Printf("%T",z)
}

