package main

import (
	"fmt"
)

func reverse(ar []int) {
	var l = len(ar)
	for i := 0; i < l/2; i++ {
		ar[i], ar[(l-1)-i] = ar[(l-1)-i], ar[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	reverse(arr)
	fmt.Println(arr)
}
