package main

import (
	"fmt"

	stack "github.com/aveplen/go-weak-stack"
)

type regex string

const r regex = "https?+:(//)*twitch\\.tv/(verni_shavermy)"

func explicitConcat(r regex) (res regex) {
	for i, token := range r {
		res = res + regex(token)
		if token == '(' || token == '|' {
			continue
		}
		if i+1 < len(r) {
			lookahead := r[i+1]
			if lookahead == '*' || lookahead == '?' || lookahead == '+' || lookahead == '|' || lookahead == ')' {
				continue
			}
			res = res + "."
		}
	}
	return res
}

type postRegex regex

func (r regex) toPostfix() postRegex {
	// not implemented
	return postRegex("")
}

func main() {
	fmt.Println("Bfr:", r)
	fmt.Println("Aft:", explicitConcat(r))

	st := stack.New()
	for i := 0; i < 10; i++ {
		st.Push(i)
		fmt.Println(st)
	}

	for j := 0; j < 8; j++ {
		x, _ := st.Pop()
		fmt.Println(st, "   ", x)
	}
}
