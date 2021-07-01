package insertionSort

// https://www.runoob.com/w3cnote/insertion-sort.html

func InsertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i] // 记录替换的起始位置
		for preIndex >= 0 && arr[preIndex] > current {
			// 如果前面的元素的值大于后面的元素，用前面的元素覆盖后面的元素
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1 // 向前遍历
		}
		arr[preIndex+1] = current // 把之前记录的数据放到遍历结束以后最前面的位置
	}

	return arr
}
