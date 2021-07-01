package main

import (
	"fmt"

	"sort/bubbleSort"
	"sort/quickSort"
)

func main() {
	arr := []int{45, 78, 4, 9, 54, 34, 23, 66, 10}

	fmt.Println(arr)
	fmt.Println(quickSort.QuickSort(arr))
	fmt.Println(bubbleSort.BubbleSort(arr))
}
