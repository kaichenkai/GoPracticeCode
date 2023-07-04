package main

import (
	"fmt"
	"sort"
)

//切片 Slice
/*
	是一个拥有相同类型元素的可变长度的序列，是数组的引用（引用类型），包含地址，大小和容量，如下图，切片结构和内存分配：
	创建切片的两种方式（make 和 声明）
		使用 make() 函数构造切片（推荐）
			make([]T, size, cap)
			T（切片的元素类型）
			size（分配元素的数量）
			cap（预分配的元素数量）
			使用 make() 函数生成的切片一定发生了内存分配操作。但从数组或切片中获取新切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作
		声明切片
			var strList []string
			声明的空切片不为 nil 而是 []

	从数组或切片生成新的切片:
		slice [start:end]　　start：开始位置索引，end：结束位置索引　　（顾头不顾尾）
		两个位置索引都缺省时，与数组或切片本身等效
		两个位置索引都为0时，等效于空切片（用于清空元素）

	使用 append() 函数为切片添加元素
		内建函数 append() 函数可以为切片添加元素，每个切片会指向一片内存空间，这片空间可以容纳一定数量的元素，当空间不足时，切片就会进行“扩容”，扩容操作往往发生在 append() 调用时
		切片在扩容时，容量的扩展规律按容量的 2 倍扩充，例如 1,2,4,8,16...

		往一个切片中不断添加元素的过程，类似于公司搬家。公司发展初期，资金紧张，人员很少，所以只需要很小的房间即可容纳所有的员工。
		随着业务的拓展和收入的增加就需要扩充工位，但是办公地的大小是固定的，无法改变。因此公司选择搬家，每次搬家就需要将所有的人员转移到新的办公点
			1.员工和工位就是切片中的元素。
			2.办公地就是分配好的内存。
			3.搬家就是重新分配内存。
			4.无论搬多少次家，公司名称始终不会变，代表外部使用切片的变量名不会修改。
			5.因为搬家后地址发生变化，因此内存“地址”也会有修改

		切片添加多个元素（直接往后面写就行了， 用逗号隔开）

		切片中添加切片

	copy切片
		内建函数 copy() ，可以将一个切片的数据复制到另外一个空间, 需要先分配足够的内存!
		使用格式：copy(destSlice, srcSlice)

	从切片中删除数据
		并没有什么函数来提供删除切片元素，需要切片本身的特性来删除元素
		切片的删除过程，本质是以被删除元素为分界点，将前后两个部分的内存重新连接起来

	切片排序
		导入 sort 包
			sort.Ints　　       整数类型排序
			sort.Strings         字符串排序
			sort.Float64s      浮点数排序
 */



func main() {
	//make([]T, size, cap)
	var sliceA []int = make([]int, 2, 10)
	fmt.Println(sliceA)
	fmt.Println(len(sliceA)) //2

	//声明新的切片, 空切片不为 nil
	var sliceB []string
	fmt.Println(sliceB)      //[]
	fmt.Println(len(sliceB)) //0

	//声明新的切片，并填充内容（字面量方式）
	var sliceTest []string = []string {"a", "b", "c"}
	sliceTest = append(sliceTest, "d")
	fmt.Printf("sliceTest type: %T \n", sliceTest)  //sliceTest type: []string

	//从数组中获取切片
	var arrayC [3]int = [3]int{1, 2, 3}
	var sliceC []int = arrayC[:]
	fmt.Printf("sliceC:%v: sliceC type:%T \n", sliceC, sliceC) //sliceC:[1 2 3]: sliceC type:[]int
	//从切片中获取切片
	var sliceD []int = sliceC[0:0]
	fmt.Println("sliceD: ", sliceD) //[] 清空元素

	//使用 append() 函数为切片添加元素
	sliceD = append(sliceD, 0)
	fmt.Println(sliceD)
	//切片扩容规律
	var numSlice []int
	for i := 0; i <= 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("len:%d, cap:%d, point:%p \n", len(numSlice), cap(numSlice), numSlice) //len() 返回切片内元素数量，cap() 函数返回切片容量大小
	}

	//切片添加多个元素（直接往后面写就行了， 用逗号隔开）
	sliceD = append(sliceD, 1, 2, 3)
	fmt.Println(sliceD) //[0 1 2 3]

	//切片中添加切片
	var sliceE []int = make([]int, 2, 10)
	sliceD = append(sliceD, sliceE...)
	fmt.Println(sliceD) //[0 1 2 3 0 0]

	//copy切片 	copy(destSlice, srcSlice)  需要先分配足够的内存
	//var sliceF []int
	//var sliceF []int = make([]int, 2)
	var sliceF []int = make([]int, len(sliceD))
	copy(sliceF, sliceD)
	fmt.Println(sliceF)

	//切片中删除数据, 例如删除 "c", 代码的删除过程，本质是以被删除元素为分界点，将前后两个部分的内存重新连接起来
	var sliceG []string = []string{"a", "b", "c", "d", "e"}
	sliceG = append(sliceG, "f")
	fmt.Printf("sliceG point:%p \n", sliceG)
	var sliceH = sliceG[:2]
	fmt.Printf("sliceH point:%p \n", sliceH)
	var sliceI = sliceG[3:]
	fmt.Printf("sliceI point:%p \n", sliceI)
	sliceG = append(sliceH, sliceI...)
	fmt.Println(sliceG) //[a b d e]
	fmt.Printf("sliceG point:%p \n", sliceG)

	//切片排序
	var sliceInt []int = []int{4, 3, 5, 2, 6, 1}
	var sliceString []string = []string{"abc", "ab", "de", "d", "cd", "c"}
	var sliceFloat []float64 = []float64{2.34, 4.3, 1.23, 5.6, 3.2}
	sort.Ints(sliceInt)
	sort.Strings(sliceString)
	sort.Float64s(sliceFloat)
	fmt.Println(sliceInt)
	fmt.Println(sliceString)
	fmt.Println(sliceFloat)
}