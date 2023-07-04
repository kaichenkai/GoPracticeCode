package main

import (
	"fmt"
	"sort"
)

//map 使用散列表（hash）实现，和 python 的 dictionary 一样，是无序的
/*
	和 slice 一样，创建 map 的两种方式（make 和 声明）
    与 slice 不一样的是：map 声明后是 nil 值，必须要初始化，直接使用触发 panic: assignment to entry in nil map

	遍历map（顺序问题）
		需要注意的是，遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果，如果你得到了顺序的结果，也不要庆幸，也许是当前的数据量不够大
		如果需要特定顺序的遍历结果，正确的做法是排序：把字典的 key 或 value 放入到切片中，然后进行排序

	map的删除和清空
		使用内建函数 delete() 从 map 中删除一组键值对（delete 没有返回值）
		清空map
		but it didn't
		唯一的方法就是重新 make 一个新 map，不用担心垃圾回收的效率，Go 语言中的并行垃圾回收效率比写一个清空函数高效多了
 */


func main() {
	//make
	var scene map[string]int = make(map[string]int)
	scene["route"] = 1
	fmt.Println(scene["route"])
	fmt.Println(scene["route2"]) //不存在的key不会报错，输出基本类型的默认值，和python不同

	//声明map并填充内容（字面量方式）
	var mapA map[string]int = map[string]int{"macbookpro": 13000, "iphone12": 6999}
	fmt.Println(mapA)
	fmt.Println(mapA["iphone12"])//6999
	fmt.Println(mapA["iphone11"])//0
	//声明空map
	var mapB map[string]int
	fmt.Println(mapB == nil)//true
	//mapB["key"] = 1 //panic: assignment to entry in nil map

	//判断某个key是否存在map中
	var mapC map[string]int = make(map[string]int)
	mapC["macBookPro"] = 13000
	value1, isExist1 := mapC["macBookPro"]
	fmt.Println(value1, isExist1) //13000 true
	value2, isExist2 := mapC["iphone10"]
	fmt.Println(value2, isExist2) //0 false

	//遍历map items
	for key, value := range mapC {
		fmt.Println(key, value)
	}
	//遍历，只需要key
	for key := range mapC {
		fmt.Println(key)
	}
	//遍历，使用匿名变量忽略key
	for _, value := range mapC {
		fmt.Println(value)
	}

	//map 按照顺序输出item
	var mapD map[string]string = make(map[string]string)
	mapD["w"] = "forward"
	mapD["a"] = "left"
	mapD["d"] = "right"
	mapD["s"] = "backend"
	fmt.Println(mapD)
	var sliceD []string
	for key := range mapD {
		sliceD = append(sliceD, key)
	}
	sort.Strings(sliceD)
	fmt.Println(sliceD)
	for _, key := range sliceD {
		fmt.Println(mapD[key])
	}

	//map 删除键值对
	var mapE map[string]string = make(map[string]string)
	mapE["macBookPro"] = "mbp"
	mapE["iphone12"] = "i12"
	delete(mapE, "iphone12")
	fmt.Println(mapE)

	//清空map 这个操作很少，基本不存在
	//but it didn't
	//唯一的方法就是重新 make 一个新 map，不用担心垃圾回收的效率，Go 语言中的并行垃圾回收效率比写一个清空函数高效多了
}