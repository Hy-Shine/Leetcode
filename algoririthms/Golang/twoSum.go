package main

import "fmt"

func main() {
	p := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(p, target))
}

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
