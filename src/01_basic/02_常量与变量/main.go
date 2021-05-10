package main

import (
	"fmt"
	"go/types"
)

//常量使用const修饰，表示是只读的，不能修改
//const只能修饰 boolean, number(int相关，浮点数，complex) 和 string类型
//语法：const identifier [type] = value（type 可省略）
const (
	name string = "skier"
	age int = 10
	salary int = 5000 / 2
	isLive = true
	//gender bool = getGender()  //const不能从函数中获取
)
//常量在编译期确定，所以可以用于数组声明
const size = 4
var arrayA [size]int


//变量
//声明一个变量：var identifier [type]
var a int = 100
//简写：(自动推导类型)
var b = 100
var c, d  = 100, 200
//优雅写法
var (
	name2 string = "johny"
	age2 int = 12
	sliceA types.Slice
)

//变量的默认值
//int 0
//float 0.0		默认推导为 float64位
//bool false
//string ""
//slice			函数，指针类型默认为nil

//遍历数组
var arrayB = []string{"hammer", "soldier", "mum"}


func main() {
	fmt.Println(arrayA)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(name2)
	fmt.Println(age2)
	fmt.Println(sliceA)
	fmt.Printf("format test %s %d \n", name, age)
	//字符串拼接
	var name3 string = name + name2
	fmt.Println(name3)
	//获取字符串长度
	fmt.Println(len(name3))
	//遍历数组
	for index, value := range arrayB {
		fmt.Println(index, value)
	}
	//
	fmt.Println(`test 
test`)

}
