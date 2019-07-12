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

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var defStyle tcell.Style

// func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
// 	for _, c := range str {
// 		var comb []rune
// 		w := runewidth.RuneWidth(c)
// 		if w == 0 {
// 			comb = []rune{c}
// 			c = ' '
// 			w = 1
// 		}
// 		s.SetContent(x, y, c, comb, style)
// 		x += w
// 	}
// }

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

// params @{index int} : index is current index
func drawList(s tcell.Screen, index int, list []string) {
	w, _ := s.Size()

	count := len(list)

	white := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	blue := tcell.StyleDefault.Background(tcell.ColorBlue)

	drawBox(s, 1, 20, w-2, 20+count+1, white, ' ')
	for i := 0; i < count; i++ {
		if i == index {
			emitStr(s, 3, 20+1+i, blue, list[i])
		} else {
			emitStr(s, 3, 20+1+i, white, list[i])
		}
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

	// f := NewFilesTodo()
	// f.FindData(".")

	defStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	box := NewBox()
	box.SetStyle(defStyle)
	box.SetTitle("Debug Box(디버그 창)")
	box.SetTitleColor(tcell.ColorYellow)

	// s.EnableMouse()
	s.Clear()

	keyfmt := "Keys: %s"
	white := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).Background(tcell.ColorBlue)

	lks := ""
	ecnt := 0
	index := 0
	for {
		box.SetRect(1, 1, w-2, 6)

		box.Draw(s)
		emitStr(s, 3, 2, white, "Press ESC twice or q to exit, C to clear.")
		emitStr(s, 3, 3, white, fmt.Sprintf("Box Size: (%d,%d) / %d x %d", box.x, box.y, box.width, box.height))
		x5, y5, x6, y6 := box.GetInnerRect()
		emitStr(s, 3, 4, defStyle, fmt.Sprintf("Inner Rect Size : (%d,%d) / %d x %d", x5, y5, x6, y6))
		emitStr(s, 3, 5, white, fmt.Sprintf(keyfmt, lks))
		// drawList(s, index, f.m.items)

		s.Show()
		ev := s.PollEvent()
		st := tcell.StyleDefault.Background(tcell.ColorBlack)
		w, h = s.Size()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			// Debug Info to lower right corner, "R" means Resize
			// s.SetContent(w-1, h-1, 'R', nil, st)
		case *tcell.EventKey:
			// ev.Rune() is key what you just pressed
			// s.SetContent(w-2, h-2, ev.Rune(), nil, st)
			// s.SetContent(w-1, h-1, 'K', nil, st)

			if ev.Key() == tcell.KeyEscape {
				ecnt++
				// escape key was pressed twice it'll quit(exit)
				if ecnt > 1 {
					s.Fini()
					os.Exit(0)
				}
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else {
				ecnt = 0
				if ev.Rune() == 'C' || ev.Rune() == 'c' {
					s.Clear()
				}
				if ev.Rune() == 'Q' || ev.Rune() == 'q' {
					s.Fini()
					os.Exit(0)
				}
				if ev.Rune() == 'J' || ev.Rune() == 'j' {
					index++
				}
				if ev.Rune() == 'K' || ev.Rune() == 'k' {
					index--
				}
			}
			lks = ev.Name()
		default:
			s.SetContent(w-1, h-1, 'X', nil, st)
		}

	}
}
