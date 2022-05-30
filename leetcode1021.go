package main

import "fmt"

func main() {
	s := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(s))

}

func removeOuterParentheses(s string) string {
	b := make([]byte, 0, len(s))
	pre, count := 0, 0
	for ind, value := range s {
		if value == '(' {
			count += 1
		} else {
			count -= 1
		}
		if count == 0 {
			b = append(b, s[pre+1:ind]...)
			pre = ind + 1
		}
	}
	return string(b)

}
