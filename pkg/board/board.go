package board

import (
	"2048/pkg/dto"
	"math/rand"
	"os"
)

const StartValue = 2

type Move interface {
	Top(dto.State) dto.State
	Right(dto.State) dto.State
	Bottom(dto.State) dto.State
	Left(dto.State) dto.State
}

type Board struct {
	move   Move
	state  [4][4]int
	render dto.Render
}

func (b *Board) MoveTop() {
	b.state = b.move.Top(b.state)
}

func (b *Board) MoveRight() {
	b.state = b.move.Right(b.state)
}

func (b *Board) MoveBottom() {
	b.state = b.move.Bottom(b.state)
}

func (b *Board) MoveLeft() {
	b.state = b.move.Left(b.state)
}

func (b *Board) SetAction(moveAction Move) {
	b.move = moveAction
}

func (b *Board) newCell() {
	for {
		row, col := rand.Intn(4), rand.Intn(4)
		if b.state[row][col] == 0 {
			b.state[row][col] = StartValue
			break
		}
	}
}

func (b *Board) Start() {
	b.newCell()
	b.Render()
}

func (b *Board) SetRender(render dto.Render) {
	b.render = render
}

func (b *Board) Render() {
	b.newCell()
	b.render(os.Stdout, b.state)
}

func New() *Board {
	return &Board{}
}
