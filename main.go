package main

import (
	"fmt"
	// "sort/bubbleSort"
	// "sort/insertionSort"
	// "sort/quickSort"
	// "sort/selectionSort"
	// "sort/hillSort"
	// "sort/mergeSort"
	"sort/heapSort"
)

func main() {
	arr := []int{45, 78, 4, 9, 54, 34, 23, 66, 10}

	fmt.Println(arr)
	// fmt.Println(quickSort.QuickSort(arr))
	// fmt.Println(bubbleSort.BubbleSort(arr))
	// fmt.Println(selectionSort.SelectionSort(arr))
	// fmt.Println(insertionSort.InsertionSort(arr))
	// fmt.Println(hillSort.HillSort(arr))
	// fmt.Println(mergeSort.MergeSort(arr))
	fmt.Println(heapSort.HeapSort(arr))
}
