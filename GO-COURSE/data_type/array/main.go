package main

import "fmt"

// 数组初始化
var arr1 = [5]int{}                // 指定了长度且不可改变，初始化时就已经分配好了默认值0
var arr2 = [5]int{1, 2}            // 只赋了前两个值，其他为默认值0
var arr3 = [5]int{1: 2, 2: 3}      // 给下标为1和2的元素赋值，其他为默认值0
var arr4 = [...]int{3, 4, 5, 6, 7} // 根据{}里元素个数推算数组长度
var arr5 = [...]struct {
	name string
	age  int
}{{"Tom", 18}, {"Jim", 20}} //数组的元素为结构体类型，结构体后的{}内存放了两个元素

func main() {
	// 数组遍历方法1
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("index: %d, value: %d\n", i, arr4[i])
	}

	// 数组遍历方法2
	for i, v := range arr4 {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}
