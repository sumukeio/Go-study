package main

import (
	"fmt"
	"time"
)

type User struct {
	id         int
	Score      float32
	name, addr string
	enrollment time.Time // time包下的Time本质是个结构体，嵌套在User内，嵌套结构体
}

// 方法
func (user User) hello(man string) string {
	return "hello " + man + ", I'm " + user.name
}

// 不同于上面的方法，它是一个函数
func hello(user User, man string) string {
	return "hello " + man + ", I'm " + user.name
}

// 如果用不上结构体里的属性，可以用_替换结构体名
func (_ User) hello1(man string) string {
	return "I'm " + man
}

func main() {
	var ws User
	fmt.Printf("id=%d, Score=%g, name=[%s], addr=[%s], enrollment=%v\n", ws.id, ws.Score, ws.name, ws.addr, ws.enrollment)

	// 单独看看enrollment
	fmt.Printf("enrollment is %v\n", ws.enrollment)  //打印结构体属性
	fmt.Printf("enrollment is %#v\n", ws.enrollment) //打印结构体属性与结构体名
	fmt.Printf("enrollment is %+v\n", ws.enrollment) //打印结构体属性
	fmt.Println(ws.enrollment)
	fmt.Println(ws) //{0 0   {0 0 <nil>}} 嵌套结构体

	ws = User{Score: 100, name: "Tom"}
	ws.Score = 94.3
	ws.enrollment = time.Now()

	a := ws.id + 24
	fmt.Printf("a=%d\n", a)

	// 成员函数——方法
	fmt.Println(ws.hello("zcy"))
	fmt.Println(hello(ws, "zcy"))
}
