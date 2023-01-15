package main

import "net/http"

// 起http服务

func boy(w http.ResponseWriter, r *http.Request) {
	// 接收ResponseWriter与Request指针
	// 从Request里拿到请求的各种参数，把响应准备好之后通过ResponseWriter写回给客户端
	w.Write([]byte("hello boy")) //函数需要传byte碎片，将其从字符串转为byte碎片
}

func girl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello girl"))
}

//handler源码
//type Handler interface {
// 	ServeHTTP(ResponseWriter, *request) //该接口必须实现ServeHTTP方法。该方法接收这两个参数
//}

func main() {
	// 创建两个路由
	//此处接收的是字符串(请求的路径)与handler，传的却是函数，所以要转化成handler
	// 原理：客户端请求这个路径，通过这个函数进行响应
	http.Handle("/boy", http.HandlerFunc(boy))
	http.Handle("/girl", http.HandlerFunc(girl))

	// 监听
	http.ListenAndServe(":5656", nil) //端口
}
