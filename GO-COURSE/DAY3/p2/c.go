package main

import "fmt"

/*
	演示匿名结构体与自定义类型
*/

func main() {
	var user User
	user.name = "abuzabi"
	fmt.Println(user.name)
	user.id = 1

	// 演示匿名结构体
	var abc struct { // abc接收了一个匿名结构体，可以直接访问元素并给它赋值
		Name string
		Age  int
	}
	abc.Name = "zcy"

	// 等价形式——更复杂
	type People struct {
		Name string
		Sex  byte
	}
	var bcd People // bcd的类型是People结构体，bcd是People类的实例
	bcd.Sex = 1

}

type People struct {
	Name string
	Sex  byte
}

// 自定义类型
type UserMap map[int]People

func (um UserMap) Get(id int) People {
	return um[id]
}
