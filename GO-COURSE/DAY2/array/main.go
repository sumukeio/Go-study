package main

import "fmt"

/*
	2023年1月12日10:58:35
	主题：数组与指针
	内容：
	数组的初始化
	数组的遍历
	访问数组的值
	通过指针修改数组的值
	通过指针获取地址修改基本类型变量的值

*/

func basic() {
	var arr1 = [5]int{}
	arr2 := [5]int{1, 2, 3}
	var arr3 = [5]int{0: 2, 2: 3} //2 0 3 0 0
	var arr4 = [...]int{1, 3, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr3[0], arr3[2])
	arr3[0] = 9
	fmt.Println(arr3[0])
	fmt.Println(arr3[len(arr3)-1])
}

// Mean 求数组平均值
func Mean(arr [5]int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("数组平均值为%d \n", sum/len(arr))
}

// Mean2 求数组平均值
func Mean2(arr [5]int) float64 {
	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}
	return sum / float64(len(arr))
}

// 传入的参数为数组的指针，通过指针直接获取到数组的地址并修改其元素
func arrPoint(crr *[5]int) {
	fmt.Printf("传入函数的arr的地址为%p\n", crr)
	fmt.Println(crr[0])
	crr[0] += 10
	fmt.Println(crr[0])
}

// 传入的参数为数组
func arrPoint2(crr [5]int) {
	fmt.Printf("传入函数的arr的地址为%p\n", &crr)
	fmt.Println(crr[0])
	crr[0] += 10
	fmt.Println(crr[0])
}

// 再次验证传入数组是原数组值的拷贝
func forRange() {
	arr := [...]int{1, 3, 5, 7, 9}
	for i, ele := range arr { //ele是arr值的拷贝，修改ele并不会影响arr的值
		fmt.Printf("%d %d\n", i, arr[i])
		ele += 8
		fmt.Printf("%d %d\n", i, arr[i])
	}
}

func forRange1() {
	arr := [...]int{2, 4, 6, 8, 10}
	for i, ele := range arr { // ele是arr值的拷贝，修改arr的值当然不会影响它的拷贝
		fmt.Printf("%d %d\n", i, ele)
		arr[i] += 10
		fmt.Printf("%d %d\n", i, ele)
		fmt.Printf("%d %d\n", i, arr[i])
	}
}

func main() {
	basic()
	arr := [5]int{23, 56, 78, 39, 82}
	Mean(arr)
	fmt.Println(Mean2(arr))
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])

	fmt.Println("==============分界线==============")
	fmt.Printf("main函数内arr的地址为%p\n", &arr) //打印出arr的地址
	arrPoint(&arr)                                //为什么修改了第0位的值，最后打印出来还是原来的值呢？
	fmt.Println(arr[0])                           //因为赋给arrPoint的是arr数组的拷贝，把备份的值修改了，原来的值没被修改

	fmt.Println("==============分界线==============")
	fmt.Printf("main函数内arr的地址为%p\n", &arr)
	arrPoint2(arr)
	fmt.Println(arr[0])

	fmt.Println("==============分界线==============")
	forRange()
	fmt.Println("==============分界线==============")
	forRange1()

	// 指针
	var a int = 9
	var b *int
	b = &a // 取地址
	fmt.Printf("%d %p\n", a, &a)
	*b = 3 //解析指针，修改值
	fmt.Printf("%d %p\n", a, &a)

}
