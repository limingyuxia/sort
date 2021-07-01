package bubbleSort

// https://www.runoob.com/w3cnote/bubble-sort.html

/*
	1. 遍历所有元素
	2. 比较相邻的元素，如果后面的比前面的大，就交换
*/

func BubbleSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
