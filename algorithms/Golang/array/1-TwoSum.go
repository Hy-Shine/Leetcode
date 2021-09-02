package array

func twoSum(nums []int, target int) []int {
	listMap := make(map[int]int)
	list := make([]int, 2)
	for i, v := range nums {
		if k, ok := listMap[target-v]; ok {
			list = []int{i, k}
			break
		} else {
			listMap[v] = i
		}
	}
	return list
}
