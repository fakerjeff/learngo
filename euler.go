package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	// 验证欧拉公式
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func GouGu(a, b int) {
	// Go中只有强制类型转换，没有隐式类型转换
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Println(int(c))
}

func main() {
	euler()
	GouGu(3, 4)
}
