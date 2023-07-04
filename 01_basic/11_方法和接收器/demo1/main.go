package main

import "fmt"

//飞行类动物
type Flying struct {}
//行走类动物
type Walkable struct {}

func (f Flying) fly(){
	fmt.Println("can fly")
}

func (w Walkable) walk(){
	fmt.Println("can walk")
}

//人是行走类动物
type Person struct {
	Walkable
}
//鸟竟是飞行类也是行走类
type Bird struct {
	Flying
	Walkable
}


func main() {
	//创建人类
	var person *Person = new(Person)
	person.walk()

	//创建鸟类
	var bird *Bird = new(Bird)
	bird.fly()
	bird.walk()
}
