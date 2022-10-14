package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	number       int8
	students     = make([]*Student, 0)
	stuNo        int
	stuName      string
	formatModify = "请输入"
)

type Student struct {
	number int
	name   string
}

func (s Student) newStudent(number int, name string) *Student {
	return &Student{
		number: number,
		name:   name,
	}
}
func (s *Student) addStudent() {
	if students == nil || len(students) < s.number {
		stu := append(students, s)
		*(&students) = stu
		fmt.Printf("添加成功！学生个数：%d\n", len(students))
		return
	} else {
		if len(students) > s.number {
			preStudents := students[:s.number]
			aftStudents := students[s.number-1:]
			students = append(preStudents, aftStudents...)
			fmt.Printf("添加成功！学生个数：%d\n", len(students))
		}
	}
}
func (s *Student) modifyStudent() {
	students[s.number] = s
}
func (s *Student) deleteStudent() {
	preStudents := students[:s.number]
	aftStudents := students[s.number:]
	stu := append(preStudents, aftStudents...)
	*(&students) = stu
	fmt.Printf("删除成功！学生个数：%d\n", len(students))
}
func queryStudent() {
	for _, student := range students {
		fmt.Printf("student：%#v\n", student)
	}
	marshal, err := json.Marshal(students)
	if err != nil {
		fmt.Println("json serial error ~")
	}
	fmt.Println(string(marshal))
}
func getInput(number int) *Student {
	var stu Student
	if number == 1 {
		fmt.Print("请输入学生学号： ")
		fmt.Scanf("%d\n", &stuNo)
		formatModify = "请输入"
		fmt.Print(formatModify + "学生姓名： ")
		fmt.Scanf("%s\n", &stuName)
		newStudent := stu.newStudent(stuNo, stuName)
		return newStudent
	} else if number == 2 {
		fmt.Print("请输入学生学号： ")
		fmt.Scanf("%d\n", &stuNo)
		formatModify = "编辑"
		fmt.Print(formatModify + "学生姓名： ")
		fmt.Scanf("%s\n", &stuName)
		modifyStudent := stu.newStudent(stuNo, stuName)
		return modifyStudent
	} else {
		fmt.Print("请输入需删除的学生学号： ")
		fmt.Scanf("%d\n", &stuNo)
		deleteStudent := stu.newStudent(stuNo, "")
		return deleteStudent
	}
}

func main() {
	for {
		fmt.Println("==========================欢迎进入go学生信息管理系统 begin==============================")
		fmt.Println("\t\t\t\t\t\t\t1.添加学生")
		fmt.Println("\t\t\t\t\t\t\t2.编辑学生")
		fmt.Println("\t\t\t\t\t\t\t3.查询所有学生")
		fmt.Println("\t\t\t\t\t\t\t4.删除学生")
		fmt.Println("\t\t\t\t\t\t\t5.退出系统")
		fmt.Println("==========================欢迎进入go学生信息管理系统 end==============================")
		fmt.Print("请选择对应操作序号: ")
		fmt.Scanf("%d\n", &number)
		switch number {
		case 1:
			student := getInput(int(number))
			student.addStudent()
			break
		case 2:
			student := getInput(int(number))
			student.modifyStudent()
			break
		case 3:
			queryStudent()
			break
		case 4:
			student := getInput(int(number))
			student.deleteStudent()
			break
		case 5:
			os.Exit(0)
			break
		default:
			fmt.Println("当前序号选择有误！请重新选择！")
		}
	}
}
