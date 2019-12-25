package main

import "fmt"

// 函数外部不能使用 :=进行变量声明与赋值
// 函数外部的变脸作用域在package内部
var aa = 1
var ss = "hello world"
var bb = true

var (
	cc = 1
	dd = 2
	ee = 3
)

func variables() {
	// go中定义的变量，可以不赋初始值，其对应的值是该对应类型的默认值，如
	// int 为 0 string 为 ""字符串
	var a int
	var b string
	fmt.Printf("%d, %q\n", a, b)
}

func initVariables() {
	// 定义变量并赋初始值
	var a, b int = 1, 4
	var c string = "hello world"
	fmt.Println(a, b, c)
}

func variableTypeDeduction() {
	// 可以不声明变量的类型，使用变量推断来声明变量
	var a, b, c, d = 1, 2, "hello world", true
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	// 方法内可以使用 := 进行变量声明并赋值
	a, b, c, d := 1, 2, "hello world", false
	b = 6
	fmt.Println(a, b, c, d)
}

func main() {
	variables()
	initVariables()
	variableTypeDeduction()
	variableTypeDeduction()
}
