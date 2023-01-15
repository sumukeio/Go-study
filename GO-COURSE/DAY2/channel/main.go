package main

import "fmt"

func main() {
	var ch chan int                            // 三大引用类型的默认值都是nil，但可以调用长度，结果为0
	fmt.Printf("ch is nil -> %t\n", ch == nil) // %t是什么
	fmt.Printf("len of ch is %d\n", len(ch))
	var m map[string]int
	fmt.Printf("m is nil -> %t\n", m == nil)
	fmt.Printf("len of m is %d\n", len(m))
	var s []int
	fmt.Printf("s is nil -> %t\n", s == nil)
	fmt.Printf("len of s is %d\n", len(s))

	// 初始化
	ch = make(chan int, 10)
	ch1 := make(chan int)
	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))
	fmt.Printf("len of ch1 is %d\n", len(ch1))
	fmt.Printf("cap of ch1 is %d\n", cap(ch1))

	// 插入数据
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))
	// 能一次插入多个数据吗？不能
	// 插入越界了怎么办？
	//ch1 <- 1
	//ch1 <- 2 // fatal error: all goroutines are asleep - deadlock!

	// 完整的验证一遍，如果插入越界了，是覆盖其中的元素还是报错
	ch2 := make(chan int, 5)
	for i := 0; i < cap(ch2); i++ {
		ch2 <- 1
	}
	fmt.Printf("len of ch2 is %d\n", len(ch2))
	fmt.Printf("cap of ch2 is %d\n", cap(ch2))
	//ch2 <- 1 // 阻塞在这行代码，无法执行，因为管道已经塞满了，又没有其他协程把管道里的元素读走
	//fmt.Printf("len of ch2 is %d\n", len(ch2))
	//fmt.Printf("cap of ch2 is %d\n", cap(ch2))
	// 然后就报一个死锁出来，报一个panic
	// fatal error: all goroutines are asleep - deadlock!

	// 初始化+赋值+遍历
	var ch3 chan int
	ch3 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch3 <- 1
	}
	<-ch3
	ch3 <- 2
	close(ch3) // 如果遍历前不关闭就直接报错
	for ele := range ch3 {
		fmt.Println(ele)
	}
	fmt.Println("-----------------")
	fmt.Printf("len of ch3 is %d\n", len(ch3)) // 0 遍历的同时把元素取走

	// 上述的forR循环操作相当于
	L := len(ch3) // 如果不专门赋给一个值，让它固定，那循环条件里的长度就会变，就不是循环len(ch3)次了
	for i := 0; i < L; i++ {
		v := <-ch
		fmt.Println(v) // 将元素取出并打印
	}
	// 用forI遍历不需要先close，而用forR遍历必须先close，因为它的本质是forI+取出
}
