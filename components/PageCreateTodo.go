package components

import (
	"todo/shared"

	"github.com/rivo/tview"
)

func CreateTodo(state *shared.AppState) tview.Primitive {
	var (
		value string = ""
		form *tview.Form
	)

	form = tview.NewForm().AddInputField("Todo", value, 64, nil, func(text string) {
		value = text
	}).AddButton("Add", func () {
		if value != "" {
			state.Todos = append(state.Todos, &shared.Todo{ Text: value, Done: false })
			RenderPage(shared.PageViewTodos, state)
		}
	}).AddButton("Quit", func() {
		RenderPage(shared.PageViewTodos, state)
	})

	return form
}