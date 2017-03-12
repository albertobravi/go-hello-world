package main

import (
	"fmt"
	"github.com/albertobravi/stringutil"
)

func main() {
    fmt.Printf("hello world%s\n\n", subFunc())

	fmt.Print("!oG ,olleH => " + stringutil.Reverse("!oG ,olleH") + "\n")
}

func subFunc() string {
	return ", con un qualcosa in più"
}

