package components

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo/shared"

	"github.com/rivo/tview"
)

func SaveTodos(state *shared.AppState) tview.Primitive {
	var (
		dir = ""
		name = "todo"
		form *tview.Form
	)

	form = tview.NewForm().AddInputField("Name", name, 64, nil, func(text string) {
		name = text
	}).AddInputField("Directory", dir, 64, nil, func(text string) {
		dir = text
	}).AddButton("Add", func () {
		if name != "" {
			data, err := json.MarshalIndent(state.Todos, "", "  ")
			shared.Check(err)

			fullPath := filepath.Join(dir, name + ".json")
			err = os.WriteFile(fullPath, data, 0644) // Octal Linux Priorities from OS 1CP S1(yamat yah!)
			if err != nil {
				RenderErrorLayout(state, err.Error())
				return;
			} 

			RenderFeedbackLayout(state, fmt.Sprintf("Saved to %s", name), func () {
				RenderPage(shared.PageViewTodos, state)
			})
		}
	}).AddButton("Quit", func() {
		RenderPage(shared.PageViewTodos, state)
	})

	return form
}
