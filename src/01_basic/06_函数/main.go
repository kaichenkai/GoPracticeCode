package main

//Go 语言支持普通函数、匿名函数和闭包
/*
	普通函数声明：func 函数名(参数列表)  (返回值列表)  {函数体}

	不支持重载，一个源文件内不能有两个相同名称的函数

	函数也是一种类型，可以赋值给变量

	函数的传参方式：
	值传递 　　     （基本数据类型都是值传递）
	引用传递          （指针，slice，map，chan，interface）
	注意：无论是值传递还是引用传递，传递给函数的都是变量的副本，不过，值传递是对值的拷贝，引用传递是地址的拷贝，一般来说，地址拷贝更为高效，而值拷贝取决于拷贝对象的大小，对象越大，则性能越低

	返回值命名：
	返回值不需要定义，直接使用（命名的返回值变量的默认值为类型的默认值，即数值为 0，字符串为 ""，布尔为 false、指针为 nil）
 */


//匿名函数：
/*
	匿名函数没有函数名，只有函数体，可以直接被当做一种类型赋值给函数类型的变量，匿名函数也往往以变量的方式被传递
	匿名函数经常被用于实现回调函数、闭包
 */



//defer 延迟调用
/*
	1.延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时
	2.多个 defer 语句，按先进后出（栈）的顺序执行
	3.defer 语句中的变量，在 defer 声明时就决定了
	4.defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题
 */


func main() {
	var c int = calc(1, 2)
	println(c)

	//可变长参数
	calcArgs(1, 2, 3, 4)

	//递归函数
	var recursionRet int = recursion(5)
	println("递归函数运行结果：", recursionRet)

	//斐波拉契实现
	var fabRet int = fab(5)
	println("斐波拉契数运行结果：", fabRet)

	//匿名函数
	func(args string) {
		println(args)
	}("hello")

	//将匿名函数赋值给变量
	f := func(args string) {
		println(args)
	}
	f("world")//调用
}

func calc(a int, b int) int {
	var c int = a + b
	return c
}

//可变长参数
func calcArgs(args... int) {
	for index, value := range args {
		println(index, value)
	}
}

//递归函数
func recursion(n int) int {
	if n <= 1 {
		return 1
	}
	return recursion(n-1) * n
}

//斐波拉契数
func fab(n int) int {
	if n<=1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}


//defer 延迟调用
//关闭文件句柄
//注意：不能将这一句代码放在第3行空行处，一旦文件打开错误，f将为空，在延迟语句触发时，将触发宕机错误
//func read(){
//	file err := open(filename)<br>
//	if err != nil{
//		return
//	}
//	defer file.Close()
//}

//锁资源的释放
//func lock(){
//	mc.Lock()
//	defer mc.Unlock()
//}


//数据库连接的释放
//func connect(){
//	conn := openDatabase()
//	defer conn.Close()
//}