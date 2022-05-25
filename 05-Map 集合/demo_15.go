package main

import "fmt"

//编辑、删除、新增
func main() {
	person := map[int]string {
		1 : "Tom",
		2 : "Araron",
		3 : "John",
	}
	fmt.Println("data :", person)

	delete(person, 2)
	fmt.Println("data :", person)

	person[2] = "Jack"
	person[4] = "kevin"
	fmt.Println("data :", person)
}
