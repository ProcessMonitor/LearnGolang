package main

import "fmt"

//父类
type Human struct {
	name     string
	gender   string
	sexLevel int
	objType  string
}

//子类（通过父类对象）
type SuperMan struct {
	Human
	skill string
}

//父类方法
func (this *Human) Eat() string {
	return this.name + "Man.eat() ..."
}
func (this Human) Walk() {
	fmt.Println(this.name + "human walking ..")
}
func (this *Human) levelUp() string {
	this.sexLevel++
	return this.name + "sexlevel ++"
}

//子类方法
func (this *SuperMan) Eat() {
	fmt.Println(this.skill + " SuperMan.eat() ...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly ...")
}

func (this *SuperMan) PrintSuperMan() {
	fmt.Println("一个性别为", this.gender, "叫做", this.name, "的",
		this.objType, "通过", this.skill, "使得", this.levelUp(), "现在的 level", this.sexLevel)
}

func main() {

	superman := SuperMan{Human{"zwq ", "male", 100, "Master"}, "fk"}
	fmt.Println("new Super man:", superman)

	superman.Walk() //父类
	superman.Eat()  //子类
	superman.Fly()  //子类

	superman.PrintSuperMan()
}
