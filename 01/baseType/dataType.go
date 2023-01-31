package main

import (
	"flag"
	"fmt"
)

// 整型
var smallInt int8 = 1 << 6
var unsignedSmallInt uint8 = 1 << 7
var shortInt int32 = 1 << 30
var unsignedShortInt uint32 = 1 << 31
var longInt int64 = 1 << 62
var unsignedLongInt uint64 = 1 << 63

func main() {
	//整形
	fmt.Println(smallInt)
	fmt.Println(unsignedSmallInt)
	fmt.Println(shortInt)
	fmt.Println(unsignedShortInt)
	fmt.Println(longInt)
	fmt.Println(unsignedLongInt)
	//指针
	var name = "zhangjie"
	ptr := &name
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
	swap()
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)
	var pointerInt *int
	var uptr uintptr
	var number = 10
	pointerInt = &number
	*pointerInt = 30
	uptr = 30
	fmt.Println(*pointerInt)
	//当前指针地址
	fmt.Println(&(*pointerInt))
	//值地址
	fmt.Println(&uptr)
	//复数
	var c1 complex64
	var c2 complex128
	c1 = 1 + 2i
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
	byteAndRune()
}

func swap() {
	var a string = "a"
	var b string = "b"
	fmt.Printf("当前变量值a:%v,b:%v\n", a, b)
	// 使用指针进行值交换
	aptr := &a
	bptr := &b
	*aptr = "b"
	*bptr = "a"
	fmt.Printf("当前变量值a:%v,b:%v\n", a, b)
}

var mode = flag.String("mode", "", "process mode")

func byteAndRune() {
	var str = "hello,你好"
	bytes := []byte(str)
	i := len(bytes)
	fmt.Printf("当前字节长度：%d\n", i)
	for a := range str {
		fmt.Printf("%c,%d\n", rune(a), a)
	}
	//
	//// 字节类型
	//for i := 0; i < 255; i++ {
	//	fmt.Printf("%c,%d\n", i,i)
	//}
	//for i := 27801; i < 1<<16; i++ {
	//	fmt.Printf("%c,%d\n", i,i)
	//}
	// utf-8 rune类型

}
