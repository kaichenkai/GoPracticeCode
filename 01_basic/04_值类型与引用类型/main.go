package main

import "fmt"

//值类型：
/*
本质上是原始类型，变量直接储存值，内存通常在栈中分配，包括 int,  float,  bool,  string 以及数组和 struct（结构体）
对值类型进行操作，一般都会返回一个新创建的值，所以把这些值传递给函数时，其实传递的是一个值的副本
 */


//引用类型
/*
引用类型与值类型恰恰相反，它的修改可以影响到任何引用到它的变量；变量存储的是地址，这个地址存储最终的值，通常在堆内存上分配，通过 GC 回收，包括 指针，select，map，chan 等
引用类型之所以可以引用，是因为我们创建引用类型的变量，其实是一个标头值，标头值里包含一个指针，指向底层的数据结构，当我们在函数中传递引用类型时，其实传递的是这个标头值的副本，它所指向的底层结构并没有被复制传递，这也是引用类型传递高效的原因。
*/

func main(){
	//传递值类型
	var args string = "johny"
	modifyValue(args)
	fmt.Println(args)

	//传递引用类型 指针
	var quote int = 6
	modifyQuote(&quote)
	fmt.Println(quote)  // 7
}

func modifyValue(args string) {
	args += args
}

func modifyQuote(quote *int) {
	*quote += 1
}
