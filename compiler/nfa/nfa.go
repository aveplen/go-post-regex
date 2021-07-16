package nfa

import state "github.com/aveplen/go-post-regex/state"

func New() *NFA {
	return &NFA{}
}

func FromEpsilon() *NFA {
	from := state.New(false)
	to := state.New(true)
	from.AddEpsilon(to)
	return &NFA{
		Start: from,
		End:   to,
	}
}

func FromSymbol(symbol rune) *NFA {
	from := state.New(false)
	to := state.New(true)
	from.AddTransition(to, symbol)
	return &NFA{
		Start: from,
		End:   to,
	}
}

func Concat(first *NFA, second *NFA) *NFA {
	first.End.AddEpsilon(second.Start)
	first.End.IsEnd = false
	return &NFA{
		Start: first.Start,
		End:   second.End,
	}
}

func Union(first *NFA, second *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	res.Start.AddEpsilon(first.Start)
	res.Start.AddEpsilon(second.Start)

	first.End.IsEnd = false
	first.End.AddEpsilon(res.End)

	second.End.IsEnd = false
	second.End.AddEpsilon(res.End)

	return res
}

func Closure(nfa *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	res.Start.AddEpsilon(res.End)
	res.Start.AddEpsilon(nfa.Start)

	nfa.End.AddEpsilon(res.End)
	nfa.End.AddEpsilon(nfa.Start)
	nfa.End.IsEnd = false

	return res
}

func Optional(nfa *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	res.Start.AddEpsilon(res.End)
	res.Start.AddEpsilon(nfa.Start)

	nfa.End.AddEpsilon(res.End)
	nfa.End.IsEnd = false

	return res
}

func Plus(nfa *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	// res.Start.AddEpsilon(res.End)
	res.Start.AddEpsilon(nfa.Start)

	nfa.End.AddEpsilon(res.End)
	nfa.End.AddEpsilon(nfa.Start)
	nfa.End.IsEnd = false

	return res
}
