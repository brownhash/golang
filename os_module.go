package main
  
import (
        "fmt"
        "os"
        "os/exec"
)

func main() {
        //var a int
        fmt.Println("Welcome . . .")
        //fmt.Println("\nEnter the age : ")
        //fmt.Scanf("%d",&a)
        cmd := exec.Command("ls")
        cmd.Stdout = os.Stdout
        //cmd.Stderr = os.Stderr
        err := cmd.Run()
        fmt.Println(err)
}
