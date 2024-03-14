package sorts

// 选择排序
func selectionSort(s []int) []int {
	min := 0
	for i := 0; i < len(s)-1; i++ {
		min = findSmallest(s[i+1:]) + i + 1
		if s[i] > s[min] {
			s[i], s[min] = s[min], s[i]
		}
	}
	return s
}

// get the min value's index of a array
func findSmallest(p []int) int {
	min := p[0]
	n := 0
	for i := 0; i < len(p); i++ {
		if p[i] < min {
			min = p[i]
			n = i
		}
	}
	return n
}

// 快速排序
func quickSort(list []int) {
	_quickSort(list, 0, len(list)-1)
}

func _quickSort(list []int, start, end int) {
	if start >= end {
		return
	}

	index := partition(list, start, end)
	_quickSort(list, start, index-1)
	_quickSort(list, index+1, end)
}

func partition(list []int, start, end int) int {
	pivot := start
	// index 索引始终作为比基准大的第一个元素
	index := start + 1
	for i := index; i <= end; i++ {
		if list[i] < list[pivot] {
			list[i], list[index] = list[index], list[i]
			index++
		}
	}
	list[pivot], list[index-1] = list[index-1], list[pivot]
	// 返回基准索引
	return index - 1
}
