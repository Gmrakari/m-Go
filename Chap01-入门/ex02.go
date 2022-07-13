package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Print("%d:\t%s\n", i + 1, arg)
	}
}
