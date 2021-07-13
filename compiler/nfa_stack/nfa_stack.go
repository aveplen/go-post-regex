package nfa_stack

import (
	"errors"

	nfa "github.com/aveplen/go-post-regex/compiler/nfa"
)

func (st *Stack) Empty() bool {
	return st.len <= 0
}

func (st *Stack) Push(x *nfa.NFA) {
	st.cont = append(st.cont, x)
	st.len++
}

func (st *Stack) Pop() (*nfa.NFA, error) {
	if st.Empty() {
		return nfa.New(), errors.New("stack is empty")
	}
	st.len--
	res := st.cont[st.len]
	st.cont = st.cont[:st.len]
	return res, nil
}

func New() Stack {
	return Stack{
		len:  0,
		cont: make([]*nfa.NFA, 0),
	}
}

func (st *Stack) Length() int {
	return st.len
}
