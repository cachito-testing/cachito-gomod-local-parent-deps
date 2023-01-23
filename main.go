package main

import (
	"fmt"
	"github.com/cachito-testing/cachito-gomod-local-parent-deps/baz-package"
)


func main() {
	fmt.Println("Hello, local dependencies.")
	baz.Hi()
}
