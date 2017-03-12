package main

import (
	"fmt"
)

func main() {
    fmt.Printf("hello world%s\n\n", subFunc())
}

func subFunc() string {
	return ", con un qualcosa in più"
}

