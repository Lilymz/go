package main

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	const five = 2
	//数组,指定索引位置初始化，未指定则采用默认值
	var scoreArray = [5]int{1: 20, 4: 82, 3: 65, five: 25}
	//遍历
	for i := 0; i < len(scoreArray); i++ {
		fmt.Println(scoreArray[i])
	}
	for _, value := range scoreArray {
		fmt.Printf("值：%v\n", value)
	}
	//基于数组建立切片
	sliceArray := scoreArray[0:3]
	sliceArray[0] = 100
	*(&scoreArray[0]) = 256
	fmt.Printf("当前数组集合%#v,len:%d,cap:%d\n", sliceArray, len(sliceArray), cap(sliceArray))

	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e.Error())
	fmt.Println(w)
	// 切片函数操作（追加，拷贝，删除，遍历，覆盖）包含
	sliceArray = append(sliceArray, 156)
	cpSliceArray := make([]int, 5, 5)
	i := copy(cpSliceArray, sliceArray)
	fmt.Printf("结果：%d,desination:%#v\n", i, cpSliceArray)
	contains := slices.Contains(cpSliceArray, 156)
	if contains {
		fmt.Println("该切片包含对应156数据！")
	}
	b := slices.Contains(cpSliceArray, 36655)
	if b {
		fmt.Println("存在36655数据")
	}
	search, b2 := slices.BinarySearch(cpSliceArray, 20)
	fmt.Printf("search 156 is exist:%t,value:%v\n", b2, search)

}
