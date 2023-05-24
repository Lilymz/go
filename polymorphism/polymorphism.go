package main

/**
  golang多态的实现
*/
import (
	"fmt"
	"golang.org/x/net/http2/hpack"
)

type Person struct {
	name   string
	school string
}

type Student struct {
	Person
}
type Teacher struct {
	Person
}
type MyError interface {
	Error() string
}

// if else continue break switch case default fallthrough goto select for range return
func (t *Teacher) studying() {
	fmt.Println("name:" + t.Person.name + "老师,school:" + t.Person.school + " 正在进行培训")
}
func (s *Student) studying() {
	fmt.Println("name:" + s.Person.name + "学生,school:" + s.Person.school + " 正在进行学习")
}
func (s *Student) Error() string {
	return "panic"
}

type Study interface {
	studying()
}

func running(runner Study) {
	runner.studying()
}
func printError(printError error) {
	fmt.Println(printError.Error())
}
func main() {
	//student := &Student{
	//	Person{
	//		"zhangjie",
	//		"清华大学",
	//	},
	//}
	//teacher := &Teacher{
	//	Person{
	//		"meixin",
	//		"清华大学教授",
	//	},
	//}
	//running(student)
	//running(teacher)
	//var (
	//	study Study
	//)
	var s1 *Student
	fmt.Println(s1)
	s2 := new(Student)
	s2.name = "zhangjie"
	fmt.Println(s2)
	//study = student
	//study.studying()
	//printError(student)
	var error hpack.DecodingError
	error.Err
	fmt.Println(error.Error())
}
