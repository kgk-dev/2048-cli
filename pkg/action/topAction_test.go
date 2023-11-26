package action

import (
	"2048/pkg/dto"
	"testing"
)

func TestTopAction(t *testing.T) {
	testcases := []struct {
		in, want dto.State
	}{
		{
			dto.State{
				{1, 0, 0, 0},
				{2, 4, 2, 4},
				{3, 2, 2, 2},
				{4, 2, 4, 4},
			},
			dto.State{
				{1, 4, 4, 4},
				{2, 4, 4, 2},
				{3, 0, 0, 4},
				{4, 0, 0, 0},
			},
		},
	}

	for _, tc := range testcases {
		result := topAction(tc.in)
		if result != tc.want {
			t.Errorf("Result: %v want: %v", result, tc.want)
		}
	}
}
