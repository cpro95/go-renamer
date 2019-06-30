package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {

	f := NewFilesTodo()
	f.FindData(".")

	app := tview.NewApplication()
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	textView := tview.NewTextView()
	fmt.Fprintf(textView,
		`Movies Renamer with Subtitles files
	Tab to Navigate to another panel
	Y to Rename it
	Q,CtrlQ,CtrlC to Quit
		`)

	moviesList := tview.NewList()
	for i := 0; i < f.m.GetItemCount(); i++ {
		moviesList.AddItem(f.m.items[i], "", rune(strconv.Itoa(i + 1)[0]), nil)
	}
	moviesList.SetBorder(true).SetBorderColor(tcell.ColorYellow).SetBorderPadding(1, 1, 1, 1)

	subtitlesList := tview.NewList()
	for i := 0; i < f.s.GetItemCount(); i++ {
		subtitlesList.AddItem(f.s.items[i], "", rune(strconv.Itoa(i + 1)[0]), nil)
	}

	subtitlesList.SetBorder(true).SetBorderColor(tcell.ColorBlue).SetBorderPadding(1, 1, 1, 1)
	flex.AddItem(textView, 0, 1, false)
	flex.AddItem(moviesList, 0, 4, false)
	flex.AddItem(subtitlesList, 0, 4, false)

	// app.SetAfterDrawFunc(func(screen tcell.Screen) {
	// 	count := f.m.GetItemCount()

	// 	tview.Print(screen, string(f.m.items[0]), 2, 1, 10, tview.AlignLeft, tcell.ColorYellow)
	// 	tview.Print(screen, string(f.m.items[1]), 2, 2, 10, tview.AlignLeft, tcell.ColorYellow)
	// 	tview.Print(screen, strconv.Itoa(count), 2, 30, 10, tview.AlignLeft, tcell.ColorYellow)
	// })

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlQ {
			app.Stop()
		}
		if event.Key() == tcell.KeyTab {
			nowPrimitive := app.GetFocus()
			if nowPrimitive == moviesList {
				app.SetFocus(subtitlesList)
			} else if nowPrimitive == subtitlesList {
				app.SetFocus(moviesList)
			}
		}
		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(moviesList).Run(); err != nil {
		panic(err)
	}

}
