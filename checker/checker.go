package checker

import (
	mult "github.com/aveplen/go-post-regex/checker/multistate"
	nfa "github.com/aveplen/go-post-regex/compiler/nfa"
)

func Match(nfa *nfa.NFA, word string) bool {
	currentStates := mult.New()
	visited := mult.New()
	mult.AddNextStates(nfa.Start, &currentStates, &visited)

	for _, symbol := range word {
		nextStates := mult.New()
		visited = mult.New()
		for _, state := range currentStates {
			next, ok := state.Transition[symbol]
			if ok {
				mult.AddNextStates(next, &nextStates, &visited)
			}
		}
		currentStates = nextStates
	}

	return currentStates.Find(nfa.End)
}
