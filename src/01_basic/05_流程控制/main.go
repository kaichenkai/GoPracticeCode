package main

import (
	"fmt"
)

func main() {
	//if else 分支判断
	var num = 6
	if num >= 6 {
		println("num >= 6")
	} else {
		println("num < 6")
	}

	//switch case
	/*
	case 后边的值可以写多个，是 或 的关系
	case 语句块末尾如果加上 fallthrough，会接着执行下一个 case 的语句块
	*/
	var num2 = 7
	switch num2 {
	case 1, 2:
		fmt.Println("num2 等于1或者2")
	case 3:
		fmt.Println("num2 等于3")
	case 7:
		fmt.Println("num2 等于7")
		//fallthrough  // 会执行下一个case的语句块
	default:
		fmt.Println("default output")
	}

	//for 循环
	for i:=0; i<=100; i++ {
		println(i)
	}

	//for 死循环
	//for {
	//	print("hello world")
	//	time.Sleep(time.Second)
	//}

	//for 加判断 （while）
	var j int = 1
	for j <= 100 {
		println(j)
		j ++
	}

	//for range
	/*
	可以使用 for range 遍历数组、切片、字符串、map 及通道（channel）。通过 for range 遍历的返回值有一定的规律：
	数组、切片、字符串返回索引和值
	map 返回键和值
	通道（channel）只返回通道内的值
	 */
	var arrayA [3]string = [3]string{"hammer", "soldier", "mum"}
	for index, value := range arrayA {
		println(index, value)
	}
	//用匿名标识符来忽略 index
	for _, value := range arrayA {
		println(value)
	}

	//goto
	/*
	可以通过标签进行代码间的无条件跳转，goto语句可以快速跳出循环 或 实现同样的逻辑 有一定的帮助：
	快速跳出循环：
	 */
	for i:=1; i<=100; i++ {
		if i == 10 {
			//直接跳转到标签函数
			goto breakHere
		}
	}

	breakHere:
		println("将要退出for循环")
	//使用goto集中处理错误：
	/*
	err := firstCheckError()
	if err != nil {
	    goto onExit
	}

	err = secondCheckError()
	if err != nil {
	    goto onExit
	}

	fmt.Println("done")
	return

	onExit:
	    fmt.Println(err)
	    exitProcess()
	 */
}
