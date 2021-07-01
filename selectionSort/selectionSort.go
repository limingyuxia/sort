package selectionSort

// https://www.runoob.com/w3cnote/selection-sort.html

/*
	1. 假设第一个元素是最小的元素
	2. 遍历数组，如果找到比设定最小元素还要小的值，交换两个元素的位置
*/

func SelectionSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
