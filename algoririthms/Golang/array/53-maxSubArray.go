package array

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	sumList := []int{}
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			sumList = append(sumList, getListSum(nums[i:j]))
		}
	}
	return maxNums(sumList)
}

func getListSum(nums []int) int {
	sum := nums[0]
	for _, v := range nums {
		sum += v
	}
	return sum
}

func maxNums(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v >= max {
			max = v
		}
	}
	return max
}
