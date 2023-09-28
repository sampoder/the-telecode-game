package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

type Room struct {
	code         string
	createdAt    string
	public	     bool
}

var rooms = make([]Room, 0)

var pages = tview.NewPages()
var roomText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm()
var roomsList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText(" (n) to create a new room \n (j) to join a room \n (q) to quit")

func main() {
	roomsList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		setRoomText(&rooms[index])
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(roomsList, 0, 1, true).
			AddItem(roomText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 110 {
			form.Clear(true)
			addRoomForm()
			pages.SwitchToPage("Create Room")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Create Room", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func addRoomList() {
	roomsList.Clear()
	for index, room := range rooms {
		roomsList.AddItem(room.code, " ", rune(49+index), nil)
	}
}

func addRoomForm() *tview.Form {

	room := Room{}

	form.AddInputField("Code", "", 20, nil, func(code string) {
		room.code = code
	})

	form.AddButton("Save", func() {
		room.createdAt = time.Now().Format(time.RFC3339)
		rooms = append(rooms, room)
		addRoomList()
		pages.SwitchToPage("Menu")
	})

	return form
}

func setRoomText(room *Room) {
	roomText.Clear()
	text := room.code + "\n" + room.createdAt
	roomText.SetText(text)
}