package go_leetcode

func removeElement(nums []int, val int) int {
	head := 0
	tail := len(nums) - 1

	for head <= tail {
		if nums[head] == val && nums[tail] != val {
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail--
		} else if nums[head] != val && nums[tail] != val {
			head++
		} else if nums[head] == val && nums[tail] == val {
			tail--
		} else {
			head++
		}
	}
	return head
}
