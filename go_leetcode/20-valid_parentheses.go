package go_leetcode

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
			continue
		}
		index := len(stack)
		if index == 0 {
			return false
		}
		switch {
		case stack[index-1] != '[' && v == ']':
			return false
		case stack[index-1] != '(' && v == ')':
			return false
		case stack[index-1] != '{' && v == '}':
			return false
		case stack[index-1] == '[' && v == ']':
			stack = push(stack)
		case stack[index-1] == '(' && v == ')':
			stack = push(stack)
		case stack[index-1] == '{' && v == '}':
			stack = push(stack)
		}
	}

	return len(stack) == 0
}

func push(stack []rune) []rune {
	l := len(stack)
	if l == 1 {
		return stack[0:0]
	}
	return stack[0 : l-1]
}
