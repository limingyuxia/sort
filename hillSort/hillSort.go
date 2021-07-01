package hillSort

// https://www.runoob.com/w3cnote/shell-sort.html

func HillSort(arr []int) []int {
	length := len(arr)
	gap := 1 // 增量

	for gap < gap/3 {
		gap = gap*3 + 1
	}

	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]                // 记录右边起始位置的元素
			j := i - gap                  // j 向前跳过一个增量
			for j >= 0 && arr[j] > temp { // 前面的元素比后面的大
				arr[j+gap] = arr[j] // 前面的元素覆盖后面的元素
				j -= gap
			}
			arr[j+gap] = temp // 起始位置的元素放到遍历结束的开头
		}

		gap = gap / 3
	}

	return arr
}
