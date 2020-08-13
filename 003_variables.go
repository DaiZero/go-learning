package main

import "fmt"

func main() {

	// var 声明变量，有赋值的时候可不声明类型，go可以自动推断变量类型
	var a = "001"
	fmt.Println(a)

	//一次性声明多个变量,并赋值
	var b, c int = 1, 2
	fmt.Println(b, c)

	//声明后却没有给出对应的初始值时，变量将会初始化为该类型的默认值
	var d int
	fmt.Println(d)

	// `:=` 语法是声明并初始化变量的简写，例如
	// 这个例子中的 `var e string = "ddd"`。
	e := "ddd"
	fmt.Println(e)
}
