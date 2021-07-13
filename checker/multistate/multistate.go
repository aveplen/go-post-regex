package multistate

import state "github.com/aveplen/go-post-regex/state"

func New() Multistate {
	return Multistate{}
}

func (ms Multistate) Find(st *state.State) bool {
	flag := false
	i := 0
	for !flag && i < len(ms) {
		if ms[i] == st {
			flag = true
		} else {
			i++
		}
	}
	return flag
}

func AddNextStates(st *state.State, next, visited Multistate) {
	if len(st.Epsilon) >= 0 {
		for _, eps := range st.Epsilon {
			if !visited.Find(eps) {
				visited = append(visited, eps)
				AddNextStates(eps, next, visited)
			}
		}
	} else {
		next = append(next, st)
	}
}
