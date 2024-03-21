package go_leetcode

func findKthLargest(list []int, k int) int {
	for i := range list {
		minEle := i
		for j := i; j < len(list); j++ {
			if list[j] > list[minEle] {
				minEle = j
			}
		}
		list[i], list[minEle] = list[minEle], list[i]
		if i == k-1 {
			return list[i]
		}
	}
	return -1
}
