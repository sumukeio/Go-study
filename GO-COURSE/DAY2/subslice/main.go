package main

import "fmt"

func subSlice() {
	arr := make([]int, 3, 5) // 0 0 0 空 空
	crr := arr[0:2]          // 前闭后开区间，包含0和1
	crr[1] = 8               // 0 8 0 空 空
	fmt.Printf("arr[1]为：%d\n", arr[1])
	crr = append(crr, 9) // 0 8 9 空 空
	fmt.Printf("arr[2]为：%d\n", arr[2])

	crr = append(crr, 10, 11)
	fmt.Printf("arr为 %v", arr) // 0 8 9 尽管底层数组的值因为append扩容已经改变，但由于arr的len还是3，所以只显示了前三个
	//fmt.Printf("arr[3]: %d, arr[4]: %d\n", arr[3], arr[4])
	//上一行代码发生panic，因为arr的length还是3，虽然底层变成5了,越界访问导致报错
	//解决办法：把arr的length改成5

	// 继续执行append，发生内存分离
	crr = append(crr, 12)
	fmt.Printf("arr指向的内存地址为 %p, crr指向的内存地址为 %p\n", &arr, &crr)
	crr[1] = 5
	fmt.Printf("arr[1]为 %d\n", arr[1])

}

/*
	由于Go语言传参传的是参数值的拷贝。
	所以 ，传切片时会把{arrayPointer，len，cap}
	这三个字段拷贝一份进来；由于传的是底层数组的指针arrayPointer，
	所以可以直接修改底层数组里的元素，因为指针指向了底层数组
*/

func updateSlice(arr []int) {
	arr[0] = 100 // 接收切片并修改对应值
}

func updateSlice2(arr *[]int) {
	(*arr)[0] = 100
}

func updateArr(arr [3]int) {
	arr[0] = 100 // 接收数组并修改对应值
}

func main() {
	subSlice()
	fmt.Println("===================")

	// 切片传参试验
	crr := []int{3, 4, 5} //arrPointer len:3 cap:3
	updateSlice(crr)
	fmt.Printf("crr[0]= %d\n", crr[0]) // 100
	// 因为拷贝过去的值里有指针，指针指向了内存地址，改的就是那个内存地址里的值

	// 数组传参实验
	brr := [3]int{1, 2, 3} // len=cap=3
	updateArr(brr)
	fmt.Printf("brr[0]= %d\n", brr[0]) // 1
	// 因为传过去的是数组的拷贝，所以，原本的数组不受影响

	updateSlice2(&crr)
}
