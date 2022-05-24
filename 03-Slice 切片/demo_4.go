package main

import "fmt"

//删除切片
func main() {
	sli := [] int {1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("len = %d cap = %d slice = %v\n",len(sli), cap(sli), sli)

	//删除尾部两个元素
	fmt.Printf("len = %d cap = %d slice = %v\n", len(sli[:len(sli) - 2]), cap(sli[:len(sli) - 2]), sli[:len(sli) - 2])

	//删除开头2个元素
	fmt.Printf("len = %d cap = %d slice = %v\n", len(sli[2:]), cap(sli[2:]), sli[2:])

	//删除中间2个元素
	sli = append(sli[:3], sli[3+2:]...)
	fmt.Printf("len = %d cap = %d slice = %v\n", len(sli), cap(sli), sli)
}