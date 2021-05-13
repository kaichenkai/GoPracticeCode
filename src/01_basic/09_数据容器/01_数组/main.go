package main

import (
	"fmt"
)

//数组是值类型，因此改变副本的值，不会影响到本身
/*
	数组的定义：var 变量名 [元素数量] T
		变量名（符合标识符要求即可）
		元素数量（整型，可以是const中的值）
		T（可以是任意基本类型，包括数组本身，当类型为数组时，可以实现多维数组）
		var a [5]int 和 var a [10]int 是不同的类型
 */


//多维度数组


//数组排序
/*
不能对数组进行排序
 */


func main() {
	//定义一个字符串数组，然后赋值
	var team [3]string
	team[0] = "zero"
	team[1] = "one"
	team[2] = "two"
	//team[3] = "three"  //invalid array index 3 (out of bounds for 3-element array)

	//也可以在声明时进行元素赋值
	var team2 [3]string = [3]string{"zero", "one"}  //没有赋值是基础类型默认值
	fmt.Println(team2)
	fmt.Printf("==%s==, %T \n", team2[2], team2[2])  //====, string

	//遍历数组 range
	for index, value := range team {
		fmt.Println(index, value)
	}
	//遍历数组 index
	for i:=0; i<len(team); i++ {
		fmt.Println(i, team[i])
	}

	//多维度数组
	var manyArray [2][5]int = [2][5]int{{1,2,3,4,5}, {6,7,8,9,0}}
	for rowIndex, rowValue := range manyArray {
		for columnIndex, columnValue := range rowValue {
			fmt.Printf("%d %d -> %d \n", rowIndex, columnIndex, columnValue)
		}
	}

	//数组排序
	//不能对数组进行排序
}