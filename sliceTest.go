package main

import "fmt"

func sliceTest() {
	a := [7]int{0, 1, 2, 3, 4, 5, 6}
	sliceA := a[2:6]
	sliceB := a[2:]
	sliceC := a[:]
	fmt.Println(sliceA, sliceB, sliceC)
}

func updateSlice(slice []int) {
	slice[0] = 100
}

func appendSlice(slice []int, value int) []int {
	slice = append(slice, value)
	return slice
}

func main() {
	sliceTest()
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := arr1[0:5]

	updateSlice(slice1)
	fmt.Println(slice1)
	fmt.Println(arr1)
	slice1 = appendSlice(slice1, 102)
	fmt.Println(slice1)
	fmt.Println(arr1)
	slice1 = appendSlice(slice1, 103)
	fmt.Println(slice1)
	fmt.Println(arr1)
	slice1 = appendSlice(slice1, 104)
	fmt.Println(slice1)
	fmt.Println(arr1)
	slice1 = appendSlice(slice1, 105)
	fmt.Println(slice1)
	fmt.Println(arr1)

}
