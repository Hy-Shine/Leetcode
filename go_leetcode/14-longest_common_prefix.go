package go_leetcode

import "github.com/hy-shine/leetcode/go_leetcode/utils"

func longestCommonPrefix(strs []string) string {
	common := strs[0]
	for i := 0; i < len(strs); i++ {
		common = longestCommonPrefixTwoString(common, strs[i])
	}
	return common
}

func longestCommonPrefixTwoString(a, b string) string {
	max := utils.MinStringLength(a, b)
	i := 0
	for i < max && a[i] == b[i] {
		i++
	}
	return a[:i]
}
