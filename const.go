package main

import (
	"fmt"
	"math"
)

func traingleGougu() {
	// 定义常量时如果未指定常量类型，则常量类型不定，不会进行类型推定
	const a, b = 3, 4
	fmt.Println(math.Sqrt(a*a + b*b))
}

func enums() {
	// 定义枚举, 使用 iota表示枚举值自增（从 0 开始）
	// 使用_可以跳过自增值
	const (
		golang = iota
		python
		java
		_
		cpp
		javascript
	)
	fmt.Println(golang, python, java, cpp, javascript)
}

func capacity() {
	// 可以使用iota的表达式来定枚举值
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	traingleGougu()
	enums()
	capacity()
}
