package main

import (
	"fmt"
)

func main() {
	var arr [2]string = [2]string{}
	fmt.Printf("%T \n", arr)
	fmt.Println(&arr)
	fmt.Println(len(arr))

}
