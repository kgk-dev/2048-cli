package action

import "2048/pkg/dto"

const EMPTYCELL = 0
const CELL = 1

type Action struct{}

func (a Action) Top(input dto.State) dto.State {
	return topAction(input)
}

func (a Action) Right(input dto.State) dto.State {
	return rightAction(input)
}

func (a Action) Bottom(input dto.State) dto.State {
	return bottomAction(input)
}

func (a Action) Left(input dto.State) dto.State {
	return leftAction(input)
}

func New() Action {
	return Action{}
}
