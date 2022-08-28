package state

type State uint

const (
	Query State = iota
	Input
	Output
	Save
)
