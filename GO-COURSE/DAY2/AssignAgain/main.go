package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
	需求：创建一个初始长度为0、容器为10的int型切片，调用rand.Intn(128)100次，
	往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素
	注释：rand是个包，Intn(128)是随机产生一个0-128的随机数
	排除重复元素，看看剩下有多少个元素
*/
// 把功能封装成函数，让用户使用方便

func GenSlice(n int) []int {
	// 随机数生成器, source是一个种子，通过变换种子，可以使得每次生成的随机序列不同
	source := rand.NewSource(100) // 输入不同的数，获得不同的种子
	rander := rand.New(source)

	sli := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		sli = append(sli, rander.Intn(128)) // rand.Intn()这种方式，每次生成的随机序列都一样
	}
	return sli
}

func CountUniq(arr []int) int {
	// 计算不重复的元素个数
	m := make(map[int]bool, len(arr))
	for _, ele := range arr {
		m[ele] = true
	}
	return len(m)
}

/*
需求：实现一个函数func arr2string(arr []int)string，比如
输入[]int{2, 4, 6}，返回“2 4 6”，输入的切片长度不一。
提示：用stringbuilder做字符串拼接
*/
func arr2string(arr []int) string {
	var str strings.Builder
	for _, ele := range arr {
		str.WriteString(strconv.Itoa(ele))
		str.WriteString(" ")
	}
	return strings.Trim(str.String(), " ")
}

func main() {
	sli := GenSlice(100)
	fmt.Printf("互不相同的元素个数为%d\n", CountUniq(sli))

	str := []int{2, 4, 6}
	fmt.Println(arr2string(str))
}
