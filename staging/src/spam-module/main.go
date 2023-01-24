package main

import (
	"fmt"
	"github.com/cachito-testing/spam-module/spam"
)

func main() {
	fmt.Println("Hello from a module which isn't really a part of this repo.")
	spam.Hi()
}
