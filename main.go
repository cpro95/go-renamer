//+build ignore
//+build ignore

// Copyright 2015 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// mouse displays a text box and tests mouse interaction.  As you click
// and drag, boxes are displayed on screen.  Other events are reported in
// the box.  Press ESC twice to exit the program.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var defStyle tcell.Style

// params @{r rune} : fill r as white space
func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, r rune) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	if y1 != y2 && x1 != x2 {
		// Only add corners if we need to
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		for col := x1 + 1; col < x2; col++ {
			s.SetContent(col, row, r, nil, style)
		}
	}
}

func loadData(f *filesTodo, list, list2 *List) {
	for _, item := range f.m.items {
		list.AddItem(item, nil)
	}
	for _, item := range f.s.items {
		list2.AddItem(item, nil)
	}
}

func handleDown(screen tcell.Screen, list, list2 *List) {
	if list.HasFocus() {
		if list.selected {
			if list.currentItem < list.GetItemCount()-1 {
				list.items[list.currentItem], list.items[list.currentItem+1] = list.items[list.currentItem+1], list.items[list.currentItem]
				handleSelect(screen, list, list2)
				handleDown(screen, list, list2)
			}
		} else {
			list.SetCurrentItem(list.currentItem + 1)
		}
	} else if list2.HasFocus() {
		if list2.selected {
			if list2.currentItem < list2.GetItemCount()-1 {
				list2.items[list2.currentItem], list2.items[list2.currentItem+1] = list2.items[list2.currentItem+1], list2.items[list2.currentItem]
				handleSelect(screen, list, list2)
				handleDown(screen, list, list2)
			}
		} else {
			list2.SetCurrentItem(list2.currentItem + 1)
		}
	}
}

func handleUp(screen tcell.Screen, list, list2 *List) {
	if list.HasFocus() {
		if list.selected {
			if list.currentItem > 0 {
				list.items[list.currentItem-1], list.items[list.currentItem] = list.items[list.currentItem], list.items[list.currentItem-1]
				handleSelect(screen, list, list2)
				handleUp(screen, list, list2)
			}
		} else {
			list.SetCurrentItem(list.currentItem - 1)
		}
	} else if list2.HasFocus() {
		if list2.selected {
			if list2.currentItem > 0 {
				list2.items[list2.currentItem-1], list2.items[list2.currentItem] = list2.items[list2.currentItem], list2.items[list2.currentItem-1]
				handleSelect(screen, list, list2)
				handleUp(screen, list, list2)
			}
		} else {
			list2.SetCurrentItem(list2.currentItem - 1)
		}
	}
}

func handleSelect(screen tcell.Screen, list, list2 *List) {
	if list.HasFocus() {
		x, y, _, _ := list.GetInnerRect()
		_, _, style, _ := screen.GetContent(x, y+list.currentItem)
		_, bg, _ := style.Decompose()
		if bg == tcell.ColorWhite {
			list.SetSelectedBackgroundColor(tcell.ColorGrey)
		} else {
			list.SetSelectedBackgroundColor(tcell.ColorWhite)
		}

		// convert selected bool
		list.selected = !list.selected
	} else if list2.HasFocus() {
		x, y, _, _ := list2.GetInnerRect()
		_, _, style, _ := screen.GetContent(x, y+list2.currentItem)
		_, bg, _ := style.Decompose()
		if bg == tcell.ColorWhite {
			list2.SetSelectedBackgroundColor(tcell.ColorGrey)
		} else {
			list2.SetSelectedBackgroundColor(tcell.ColorWhite)
		}

		// convert selected bool
		list2.selected = !list2.selected
	}
}

// This program just shows simple mouse and keyboard events.  Press ESC twice to
// exit.
func main() {
	encoding.Register()

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	// get the w, h of Screen of tcell
	w, h := s.Size()

	defStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	box := NewBox()
	box.SetStyle(defStyle)
	box.SetTitle("Debug Box(디버그 창)")
	box.SetTitleColor(tcell.ColorYellow)
	box.SetRect(1, 1, w-2, 6)

	f := NewFilesTodo()
	f.FindData(".")

	list := NewList()
	list.Box.SetRect(1, 8, w-2, f.m.GetItemCount()+2)

	list.Box.SetBorderPadding(0, 0, 1, 1)

	list2 := NewList()
	list2.Box.SetRect(1, 8+1+f.m.GetItemCount()+2, w-2, f.s.GetItemCount()+2)

	list2.Box.SetBorderPadding(0, 0, 1, 1)
	list2.SetFocus(false)

	loadData(f, list, list2)

	// s.EnableMouse()
	s.Clear()

	keyfmt := "Keys: %s"
	white := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).Background(tcell.ColorBlue)

	lks := ""
	for {

		box.Draw(s)
		emitStr(s, 3, 2, white, "Press Q to exit, R to Reload, JK to Up & Down, D to Delete Item, Y to Rename This")
		emitStr(s, 3, 3, white, "Tab to Switching, Space to Selecting & Greying Item and then JK to Up & Down the selected item")
		x5, y5, x6, y6 := box.GetInnerRect()
		emitStr(s, 3, 4, defStyle, fmt.Sprintf("Box Size: (%d,%d) / %d x %d - Inner Rect Size : (%d,%d) / %d x %d", box.x, box.y, box.width, box.height, x5, y5, x6, y6))
		emitStr(s, 3, 5, white, fmt.Sprintf(keyfmt, lks))

		list.Box.SetTitle(strconv.Itoa(list.currentItem) + " " + strconv.FormatBool(list.selected))
		list.Draw(s)
		list2.Box.SetTitle(strconv.Itoa(list2.currentItem) + " " + strconv.FormatBool(list2.selected))
		list2.Draw(s)

		s.Show()
		ev := s.PollEvent()
		w, h = s.Size()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Sync()
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Key() == tcell.KeyTab {
				if list.HasFocus() {
					list.SetFocus(false)
					list2.SetFocus(true)
				} else if list2.HasFocus() {
					list.SetFocus(true)
					list2.SetFocus(false)
				}
			} else if ev.Key() == tcell.KeyDown {
				handleDown(s, list, list2)
			} else if ev.Key() == tcell.KeyUp {
				handleUp(s, list, list2)
			} else {
				if ev.Rune() == 'R' || ev.Rune() == 'r' {
					s.Clear()
					list.Clear()
					list2.Clear()
					loadData(f, list, list2)
				}
				if ev.Rune() == 'Q' || ev.Rune() == 'q' {
					s.Fini()
					os.Exit(0)
				}
				if ev.Rune() == 'J' || ev.Rune() == 'j' {
					handleDown(s, list, list2)
				}
				if ev.Rune() == 'K' || ev.Rune() == 'k' {
					handleUp(s, list, list2)
				}
				if ev.Rune() == 'D' || ev.Rune() == 'd' {
					if list.HasFocus() {
						list.RemoveItem(list.currentItem)
					} else if list2.HasFocus() {
						list2.RemoveItem(list2.currentItem)
					}
				}
				if ev.Rune() == ' ' {
					handleSelect(s, list, list2)
				}

			}
			lks = ev.Name()
		default:
			s.SetContent(w-1, h-1, 'X', nil, defStyle)
		}

	}
}
