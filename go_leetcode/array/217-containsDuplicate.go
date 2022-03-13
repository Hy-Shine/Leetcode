package array

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			return true
		} else {
			numsMap[v] = 1
		}
	}
	return false
}
