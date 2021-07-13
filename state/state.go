package state

func New(isEnd bool) *State {
	return &State{
		IsEnd:      isEnd,
		Transition: make(map[rune]*State),
		Epsilon:    make([]*State, 0),
	}
}

func (from *State) AddEpsilon(to *State) {
	from.Epsilon = append(from.Epsilon, to)
}

func (from *State) AddTransition(to *State, symbol rune) {
	from.Transition[symbol] = to
}
