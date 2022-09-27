package state

type State uint

const (
	Query State = iota
	Running
	Input
	Output
	Save
)
