package main

import (
	"time"
)

type User struct {
	id         int
	Score      float32
	name, addr string
	enrollment time.Time // time包下的Time本质是个结构体，嵌套在User内，嵌套结构体
}
