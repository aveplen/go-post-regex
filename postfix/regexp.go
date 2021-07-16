package postfix

import (
	"strings"

	stack "github.com/aveplen/go-post-regex/postfix/stack"
)

var presedence = map[rune]int{
	'(': 0,
	'|': 1,
	'.': 2,
	'?': 3,
	'*': 3,
	'+': 3,
}

func (r Regex) ToPostfix() (res RegexP) {
	st := stack.New()
	for _, token := range r {
		if strings.Contains(".*?+|", string(token)) {
			for {
				if top, err := st.Peek(); err == nil {
					preTop := presedence[top]
					if preTok, ok := presedence[token]; ok && preTop >= preTok {
						res = res + RegexP(top)
						st.Pop()
						continue
					}
					break
				}
				break
			}
			st.Push(token)
		} else {
			if token == '(' || token == ')' {
				if token == '(' {
					st.Push(token)
				} else {
					for top, err := st.Pop(); err == nil && top != '('; top, err = st.Pop() {
						res = res + RegexP(top)
					}
				}
			} else {
				res = res + RegexP(token)
			}
		}
	}
	for top, err := st.Pop(); err == nil; top, err = st.Pop() {
		res = res + RegexP(top)
	}
	return res
}
