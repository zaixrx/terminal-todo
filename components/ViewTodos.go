package components

import (
	"todo/shared"

	"github.com/rivo/tview"
)

func ViewTodos(state *shared.AppState) tview.Primitive {
	var (
		list *tview.List = tview.NewList()
		getCheckmark = func (checked bool) string {
			if checked {
				return "Done: ✓"
			}
	
			return "Done: ✗"
		}
	)

	for i, todo := range state.Todos {
		i, todo := i, todo 
		list.AddItem(todo.Text, getCheckmark(todo.Done), rune('0' + i), func () {
			todo.Done = !todo.Done
			go state.App.QueueUpdateDraw(func () {
				list.SetItemText(i, todo.Text, getCheckmark(todo.Done))
			})
		})
	}

	list.
		AddItem("Add", "Add a todo", '+', func() {
			RenderPage(shared.PageCreateTodo, state)
		}).
		AddItem("Save", "Save the list", 's', func() {
			RenderPage(shared.PageSaveTodos, state)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			state.App.Stop()
		})

	return list
}