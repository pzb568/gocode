package main

import (
	"fmt"
)

// 变量声明

// 第一种方法
/*
var a = 12
var b = 25
var c = 33
*/
// 第二种方法
var a, b, c = 5, 8, 9

func main() {

	fmt.Println("a加c除以b乘以e等于", a+c/b)
}
