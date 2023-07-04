package main

import "fmt"

//结构体
/*
	结构体定义
		结构体的定义只是一种内存布局的描述（相当于是一个模板），只有当结构体实例化时，才会真正分配内存空间
		结构体是一种复合的基本类型，通过关键字 type 定义为 自定义 类型后，使结构体更便于使用
		定义一个简单的结构体：
			1.类型名（Point）在同一个包内不能重复
			2.字段名（X,  Y）必须唯一

	实例
		结构体实例与实例间的内存是完全独立的，
		可以通过多种方式实例化结构体，根据实际需求选用不同的写法
		1.基本的实例化形式（不推荐）
			结构体是一种[值类型]，可以像整型，字符串一样，以 var 开头的方式声明结构体即可完成实例化
			基本实例化格式：
				var ins T
				其中，T 为结构体类型，ins 为结构体的实例 var ins *T 创建指针类型的结构体，T 为结构体指针类型
				用声明的方式创建指针类型的结构体，然后进行赋值会触发 panic（空指针引用）
		2.创建指针类型的结构体（推荐使用）
			Go语言中，可以使用 new 关键字对值类型（包括结构体，整型，字符串等）进行实例化，得到指针类型
				ins := new(T)
				其中，T 为结构体类型，ins 为指针类型 *T
		3.使用 键值对 初始化结构体（实例化时直接填充值）
			每个键对应结构体中的一个字段，值对应字段中需要初始化的值
			键值对的填充是可选的，不需要初始化的字段可以不在初始化语句块中体现（字段的默认值，是字段类型的默认值，例如：整数是0，字符串是 ""，布尔是 false，指针是 nil 等）

	结构体的嵌套（递归）
		结构体成员中只能包含[自身]结构体的指针类型，包含非指针类型会引起编译错误（invalid recursive type 'People'）

	匿名结构体

	模拟构造函数 初始化结构体
		如果使用结构体来描述猫的特性，那么根据猫的名称与颜色可以有不同的类型，那么可以使用不同名称与颜色可以构造不同猫的实例：

	带有 继承关系 的结构体的构造与初始化（例子）
		猫是基本结构体（只有姓名和颜色）
		黑猫继承自猫，是子结构体（不仅有姓名和颜色，还有技能）
		使用不同的两个构造函数分别构造猫与黑猫两个结构体实例：
		Cat 结构体类似于面向对象中的“基类”。BlackCat 嵌入 Cat 结构体，类似于面向对象中的“派生”。实例化时，BlackCat 中的 Cat 也会一并被实例化

	结构体匿名字段
		上面的 黑猫 与 猫 的继承关系中，定义黑猫字段的时候，就用到了匿名字段：
			黑猫的结构体（继承了猫，增加了技能字段）
			type BlackCat struct{
    			Cat
    			Skill string
			}
		一个结构体中只能有一个同类型的匿名字段，不需要担心结构体字段重复问题

	结构体字段标签
		结构体标签是指对结构体字段的额外信息，进行 json 序列化及对象关系映射（Object Relational Mapping）时都会用到结构体标签，标签信息都是静态的，无须实例化结构体，可以通过反射拿到（反射后面会有记录）
		Tag 在结构体字段后面书写，格式如下：
		由一个或多个键值组成，键值对之间使用空格分隔：
			`key1:"value1" key2:"value2"`

	结构体标签格式错误导致的问题
		package main
		import (
			"fmt"
			"reflect"
		)
		func main() {
			type cat struct {
				Name string
				Type int `json: "type" id:"100"`
			}
			typeOfCat := reflect.TypeOf(cat{})
			if catType, ok := typeOfCat.FieldByName("Type"); ok {
				fmt.Printf("'%v'", catType.Tag.Get("json"))
			}
		}

		运行结果：
		''  //空字符串

		在json:和"type"之间增加了一个空格。这种写法没有遵守结构体标签的规则，因此无法通过 Tag.Get 获取到正确的 json 对应的值
*/


//定义一个坐标的结构体
type point struct {
	//x int
	//y int
	//同类型变量可以写成一行
	x, y int
}

//定义一个人的结构体
type People struct{
	Name string
	Child *People
}

//定义一个猫的结构体
type Cat struct {
	Name string
	Color string
}

//模拟构造函数，初始化结构体
func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

//带有 继承关系 的结构体的构造与初始化（例子）
/*
	猫是基本结构体（只有姓名和颜色）
	黑猫继承自猫，是子结构体（不仅有姓名和颜色，还有技能）
	使用不同的两个构造函数分别构造猫与黑猫两个结构体实例：
 */
//黑猫的结构体
type BlackCat struct {
	Cat
	Skill string
}
// 构造黑猫的函数
func NewBlackCat(name, color, skill string) *BlackCat {
	var cat Cat = Cat{
		Name: name,
		Color: color,
	}
	return &BlackCat{
		Cat: cat,
		Skill: skill,
	}
}

// 构造黑猫的函数
//func NewBlackCat(name, color, skill string) *BlackCat {
//	blackCat := &BlackCat{}
//	blackCat.Name = name
//	blackCat.Color = color
//	blackCat.Skill = skill
//	return blackCat
//}



func main() {
	//基本实例化格式：（不推荐）
	var p1 point
	p1.x = 1
	p1.y = 1
	fmt.Println(p1)

	//用声明的方式创建指针类型的结构体，然后进行赋值会触发 panic（空指针引用）
	//var p2 *point
	//p2.x = 3
	//p2.y = 3
	//fmt.Println(p2)  //invalid memory address or nil pointer dereference

	//创建指针类型的结构体（推荐使用）
	var p3 *point = new(point)
	p3.x = 3
	p3.y =3
	fmt.Printf("point3: %v,type: %T \n", p3, p3) //point3: &{3 3},type: *main.point

	//使用 键值对 初始化结构体（实例化时直接填充值）
	var p4 point = point{
		x: 4,
		y: 4,
	}
	fmt.Printf("point4: %v,type: %T \n", p4, p4) //point4: {4 4},type: main.point
	//实例化[指针类型]的结构体
	var p5 *point = &point{
		x: 5,
		y: 5,
	}
	fmt.Printf("point5: %v,type: %T \n", p5, p5) //point5: &{5 5},type: *main.point

	//结构体的嵌套
	var relation People = People{
		Name: "grandPa",
		Child: &People{
			Name: "father",
			Child: &People{
				Name: "me",
			},
		},
	}
	// 与上面等效
	var me People = People{
		Name: "me",
	}
	var father People = People{
		Name: "father",
		Child: &me,
	}
	var grandPa People = People{
		Name: "grandPa",
		Child: &father,
	}
	fmt.Printf("%T \n %v \n", relation, relation.Child.Child.Name)
	fmt.Printf("%T \n %v \n", grandPa, grandPa.Child.Child.Name)

	//匿名结构体
	var perpleA = &struct{
	//var perpleA = struct{
		name string
		age int
	}{
		name: "johny",
		age: 12,
	}
	fmt.Printf("perpleA: %v, perpleA type: %T \n", perpleA, perpleA)

	//模拟构造函数，初始化结构体
	var tom *Cat = NewCatByName("Tom")
	fmt.Println(tom)

	//带有继承关系的结构体构造
	var blackCat *BlackCat = NewBlackCat("blackCat", "black", "climb tree")
	fmt.Println(blackCat)
}