package main

import (
	"fmt"
)

func main() {

	var a [10]int // это массив(arrays) длиной 10 позиций (от 0 до 9). Размер не меняется
	a[1] = 1000
	slice := a[0:5] // это уже срез (slice), который является ЧАСТЬЮ массива

	fmt.Println(a)
	fmt.Println(a[1])
	fmt.Println(slice)

	slice[1] = 9

	fmt.Println(slice[1])
	fmt.Println(slice)
}
