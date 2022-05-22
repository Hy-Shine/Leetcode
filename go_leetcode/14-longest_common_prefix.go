package go_leetcode

import "github.com/hy-shine/leetcode/go_leetcode/utils"

func longestCommonPrefix(strList []string) string {
	if len(strList) == 0 {
		return ""
	}
	common := strList[0]
	for i := 0; i < len(strList); i++ {
		common = longestCommonPrefixTwoString(common, strList[i])
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
