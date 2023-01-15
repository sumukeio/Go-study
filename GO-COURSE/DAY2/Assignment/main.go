package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	/*
		需求：创建一个初始长度为0、容器为10的int型切片，调用rand.Intn(128)100次，
		往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素
		注释：rand是个包，Intn(128)是随机产生一个0-128的随机数
		排除重复元素，看看剩下有多少个元素
	*/
	//s := make([]int, 10)
	//for i := 0; i < 100; i++ {
	//	s = append(s, rand.Intn(128))
	//}
	//s1 := make(map[int]int, 100)
	//for index, value := range s {
	//	s1[index] = value
	//}
	//// 互不相同的value，key都是互不相同的呀
	//num := len(s)
	//for i := 0; i < num; i++ {
	//	// 找到重复元素，然后将其从map中移出，算剩下的len
	//	ele1 := s[i]
	//	for j := 0; j < num; j++ {
	//		ele2 := s[j]
	//		if ele1 == ele2 {
	//			delete(s1, s1[j])
	//		}
	//	}
	//}
	//fmt.Printf("不重复的元素个数为%d\n", len(s1))
	//
	//sli1 := []int{2, 4, 6}
	//fmt.Println(arr2string(sli1))
	//
	//var str string
	//fmt.Println(str)

	// 张朝阳的讲解
	arr := GenSlice(100) // 传入100个值
	cnt := CountUniq(arr)
	fmt.Printf("uniq ele count %d\n", cnt)

	// 第二题
	// 要考虑的边界情况问题：如果传入一个nil会怎么样？不会报错
	// 传入一个空切片呢？也没有问题
	s := Concat(nil)
	fmt.Println(s)
	s1 := Concat([]int{})
	fmt.Println(s1)
	s2 := Concat(arr)
	fmt.Printf("[%s]\n", s2)
}

/*
需求：实现一个函数func arr2string(arr []int)string，比如
输入[]int{2, 4, 6}，返回“2 4 6”，输入的切片长度不一。
提示：用stringbuilder做字符串拼接
*/
//func arr2string(arr []int) string {
//	var str string
//	fmt.Println(str)
//	for i := 0; i < len(arr); i++ {
//		str = str + " " + string(arr[i])
//	}
//	return str
//}

// GenSlice 张朝阳的讲解
// 尽可能都把功能封装成函数，而不是直接在main里写
/*
	需求：创建一个初始长度为0、容器为10的int型切片，调用rand.Intn(128)100次，
	往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素
	注释：rand是个包，Intn(128)是随机产生一个0-128的随机数
	排除重复元素，看看剩下有多少个元素
*/
// 	创建切片并往里追加n个元素
func GenSlice(n int) []int { // 接收元素个数
	//var arr []int
	//arr = make([]int, 0, 10)
	arr := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(128))
	}
	return arr // 容量肯定会扩充到很大，len扩充到100
}

// 检验切片中互不相同的元素个数
func CountUniq(arr []int) int {
	m := make(map[int]bool, len(arr)) // 创建一个index为int，value为bool的map，len为arr的len
	fmt.Printf("m = %v\n", m)         // 一个空的map
	// map的键为int类型，值为bool类型
	// 我想的是把原先切片的值传给map，而他是直接创建一个map
	for _, ele := range arr { // 遍历切片里的元素，int类型
		m[ele] = true
		/* 把切片的元素作为map的下标，只要是有这个元素，就将值标记为true，
		这样哪怕重复，也不会改变值。然后统计map的len，就知道不重复的元素个数
		 map的下标只是一个数而已，不代表len的多少，它是键值对，不是数组 */
	}
	return len(m)
}

// 第二题-张朝阳讲解
// 拼接函数
func Concat(arr []int) string {
	// 如果要拼接的元素很多，那用+的效率比较低，这种方式会导致大量的string创建、销毁和内存分配
	// fmt.Sprintf  strings.builder
	var sb strings.Builder //我的理解：sB是个工具，它可以调用WriteString函数，接收变量，写入字符串
	for _, ele := range arr {
		sb.WriteString(strconv.Itoa(ele)) // 把整型元素转化成字符串后写入字符串
		//sb.WriteString(strconv.FormatInt(int64(ele), 10)) //第二种将整型元素转化成字符串的方式，传入的必须是int64类型，显示为10进制
		sb.WriteString(" ")
	}
	return strings.Trim(sb.String(), " ") // 把最后多出来的空格删除     sb.String()是把string拿出来
	// strings.Trim的作用是删除字符串前后的，你给出的字符

	// 采用子切片的形式删除末尾多余的空格
	//s := sb.String()
	//return s[0 : len(s)-1]
}
