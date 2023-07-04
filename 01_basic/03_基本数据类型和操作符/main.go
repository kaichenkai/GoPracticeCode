package main

import (
	"fmt"
	"math"
)

//数字类型：int8,  int16,  int32,  int64,　　uint8,  uint16, uint32,  uint64 //
//bool 类型：ture,  false　　（bool 型无法参与数值运算，也无法与其他类型进行转换。）
//浮点类型：float32,  float64
//字符类型：byte


//字符串类型：字符串实现基于 UTF-8 编码
/*
	""　　双引号，定义单行字符串
	``　　反引号，定义多行字符串（在这种方式下，反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出）
                多行字符串一般用于内嵌源码和内嵌数据等
 */


//操作符：
/*
数字操作符：+,  -,  *,  /,  %
比较运算符：>,  >=,  <,  <=,  ==,  !=
 */


//字符串拼接
/*
字符串不能与数字进行拼接
 */


func main() {
	//类型转换
	var a int = 8
	fmt.Printf("a的类型是：%T \n", a)
	var b int32 = int32(a)
	fmt.Println(a, b)
	//
	//浮点数转换成 int 类型，精度会丢失
	var c float32 = math.Pi
	var d = int(c)
	fmt.Println(d)
	//
	//int32转换成int16，数据会截断
	var e int32 = 1047483647
	fmt.Printf("int32: 0x%x %d \n", e, e)

	var f int16 = int16(e)
	fmt.Printf("int16: 0x%x %d \n", f, f)
	//简单写法，自动推导类型
	g := int16(e)
	fmt.Println(g)

	//字符串拼接: +
	var str1 string = "hello"
	var str2 string = "world"
	var str3 string = str1 + str2
	fmt.Println(str3)

	//获取字符串长度
	var length int = len(str3)
	fmt.Println(length)

	//字符串切片
	var substr string = str3[:5]
	fmt.Println(substr)
}