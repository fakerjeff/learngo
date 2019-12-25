package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func add(a, b int) (int, error) {
	return a + b, nil
}

func sub(a, b int) (int, error) {
	return a - b, nil
}

func multi(a, b int) (int, error) {
	return a * b, nil
}

func div(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, fmt.Errorf("divide by zero")
	}
}

func eval3(op string, a, b int) (int, error) {
	switch {
	case op == "+":
		return add(a, b)
	case op == "-":
		return sub(a, b)
	case op == "*":
		return multi(a, b)
	case op == "/":
		return div(a, b)
	default:
		return 0, fmt.Errorf("unsurport opreate")
	}
}

func eval4(op func(int, int) (int, error), a, b int) (int, error) {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("call method %s, with parameter %d, %d \n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	// 不定长参数
	sum := 0
	for i := range numbers {
		sum += i
	}
	return sum

}

func main() {

	if result, err := eval4(func(a, b int) (int, error) {
		return int(math.Pow(float64(a), float64(b))), nil
	}, 5, 2); err != nil {
		fmt.Printf("some bad things happend: %s", err)
	} else {
		fmt.Printf("the result is %d", result)
	}

}
