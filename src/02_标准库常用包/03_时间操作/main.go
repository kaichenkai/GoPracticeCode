package main

import (
	"fmt"
	"time"
)

//使用之前需要导入 time 包
/*
	time 包中的常量
	const (
		Nanosecond  Duration = 1
		Microsecond  = 1000 * Nanosecond
		Millisecond  = 1000 * Microsecond
		Second       = 1000 * Millisecond
		Minute       = 60 * Second
		Hour         = 60 * Minute
	)
	时间的类型为 Duration，而 Duration 实际是一个 int64 的类型，可以转换；
	Duration.String() 可以将 Duration 的值转为字符串
*/


//自定义时间
/*
	格式：func Parse(layout, value string) (Time, error)
	Parse 解析一个格式化的时间字符串并返回它代表的时间。
	layout 定义输入的时间格式，value 的时间格式需与 layout 保持一致
*/


//获取time.Time 类型的 年，月，日，时，分，秒
/*
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
*/


//获取当前时间戳
/*
	// 秒级
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	// 毫秒
	milliTimestamp := time.Now().UnixNano() / 1e6
	fmt.Println(milliTimestamp)

	// 纳秒
	NanoTimestamp := time.Now().UnixNano()
	fmt.Println(NanoTimestamp)
 */


func main() {
	//获取当前时间
	var now time.Time = time.Now()
	fmt.Printf("type:%T \n", now)
	fmt.Printf("value:%v \n", now)

	//自定义时间
	t, err := time.Parse("2006-01-02 15:04:05", "2019-05-25 19:00:00")
	if err == nil{
		fmt.Printf("%T \n", t)
		fmt.Printf("%v \n", t)
	}

	//获取当前时间戳
	// 秒级
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	// 毫秒
	milliTimestamp := time.Now().UnixNano() / 1e6
	fmt.Println(milliTimestamp)

	// 纳秒
	NanoTimestamp := time.Now().UnixNano()
	fmt.Println(NanoTimestamp)

	//统计程序运行耗时
	var startTime int64 = time.Now().UnixNano() / 1e6
	testPro()
	var endTime int64 = time.Now().UnixNano() / 1e6
	fmt.Printf("use time: %d ms \n", endTime - startTime)
}

func testPro() {
	time.Sleep(time.Second * 2)
}