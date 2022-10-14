package main

import (
	"fmt"
	"unsafe"
)

type student struct {
	Name string
	Age  int
}

func main() {
	students := new(student)
	name := (*string)(unsafe.Pointer(students))
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(students)) + uintptr(unsafe.Offsetof(students.Age))))
	*name = "zhangjie"
	*age = 24
	u := uintptr(unsafe.Pointer(name))
	pointer := (*string)(unsafe.Pointer(u))
	fmt.Println(students)
	fmt.Println(*pointer)
	fmt.Println(1 ^ 2 ^ 3 ^ 1 ^ 2 ^ 56 ^ 58 ^ 26 ^ 58 ^ 26 ^ 56 ^ 104 ^ 3)
}

// 有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
func findOnlyOnce(array []int) bool {
	return false
}
