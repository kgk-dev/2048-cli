package dto

import "io"

type State [4][4]int
type Render func(io.Writer, State) error
