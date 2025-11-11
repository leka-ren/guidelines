package main

import (
	"fmt"
	"slices"
)

func main() {

	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.All(arr))
}
