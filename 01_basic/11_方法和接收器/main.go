package main

import (
	"fmt"
	"time"
)

//方法和接收器
/*
	方法（method）是一种作用于特定类型的函数，这种特定类型叫做接收器（receiver）（目标接收器）
	如果将特定类型理解为结构体或“类”时，接收器的概念就相当于是实例，也就是其它语言中的 this 或 self
	接收器的类型可以是任何类型，
	不仅仅是结构体，任何类型都可以拥有方法

	为结构体添加方法:
		需求说明：使用背包作为“对象”，将物品放入背包的过程作为“方法”，通过面向过程的方式和结构体的方式来解释“方法”的概念
		1.面向过程方式
		2.结构体方法
			为 *Bag 创建一个方法，(bag *Bag) 表示接收器，即 put 方法作用的对象实例

	结构体方法的继承(demo1中)
		模拟面向对象的设计思想（人和鸟的特性）

	为任意类型添加方法
		因为结构体也是一种类型，给其它类型添加方法和给结构体添加方法一样
		可以给基本类型添加方法：

	time 包中的基本类型方法：
		time.Second 的类型是 Duration，而 Duration 实际是一个 int64 的类型
		对于 Duration 类型有一个 String() 方法，可以将 Duration 的值转为字符串
 */


//接收器
/*
	每个方法只能有一个接收器，如下图：
		接收器的格式如下：
		func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
			函数体
		}

	各部分说明：
		接收器变量：命名，官方建议使用接收器类型的第一个小写字母，例如，Socket 类型的接收器变量应该为 s，Connector 类型应该命名为 c
		接收器类型：与参数类似，可以是指针类型和非指针类型，两种接收器会被用于不同性能和功能要求的代码中（需要做更新操作时，用指针类型）
		方法名、参数列表、返回参数：与函数定义相同
		1.理解指针类型的接收器
			更接近于面向对象中的 this 或 self；由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的
		2.理解非指针类型的接收器
			当方法作用于非指针接收器时，会在代码运行时将接收器的值复制一份；可以获取接收器的成员值，但修改后无效
		3.关于指针和非指针接收器的使用
			在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器。大对象因为复制性能较低，适合使用指针接收器（在接收器和参数间传递时不进行复制，只是传递指针）
 */


//背包结构体
type Bag struct {
	items []string
}


//面向过程方式的方法
func put(b *Bag, item string) {
	b.items = append(b.items, item)
}

//结构体方法
func (b *Bag) put(item string) {
	b.items = append(b.items, item)
}

//给基本类型添加方法
type MyInt int
func (a *MyInt) set(i int){
	*a = MyInt(i)
}

//接收器
type Person struct {
	Name string
}
//设置name
func (p *Person) setName(name string) {
	p.Name = name
}
//获取name
func (p *Person) getName() (name string) {
	return p.Name
}


func main() {
	//为结构体添加方法: 面向过程方式
	var bag *Bag = new(Bag)
	put(bag, "macBookPro")
	fmt.Println(bag)

	//结构体方法
	bag.put("food")
	fmt.Println(bag)

	//给基本类型添加方法：
	var a *MyInt = new(MyInt)
	a.set(66)
	fmt.Println(*a) //66

	//time包中的基本类型方法
	var b time.Duration = time.Second * 2
	fmt.Printf("b: %v b type: %T \n", b, b) //b: 2s b type: time.Duration
	var c string = b.String()
	fmt.Printf("c: %v c type: %T \n", c, c) //c: 2s c type: string

	//指针类型的接收器
	var person *Person = new(Person)
	person.setName("johny")
	var name string = person.getName()
	fmt.Println(name)
}