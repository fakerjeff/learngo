package main

import (
	"fmt"
)

func arrayTest() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [4][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {5, 4, 3, 2, 1}, {10, 9, 8, 7, 6}}
	// default value [0, 0, 0, 0, 0]
	c := [5]int{}
	var d [5]int
	var e = [...]int{1, 2, 3}
	var f [4][5]int
	fmt.Printf("a:%d\n", a)
	fmt.Printf("b:%d\n", b)
	fmt.Printf("c:%d\n", c)
	fmt.Printf("d:%d\n", d)
	fmt.Printf("e:%d\n", e)
	fmt.Printf("f:%d\n", f)
}

func loopThroughArrays(arr [4]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i := range arr {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		// 同时遍历 index和value
		fmt.Printf("%d:%d\n", i, v)
	}
	// 省略索引
	for _, v := range arr {
		fmt.Println(v)
	}
}
func variablePass(array1 [5]int) {
	// [5]int 和 [10]int是不同的类型
	array1 = [5]int{5, 4, 3, 2, 1}
	array1[0] = 100
	fmt.Println("calling variable pass")
	fmt.Println(array1)
}

func arrPtr(aPtr *[5]int) {
	// 使用指针可以直接改变数组的值
	// 不需要使用 *aPtr[1] = 100, 使用 *aPtr会报错
	//*aPtr[1] = 100  Invalid indirect of 'aPtr[1]' (type 'int')
	aPtr[1] = 100
}

// 数组是值类型
func main() {
	arrayTest()
	var array0 [5]int
	fmt.Printf("array0:%d\n", array0)
	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("before variable pass")
	fmt.Println(array1)
	variablePass(array1)
	fmt.Println("after variable pass")
	fmt.Println(array1)
	loopThroughArrays([4]int{1, 2, 3, 4})

	array2 := [5]int{}
	arrPtr(&array2)
	fmt.Println(array2)
}
