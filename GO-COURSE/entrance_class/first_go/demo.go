package main // main函数必须得在main包中

import (
	"fmt"                                          //从标准库中引入包
	"go_project/src/GO-COURSE/entrance_class/util" //从工作目录中引入包，必须是完整目录
	//从第三方依赖库引入包,还没引入过
)

func main() { // go程序的唯一入口，开始执行的地方
	var a int = 4
	b := 8
	c := util.Add(a, b)
	fmt.Printf("c = %d = %d + %d \n", c, a, b)

	//x := []float64{1.1, 2.2, 5.0} //定义一个切片，名为x

}
