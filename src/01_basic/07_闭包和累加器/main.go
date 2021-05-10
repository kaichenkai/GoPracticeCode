package main

import "fmt"

//闭包（Closure）
/*
	闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境 也不会被释放或者删除，在闭包中可以继续使用这个自由变量（闭包（Closure）在某些编程语言中也被称为 Lambda 表达式）
	简单的说：
	闭包 = 函数 + 引用环境
 */



func main(){
	f := closureFunc("world~")
	// 调用闭包函数
	f()

	//累加器的实现
	accumulator := accumulate(10)
	ret1 := accumulator()
	fmt.Println(ret1)
	ret2 := accumulator()
	fmt.Println(ret2)
}


//实现一个简单的闭包：
func closureFunc(str string) func(){
	wapper := func(){
		fmt.Println("hello", str)
	}
	return wapper
}

//累加器的实现（闭包的记忆效应）
func accumulate(num int) func() int {
	wapper := func() int {
		num += 1
		return num
	}
	return wapper
}