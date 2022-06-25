package main

type node struct {
	c   byte
	idx int
}

func removeOuterParentheses(s string) string {
	stack := []*node{}
	del := []int{}
	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, &node{c: byte(ch), idx: i})
		} else {
			if len(stack) == 1 {
				del = append(del, stack[0].idx)
				del = append(del, i)
			}
			stack = stack[:len(stack)-1]
		}
	}
	j := 0
	ans := ""
	for i := range s {
		if del[j] == i {
			j++
			continue
		}
		ans += string(s[i])
	}
	return ans
}

// func main() {
// 	//"(()())(())"
// 	//"(()())(())(()(()))"
// 	//"()()"
// 	fmt.Println(removeOuterParentheses("(()())(())"))
// 	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
// 	fmt.Println(removeOuterParentheses("()()"))
// }
