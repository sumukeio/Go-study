package main

import "fmt"

// 想执行，生成可执行文件就必须改为main包

/*
	演示结构体中的匿名成员与结构体指针
*/

// Student 结构体中的匿名成员
type Student struct {
	Name string
	Age  int
	byte // 匿名成员
	bool
	string
}

// NewStudent 结构体指针  ——  学生实例的构造函数
func NewStudent(name, city string, age int) *Student {
	// 四种书写方式(Student内部可以只有字段值，但必须全写上而且按顺序)

	//var s Student
	//s.Name = name
	//s.Age = age
	//s.string = city
	//
	//return &s

	//s := Student{Name: name, string: city, Age: age}
	//return &s

	//s := &Student{Name: name, string:city, Age: age}
	//return s

	return &Student{Name: name, string: city, Age: age}
}

// 结构体的成员方法  参数不同  传入结构体的拷贝或结构体的指针
func (s Student) say() { // 传入的是结构体的拷贝
	s.Name = "123"
	fmt.Printf("say()内部的Name为 ")
}

func (s *Student) sleep() { // 传入的是结构体的指针
	s.Name = "123"
}

func main() {
	// 给结构体中的匿名成员赋值
	var s Student
	s.Name = "zcy"
	s.Age = 18
	s.string = "浙江"
	s.byte = 1
	s.bool = true

	// 结构体指针
	t := Student{Name: "zcy", Age: 19, string: "wuhan"}
	tPtr := &t
	fmt.Printf("Name: %s, Age: %d, city: %s\n", tPtr.Name, tPtr.Age, tPtr.string)
	tPtr.say() // 结构体指针，既可以访问结构体成员变量也可以访问成员方法

	// 看看结构体指针访问的字段的地址
	fmt.Printf("addr of Name is %p\n", &(tPtr.Name))

	t2 := *tPtr // 指针解析
	t2.say()
	fmt.Printf("Name: %s, Age: %d, city: %s\n", t2.Name, t2.Age, t2.string)
	// 结构体实例也可以，只不过访问的是结构体成员值的拷贝

	// 看看结构体实例访问的字段的地址
	fmt.Printf("addr of Name is %p\n", &(t2.Name))

	// 打印看看指针的结构体
	fmt.Println(tPtr)
	fmt.Println(t2)
	fmt.Printf("%p %v\n", tPtr, tPtr)

	// 调用结构体的成员方法
	st := Student{Name: "zcy", Age: 19, string: "chengdu"}

	st.say()
	fmt.Printf("Name = %s\n", st.Name) // zcy
	st.sleep()
	fmt.Printf("Name = %s\n", st.Name) // 123
}
