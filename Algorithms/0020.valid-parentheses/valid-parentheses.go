package problem0020

func isValid(s string) bool {
	size := len(s)

	stack := make([]int32, size)
	top := -1

	for _, v := range s {
		switch v {
		case '(':
			top++
			stack[top] = ')'
		case '{':
			top++
			stack[top] = '}'
		case '[':
			top++
			stack[top] = ']'
		case ')', '}', ']':
			if top == -1 || v != stack[top] {
				return false
			} else {
				top--
			}
		}
	}

	return top == -1
}
