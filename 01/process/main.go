package main

import "fmt"

const (
	land = iota
	water
)

type animal struct {
	name  string
	class int
}

func (a animal) run() {
	fmt.Println(a.name + "正在奔跑....")
}
func (a animal) getClass() {
	fmt.Println("动物名称：" + a.name + "，所属类别：" + classStr(a.class))
}
func classStr(class int) string {
	if class == 0 {
		return "陆地"
	}
	if class == 1 {
		return "水下"
	}
	return "未知"
}

type Behavior interface {
	run()
	getClass()
}

func call(behavior Behavior) {
	behavior.run()
	behavior.getClass()
}
func get(hello *struct {
	name string
	age  string
}) {
	hello.name = "sss"
}
func main() {
	var unMemAnimal [5]animal
	fmt.Println(unMemAnimal)
	animalPoint := new(animal)
	fmt.Println(&animalPoint)
	var hello struct {
		name string
		age  string
	}
	hello.name = "s"
	get(&hello)
	fmt.Println(hello.name)
	var fish = animal{
		"鱼",
		water,
	}
	var cat = animal{
		"猫",
		land,
	}
	call(fish)
	call(cat)
}
