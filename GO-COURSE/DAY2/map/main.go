package main

import (
	"fmt"
)

func main() {
	// 声明一个数组
	var a = [5]int{1, 2, 3}
	fmt.Println(a)
	// 初始化map的四种方式
	var m map[string]int
	fmt.Println(len(m)) // 0

	// 反复敲，反复理解
	var mm map[string]int
	fmt.Println(mm)

	m1 := make(map[string]int) // 等价于 m1 := make(map[string]int, 0)
	fmt.Println(len(m1))       // 0

	// 反复敲，反复理解
	mm1 := make(map[string]int)
	fmt.Println(mm1)

	m2 := make(map[string]int, 200)
	fmt.Println(len(m2)) // 0，cap是200(虽说不支持)，len是0，因为只是申请了内存，还没有存入地址呢

	// 反复敲，反复理解
	mm2 := make(map[string]int, 300)
	fmt.Println(len(mm2))

	m3 := map[string]int{"A": 0, "B": 4}
	fmt.Println(len(m3)) // 2

	// 反复敲，反复理解
	mm3 := map[string]int{"A": 1, "B": 2}
	fmt.Println(len(mm3))

	// 跟敲
	m4 := make(map[string]int, 10) // 此时的len依旧为0，因为没有键值对存入
	m4["A"] = 1
	m4["B"] = 2
	m4["C"] = 3
	m4["D"] = 4
	m4["E"] = 5 // 此时的len为5

	mm4 := make(map[string]int)
	mm4["A"] = 1
	mm4["B"] = 2
	mm4["C"] = 3
	mm4["D"] = 4

	// 添加元素
	m3["C"] = 5
	fmt.Println(len(m3)) // 3

	// 删除元素
	delete(m3, "B")
	fmt.Println(len(m3)) // 2
	fmt.Printf("%v\n", m3)

	// 跟敲
	delete(mm3, "B")

	// 删除一个不存在的元素
	delete(m3, "D")      // 删除一个nil或者一个不存在的元素，结果就是删除一个noop(和没有这条语句一样，什么都没发生No operation)
	fmt.Println(len(m3)) // 2

	// 访问一个不存在的value
	value := m["D"]
	fmt.Println(value) // 0

	// 跟敲
	fmt.Println(mm3["B"])
	// 或者
	v := mm3["B"]
	fmt.Println(v)

	// 使用前判断key是否存在
	key := "L"
	v, ok := m[key] // 我不理解为什么可以用两个变量来接收？
	if ok {
		fmt.Println(v)
	} else {
		fmt.Printf("key %s 不存在\n", key)
	}
	// 以上区分方法应对的尴尬问题
	value = m3["A"]
	fmt.Println(value) // 0 因为m3这个map中“A”对应value就是0
	value = m3["E"]
	fmt.Println(value) // 0 此处的0是因为m3这个map中"E"不存在

	// 遍历map
	for key, value := range m3 {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	// 跟敲
	for key, value := range mm3 {
		fmt.Println(key, value)
	}

	// 验证map中的值拷贝
	for key, value := range m4 { // 此处的参数value是对原始value的拷贝 原始value分别是1 2 3 4 5
		value += 2 // 将m这个map中的value都加2
		fmt.Printf(" key: %s, 修改后的拷贝value为 %d, 修改后真实value为 %d\n", key, value, m4[key])
		// 3 1   4 2   5 3   6 4   7 5
	}
}
