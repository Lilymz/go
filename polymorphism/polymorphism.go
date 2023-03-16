package main

/**
  golang多态的实现
*/
import "fmt"

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

// if else continue break switch case default fallthrough goto select for range return
func (t *Teacher) studying() {
	fmt.Println("name:" + t.Person.name + "老师,school:" + t.Person.school + " 正在进行培训")
}
func (s *Student) studying() {
	fmt.Println("name:" + s.Person.name + "学生,school:" + s.Person.school + " 正在进行学习")
}

type Study interface {
	studying()
}

func running(runner Study) {
	runner.studying()
}
func main() {
	student := &Student{
		Person{
			"zhangjie",
			"清华大学",
		},
	}
	teacher := &Teacher{
		Person{
			"meixin",
			"清华大学教授",
		},
	}
	running(student)
	running(teacher)
	var (
		study Study
	)
	study = student
	study.studying()
}
