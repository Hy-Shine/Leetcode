package go_leetcode

import "github.com/hy-shine/leetcode/go_leetcode/stack"

func IsValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	nodeStack := stack.NewStack()
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			nodeStack.Push(v)
			continue
		}
		if v == ']' && nodeStack.Value() == '[' {
			nodeStack.Pop()
		}
		if v == ')' && nodeStack.Value() == '(' {
			nodeStack.Pop()
		}
		if v == '}' && nodeStack.Value() == '{' {
			nodeStack.Pop()
		}
	}

	return nodeStack.IsEmpty()
}
