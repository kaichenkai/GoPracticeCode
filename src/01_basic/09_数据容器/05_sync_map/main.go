package main

import (
	"fmt"
	"sync"
)

//sync map
/*
	同时读写 map 的问题
	map 在并发情况下，只读线程是安全的，同时读写线程不安全
	运行时输出提示：并发的 map 读写，也就是说使用了两个并发函数不断地对 map 进行读和写而发生了竞态问题，map 内部会对这种并发操作进行检查并提前发现
	fatal error: concurrent map read and map

	引入 sync.Map
	并发读写时，一般的做法是加锁，但这样性能并不高，Go 在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map，与 map 不同，它不是以语言原生形态提供，而是在 sync 包下的特殊结构
	声明一个 sync.Map（不能使用 make 创建）

	sync.Map 特性：
		1.不能使用 map 的操作方式不同，是使用 sync.Map 的方法进行调用，Store 用来存储，Load 用来获取，Delete 用来删除
		2.遍历：使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数回调函数的返回值是 bool 类型，需要继续迭代时，返回 true，终止迭代时，返回 false
		3.没有提供获取数量的方法，替代方法是遍历时自行计算数量
		4.sync.Map 为了保证并发安全，会有一些性能损失，因此在非并发读写情况下，使用 map
*/

func main() {
	//下面来看下并发情况下读写 map 时会出现的问题： fatal error: concurrent map read and map
	//var mapA map[int]int = make(map[int]int)
	//go write(mapA)
	//read(mapA)

	//sync.Map
	var mapA sync.Map
	go syncWrite(mapA)
	syncRead(mapA)

	//遍历sync.Map
	rangeSyncMap()
}

func write(mapA map[int]int) {
	for {
		mapA[1] = 1
		fmt.Println("write")
	}
}
func read(mapA map[int]int) {
	for {
		_ = mapA[1]
		fmt.Println("read")
	}
}

func syncWrite(mapA sync.Map) {
	for {
		mapA.Store(1, 1)
		fmt.Println("sync write")
	}
}
func syncRead(mapA sync.Map) {
	for {
		mapA.Load(1)
		fmt.Println("sync read")
	}
}

//遍历 syncMap
func rangeSyncMap() {
	var syncMap sync.Map

	// 设置值
	syncMap.Store("hebei", "beijing")
	syncMap.Store("hubei", "wuhan")
	syncMap.Store(1, "shenzhen")
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		if key == "hebei" {
			return false
		} else {
			return true
		}
	})
}
