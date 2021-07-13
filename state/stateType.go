package state

type State struct {
	IsEnd      bool
	Transition map[rune]*State
	Epsilon    []*State
}
