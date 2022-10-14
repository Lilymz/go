package main

import "fmt"

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	ints := s[:1]
	fmt.Printf("%+v\n", ints)
	fmt.Printf("%p\n,leng:%d\n", s, len(s))
	s = append(s[:1], s[2:]...)
	fmt.Printf("%p\n,leng:%d\n", s, len(s))
	fmt.Printf("%+v\n", s)
	value, _ := m["q1mi"]
	fmt.Printf("%p\n,leng:%d\n", value, len(value))
	fmt.Printf("%+v\n", m["q1mi"])
	// 引用的是不同的切片 但是指向的是同一个数组 切片组成是[]int 也就是长度+容量+数据，有一个不一样就是一个不同的切片，即使他们使用的是同一组数组
	fmt.Printf("===========================================\n")
	studentMap := make(map[string]string, 1)
	fmt.Printf("当前指针地址：%p\n", &studentMap)
	studentMap["zhangjie"] = "厦门集微科技有限公司"
	studentMap["mengmeixin"] = "a1食品科技有限公司"
	fmt.Printf("%#v,len:%d\n", studentMap, len(studentMap))
	//验证map是不是存在底层扩容
	studentMap["jiangyuhang"] = "厦门致上信息科技有限公司"
	fmt.Printf("当前指针地址：%p\n", &studentMap)
	studentMap["huangliangtao"] = "厦门星网锐捷科技有限公司"
	studentMap["sfsf"] = "sfsaf"
	studentMap["gsagas"] = "sagasg"
	fmt.Printf("%#v,len:%d\n", studentMap, len(studentMap))
	fmt.Printf("当前指针地址：%p\n", &studentMap)
}
