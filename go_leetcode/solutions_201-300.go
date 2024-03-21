package go_leetcode

import (
	"fmt"
	"strconv"
)

// problems 200-300 solutions

// No.268
func missingNumber(nums []int) int {
	n := len(nums)
	var m = make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i <= n; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

// No.268
func missingNumber2(nums []int) int {
	n := len(nums)
	expect := n * (n + 1) / 2

	got := 0
	for i := 0; i < n; i++ {
		got += nums[i]
	}
	return expect - got
}

// No.283
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			fast++
			continue
		}
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		slow++
		fast++
	}
}

// No.228
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	f := func(head, tail int, l []int) string {
		if head == tail {
			return strconv.Itoa(l[head])
		}
		return fmt.Sprintf("%d->%d", l[head], l[tail])
	}

	var head int
	var rangeList []string
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[head] == i+1-head && i != len(nums)-2 {
			continue
		}

		if i != len(nums)-2 {
			rangeList = append(rangeList, f(head, i, nums))
			head = i + 1
		} else if nums[len(nums)-1]-nums[head] != len(nums)-1-head {
			rangeList = append(rangeList, f(head, len(nums)-2, nums))
			rangeList = append(rangeList, f(len(nums)-1, len(nums)-1, nums))
		} else {
			rangeList = append(rangeList, f(head, len(nums)-1, nums))
		}
	}
	return rangeList
}
