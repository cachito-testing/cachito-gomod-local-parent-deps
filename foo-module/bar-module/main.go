package main

import (
	"fmt"
	"github.com/cachito-testing/cachito-gomod-local-parent-deps/baz-package"
	"github.com/cachito-testing/cachito-gomod-local-parent-deps/foo-module/bar-module/bar-package"
)


func main() {
	fmt.Println("Hello, local dependencies.")
	bar.Hi()
	baz.Hi()
}
