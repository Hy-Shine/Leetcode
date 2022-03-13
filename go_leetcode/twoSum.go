package go_leetcode

func twoSum(nums []int, target int) []int {
	temp := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				temp = append(temp, i, j)
				break
			}
		}
	}
	return temp
}
