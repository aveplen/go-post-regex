package nfa

import state "github.com/aveplen/go-post-regex/state"

type NFA struct {
	Start *state.State
	End   *state.State
}
