package main

import (
	"fmt"
)

func printer(ch chan int){
	for data := range ch{
		if data == 0 {
			break
		}
		for i:=100; i>=0; i-- {
			fmt.Print("*" )
		}
		fmt.Println("" )

	}
	//返回数据输入端，打印完了
	fmt.Println("打印完了哈")
	ch <- 1
}

func main(){
	var ch chan int = make(chan  int)
	go printer(ch)
	//输送数据
	for i:=10; i>=0; i-- {
		ch <- i
	}

	//接收任意一个数据，如果接收到，表示写入完成
	<- ch
	fmt.Println("结束")
}

