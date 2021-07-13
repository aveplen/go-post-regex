package stack

import (
	"errors"
	"fmt"
)

func (st *Stack) Push(x rune) {
	st.cont = append(st.cont, x)
	st.len++
}

func (st *Stack) Pop() (rune, error) {
	if st.Empty() {
		return 0, errors.New("stack is empty")
	}
	st.len--
	res := st.cont[st.len]
	st.cont = st.cont[:st.len]
	return res, nil
}

func (st *Stack) Peek() (rune, error) {
	if st.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return st.cont[st.len-1], nil
}

func (st *Stack) Empty() bool {
	return st.len <= 0
}

func New() Stack {
	return Stack{len: 0}
}
