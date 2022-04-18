// Exercise 1.1
// Echo4 prints it's command-line arguments including the executables name.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
