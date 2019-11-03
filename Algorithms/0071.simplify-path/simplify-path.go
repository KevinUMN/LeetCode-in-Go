package problem0071

import "strings"

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, v := range split {
		if v == ".." && len(stack) > 0 {
			stackLast := len(stack) - 1
			stack = stack[:stackLast]
		} else if v != "." && v != ".." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}
