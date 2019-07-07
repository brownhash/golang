package main

import (
	"fmt"
	"log"
	"os/exec"
	. "runtime"
)

func main() {
	fmt.Println("Enter a URL")
	var a string
	fmt.Scanf("%s",&a)
	openbrowser(a)
}
func openbrowser(url string) {
	var err error

	switch GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
