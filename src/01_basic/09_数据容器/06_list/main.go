package main

import (
	"container/list"
	"fmt"
)

//列表 list
/*
	和 python 的列表很类似，它可以存放任意类型，可以插入数据，可以移除数据，可以获取数据（不能通过索引取值）
	有两种方式初始化列表（得到的类型不同）
		通过 container/list 包的 New 方法：
 		直接声明一个 list

	列表增加元素
		PushFront
		PushBack

	列表删除元素
		遍历双链表需要配合 Front() 函数获取头元素句柄，Value 属性可以获取句柄的值，遍历时只要元素不为空就可以继续进行。每一次遍历调用元素的 Next
 */

func main() {
	//通过 container/list 包的 New 方法： new 的是指针
	var listA *list.List = list.New()
	fmt.Printf("listA type: %T \n", listA)  //*list.List

	//直接声明
	var listB list.List
	fmt.Printf("listB type: %T \n", listB)  //list.List

	//列表插入数据
	var listC list.List
	//var listC *list.List = list.New()
	element1 := listC.PushFront("first")// 返回一个元素句柄 (*list.Element 结构)
	element2 := listC.PushBack(666)
	fmt.Println(element1)
	fmt.Println(element2)

	//列表删除元素
	var listD list.List
	// 取句柄的值
	element3 := listD.PushFront("first")
	element4 := listC.PushBack(666)
	listD.PushBack(666)
	fmt.Println("element3: ", element3.Value)
	fmt.Println("element4: ", element4.Value)

	// 遍历列表
	for i:=listD.Front(); i!=nil; i=i.Next(){
		fmt.Println(i.Value)
	}
}