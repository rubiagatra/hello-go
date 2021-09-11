package main

import (
	"fmt"
	"os"
)

var name = os.Getenv("NAMA")

func main() {
	fmt.Println("Hello", name)
}
