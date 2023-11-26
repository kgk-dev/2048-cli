package action

import (
	"2048/pkg/dto"
)

func bottomAction(input dto.State) dto.State {
	var output dto.State
	for colId := 0; colId < len(input); colId++ {
		for rowId, outputRowId := len(input)-1, 2; rowId >= 0; rowId-- {
			cell := input[rowId][colId]
			if cell != EMPTYCELL {
				if output[outputRowId+1][colId] == EMPTYCELL {
					output[outputRowId+1][colId] = cell
				} else if output[outputRowId+1][colId] == cell {
					output[outputRowId+1][colId] += cell
					outputRowId--
				} else {
					output[outputRowId][colId] = cell
					outputRowId--
				}
			}
		}
	}
	return output
}
