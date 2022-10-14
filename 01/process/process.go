package main

import "fmt"

//func main() {
//	addData, subData := calc()
//	sys.Println(addData(10))
//	sys.Println(subData(10))
//	println(greeting.Version())
//}
////if else
//func studentGrade(score float32)  {
//	if score :=score+1;score < 60 {
//		fmt.Println("E")
//	}else if score <= 70{
//		fmt.Println("D")
//	}else if score <=80 {
//		fmt.Println("C")
//	}else if score <=90 {
//		fmt.Println("B")
//	}else {
//		fmt.Println("A")
//	}
//	if false {
//		fmt.Println("进入标签")
//	}
//
//}

// SwitchDemo switch 可以进行多值判断，并且可以结合fallthrough
func SwitchDemo() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
func calc() (addData func(data int) int, subData func(data int) int) {
	var base = 20
	addData = func(add int) int {
		return base + add
	}
	subData = func(sub int) int {
		return base - sub
	}
	return
}
