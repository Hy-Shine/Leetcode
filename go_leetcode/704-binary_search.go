package go_leetcode

func binarySearch(list []int, target int) int {
	head := 0
	tail := len(list) - 1

	for head <= tail {
		mid := (head + tail) / 2
		if list[mid] == target {
			return mid
		} else if list[mid] < target {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return -1
}
