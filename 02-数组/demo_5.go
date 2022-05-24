package main

import "fmt"

//数组是值类型问题，在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销。

func main() {
	arr := [5] int {1, 2, 3, 4, 5}
	modifyArr_1(arr)
	fmt.Println(arr)

	modifyArr_2(&arr)
	fmt.Println(arr)
}

func modifyArr_1(a [5] int) {
	a[1] = 20
}

func modifyArr_2(a *[5] int) {
	a[1] = 20
}

