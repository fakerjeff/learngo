package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}

func readFile1() {
	const filename = "abc.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func eval(opt string, a, b int) int {
	if opt == "+" {
		return a + b
	} else if opt == "-" {
		return a - b
	} else if opt == "*" {
		return a * b
	} else if opt == "/" {
		return a / b
	} else {
		panic("不支持的操作")
	}
}

func eval1(opt string, a, b int) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("不支持的操作")
	}
}

func eval2(opt string, a, b int) int {
	switch {
	case opt == "+":
		return a + b
	case opt == "-":
		return a - b
	case opt == "*":
		return a * b
	case opt == "/":
		return a / b
	default:
		panic("不支持的操作")
	}
}

func main() {
	fmt.Println(eval("+", 3, 4))
	fmt.Println(eval("-", 3, 4))
	fmt.Println(eval("*", 3, 4))
	fmt.Println(eval("/", 3, 4))
	//fmt.Println(eval("#", 3, 4))
	readFile()
	readFile1()
}
