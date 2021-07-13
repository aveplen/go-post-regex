package compiler

import (
	nfa "github.com/aveplen/go-post-regex/compiler/nfa"
	stack "github.com/aveplen/go-post-regex/compiler/nfa_stack"
	p "github.com/aveplen/go-post-regex/postfix"
)

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Compile(reg p.RegexP) *nfa.NFA {
	st := stack.New()
	if reg == "" {
		return nfa.FromEpsilon()
	}

	for _, token := range reg {
		switch token {
		case '*':
			{
				top, err := st.Pop()
				errPanic(err)
				st.Push(nfa.Closure(top))
			}
		case '|':
			{
				right, err := st.Pop()
				errPanic(err)
				left, err := st.Pop()
				errPanic(err)
				st.Push(nfa.Union(left, right))
			}
		case '.':
			{
				right, err := st.Pop()
				errPanic(err)
				left, err := st.Pop()
				errPanic(err)
				st.Push(nfa.Concat(left, right))
			}
		default:
			{
				st.Push(nfa.FromSymbol(token))
			}
		}
	}
	res, err := st.Pop()
	errPanic(err)
	if !st.Empty() {
		println(st.Length())
		panic("stack is not empty in the end")
	}
	return res
}
