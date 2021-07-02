package search

func findRepeatNumber(nums []int) int {
	set := map[int]int{nums[0]: 0}
	for i := 1; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return nums[i]
		}
		set[nums[i]] = i
	}
	return -1
}
