package main

import (
	"fmt"
	"reflect"
)

func reflectTypeOf(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}
func reflectValueOf(i interface{}) {
	fmt.Println(reflect.ValueOf(i))
}
func ReflectUser(ii interface{}) {
	//获取input的value
	inputValue := reflect.ValueOf(ii)
	fmt.Println("inputVal ", inputValue)
	inputType := reflect.TypeOf(ii)
	fmt.Println("inputType ", inputType)
	//1.获取interface的reflect.Type, 通过Type获得Num of Field,遍历
	for i := 0; i < inputType.NumField(); i++ {
		//2.找到每个field 数据类型
		fieldName := inputType.Field(i).Name
		//3.field的interface()方法得到对应的value
		value := inputValue.Field(i).Interface()
		fmt.Println("Field Name: ", fieldName, " value:", value)

	}
	//通过Type获取其中的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println("MethodName ,Type :", m.Name, m.Type, m.Func)
	}
}

type User struct {
	//每一行是一个field
	Id   int
	Name string
	Age  int
}

//必须大写不然ReflectUser中for循环检索不到此方法
func (this User) PrintUser() {
	fmt.Println("User :", this)
}
func (this User) Call() {
	fmt.Println("User :", this)
}
func main() {
	//var a, b, c = 1, "s t r i n g", 0.356565

	//reflectTypeOf(a)
	//reflectValueOf(a)
	//reflectTypeOf(b)
	//reflectValueOf(b)
	//reflectTypeOf(c)
	//reflectValueOf(c)

	user := User{123, "ban", 11}
	ReflectUser(user)

}
