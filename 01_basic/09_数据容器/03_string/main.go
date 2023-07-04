package main

import "fmt"

//string 是值类型，底层是一个 byte 数组，因此，可以对 string 进行切片操作
/*
Go 语言中，string本身是不可变的，如何改变 string 中的字符？
 */


func main() {
	var a string = "hello world"
	//将字符串转换为字节切片 (数据转换)
	var sliceA []byte = []byte(a)
	//改变切片的值
	sliceA[0] = 'A'
	//再转换成字符串(数据转换)
	a = string(sliceA)
	//
	fmt.Println(a)

	//如果是中文呢 ？
	//var b string = "你好 世界"
	////将字符串转换为字节切片 (数据转换)
	//var sliceB []byte = []byte(b)
	////改变切片的值
	//sliceB[0] = '哈'
	////再转换成字符串(数据转换)
	//b = string(sliceB)
	////
	//fmt.Println(b)
}
