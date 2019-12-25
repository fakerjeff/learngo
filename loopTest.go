package main

import (
	"fmt"
	"strconv"
)

func onePlusToHundreds() int {
	sum := 0
	// 循环基本写法
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func convertToBin(n int) string {
	result := ""
	if n < 0 {
		panic("非法参数")
	}
	if n == 0 {
		return "0"
	}
	// 循环省略初始条件
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func convertToBin1(n int) string {
	result := ""
	if n < 0 {
		panic("非法参数")
	}
	if n == 0 {
		return "0"
	}
	// 循环省略 初始条件和递增表达式
	for n > 0 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
		n /= 2
	}

	return result

}

func unlimitedLoop() {
	// go语言没有while
	for {
		fmt.Println("I'm a unlimited loop")
	}
}

func main() {
	fmt.Println(onePlusToHundreds())
	fmt.Println(convertToBin(78654))
	fmt.Println(convertToBin1(78654))
	unlimitedLoop()
}
