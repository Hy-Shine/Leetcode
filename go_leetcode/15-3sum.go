package go_leetcode

import (
	"sort"
)

func twoSumSlice(nums []int, target int) [][]int {
	sort.Ints(nums)

	var result [][]int
	low, hight := 0, len(nums)-1
	m := make(map[int]struct{})
	for low <= hight {
		sum := nums[low] + nums[hight]
		if sum == target {
			if _, ok := m[nums[low]]; !ok {
				result = append(result, []int{nums[low], nums[hight]})
				m[low] = struct{}{}
			}
			low++
			hight--
		} else if sum > target {
			hight--
		} else {
			low++
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	m := make(map[int]struct{})
	var result [][]int
	for i := 0; i < len(nums)-1; i++ {
		l := twoSumSlice(nums[i+1:], 0-nums[i])
		if len(l) == 0 {
			continue
		}
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = struct{}{}
		for k := range l {
			result = append(result, []int{nums[i], l[k][0], l[k][1]})
		}

	}
	return result
}
