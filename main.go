package main

import (
	"2048/pkg/action"
	"2048/pkg/board"
	"2048/pkg/ui"
	"fmt"
	"log"
	"os"
	"path"
)

const PROMPT = "> "

func main() {
	var input string
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	uiFileName := path.Join(cwd, "pkg/ui/CLIUI.txt")
	render := ui.New(uiFileName)
	board := board.New()
	board.SetAction(action.New())
	board.SetRender(render)
	board.Start()
	for {
		fmt.Print(PROMPT)
		if _, err := fmt.Scan(&input); err != nil {
			log.Fatal(err)
		}
		switch input {
		case "e":
			fmt.Println("Exiting......")
			return
		case "t":
			board.MoveTop()
			board.Render()
		case "r":
			board.MoveRight()
			board.Render()
		case "b":
			board.MoveBottom()
			board.Render()
		case "l":
			board.MoveLeft()
			board.Render()
		}
	}
}
