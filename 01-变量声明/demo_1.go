package main

import (
	"fmt"
)

func main() {
	const name string = "Tom"
	fmt.Println(name)

	const age = 20
	fmt.Println(age)

	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)
}