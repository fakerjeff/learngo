package main

import "fmt"

// go语言参数为值传递
func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap1(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)
	a, b = swap1(a, b)
	fmt.Println(a, b)
}
