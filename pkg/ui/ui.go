package ui

import (
	"2048/pkg/dto"
	"fmt"
	"io"
	"text/template"
)

func SetUI(uiFilename string) dto.Render {
	t, err := template.ParseFiles(uiFilename)
	if err != nil {
		fmt.Println("err in template", err)
		return nil
	}
	return func(w io.Writer, s dto.State) error {
		return t.Execute(w, s)
	}
}

func New(uiFileName string) dto.Render {
	return SetUI(uiFileName)
}
