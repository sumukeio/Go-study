package main

import (
	"fmt"
)

// slice初始化
func sliceInit() {
	s2d := [][]int{ //二维切片
		{1},
		{3, 5, 7}, // 二维切片每列长度可以不相等
	}
	fmt.Println(len(s2d))    // 2
	fmt.Println(len(s2d[0])) // 1
	fmt.Println(len(s2d[1])) // 3

	var s []int
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	s1 := make([]int, 3, 5)
	fmt.Printf("s1: %v. len: %d, cap: %d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 3)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
}

// 追加元素
func sliceAppend() {
	s := make([]int, 3, 5) //创建一个切片，类型为int，len为3，cap为5
	s = append(s, 100)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s)) // 4 5,并未充满
	s = append(s, 101)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s)) // 5 10，预留空间被用完

	s = append(s, 102)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s)) // 6 10，cap容量扩容为原先的两倍，然后进行追加操作
	fmt.Println(s)
}

// 用函数直观呈现切片扩容
func funCap() {
	s := make([]int, 0, 5)
	prepCap := cap(s)           // 之前的容量
	for i := 0; i < 2000; i++ { //遍历1024次，每次追加一个0
		s = append(s, 0)
		currCap := cap(s) // 扩容之后的容量
		// 由于扩容到原先的两倍后，之后不断追加，到下一次扩容还有一段距离，所以，只有当之后的容量大于原先的容量时我们才打印一次容量对比
		if currCap > prepCap {
			fmt.Printf("之前的容量：%d => 现在的容量：%d\n", prepCap, currCap)
			prepCap = currCap //并把之前的容量替换成现在的容量，否则，一直都是最初未扩容时的容量——5
		}
	}
}

func main() {
	sliceInit()
	fmt.Println("========================")
	sliceAppend()
	fmt.Println("========================")
	funCap()
}
