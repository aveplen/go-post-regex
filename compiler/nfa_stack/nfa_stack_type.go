package nfa_stack

import nfa "github.com/aveplen/go-post-regex/compiler/nfa"

type Stack struct {
	cont []*nfa.NFA
	len  int
}
