// Exercise 1.2
// Echo5 prints each of its arguments and it's corresponding index
// on a new line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
