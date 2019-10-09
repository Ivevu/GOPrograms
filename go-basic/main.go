package main

import (
	"fmt"
	"strconv"
)

// 无返回值 无参数
func noParamNoReturn() {
	fmt.Println("无返回值无参数")
}

// 无返回值的 带参参数
func printStr(s string) {
	fmt.Println(s)
}

// 带多个参数 无返回值
func noreturnWithParams(args ...int) {
	for index, value := range args {
		fmt.Println(index, value)
	}
}

// 带一个返回值
func withReturn() int64 {
	return 1
}

// 带多个返回值
func withreturns() (int, int) {
	return 1, 1
}

// 给返回值命名
func nameReturn() (r int) {
	return 1
}

// 值传递
func copyValue(t int) int {
	t = t + 1
	return t
}

// 地址（指针）传递
func copyAddress(p *int) int {
	*p = *p + 1
	return *p
}

// 函数作为参数
func useFuncAsParam(function func(int) int) int {
	return function(2)
}

// 匿名函数
var noNameFunc = func(s int) int {
	return 1
}

func main() {
	noParamNoReturn()
	printStr("无返回值的 带参参数")
	noreturnWithParams(1, 2)
	fmt.Println(strconv.FormatInt(withReturn(), 10) + "变成字符串了")
	fmt.Println(withreturns())
	r := nameReturn()
	fmt.Println(r)

	var num int = 11
	copyValue(num)
	fmt.Println(num)
	copyAddress(&num)
	fmt.Println(num)

	s := useFuncAsParam(copyValue)
	fmt.Println(s)
	noNameFunc(1)
}
