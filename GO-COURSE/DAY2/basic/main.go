package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	var a int
	var b bool
	var c [5]int
	var u User
	// 每个变量，无论其值是基本类型还是复合类型，都有内存地址

	fmt.Printf("%p \n", &a) //取地址并打印
	fmt.Printf("%p \n", &b)
	fmt.Printf("%p \n", &c)
	fmt.Printf("%p \n", &u)
}
