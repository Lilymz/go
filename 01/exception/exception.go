package main

import "fmt"

func main() {
	fmt.Println("执行开始！")
	executeWork()
	fmt.Println("执行结束！")
}
func executeWork() {
	defer canException()
	throwException()
}
func canException() {
	// a2就是触发panic能够获取到的数据
	a2 := recover()
	if a2 != nil {
		fmt.Println("已经进行错误处理,", a2)
	}
}
func throwException() {
	panic("丢出一个异常")
}
