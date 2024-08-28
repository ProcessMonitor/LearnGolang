package main

import (
	"fmt"
)

// interface in go 是一个指针AnimalIF是指针名
type AnimalIF interface {
	Sleep()
	GetColor() string //获取抽象颜色
	GetType() string  //获取抽象种类
}

//子类（通过实现全部方法）
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println(this.color + " Cat sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println(this.color + " Dog sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	fmt.Println(animal.GetColor() + " " + animal.GetType())
}

func main() {
	var animal_test AnimalIF //接口数据类型 本质是指针
	animal_test = &Cat{"gray"}
	animal_test.Sleep() //cat 指针过去 就是cat的sleep
	animal_test = &Dog{"yellow"}
	animal_test.Sleep() //dog 指针过去 就是dog的sleep

	cat := Cat{"red"}
	dog := Dog{"green"}
	showAnimal(&cat)
	showAnimal(&dog)
}
