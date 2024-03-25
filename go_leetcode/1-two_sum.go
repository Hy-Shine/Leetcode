package go_leetcode

func twoSum(nums []int, target int) []int {
	list := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				list = append(list, i, j)
				break
			}
		}
	}
	return list
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := m[target-nums[i]]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}
