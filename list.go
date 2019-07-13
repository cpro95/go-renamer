package main

import (
	"strings"

	"github.com/gdamore/tcell"
)

// listItem represents one item in a List.
type listItem struct {
	MainText string // The main text of the list item.
}

// List displays rows of items, each of which can be selected.
//
// See https://github.com/rivo/tview/wiki/List for an example.
type List struct {
	*Box

	// The items of the list.
	items []*listItem

	// The index of the currently selected item.
	currentItem int

	// Selected if currentItem
	selected bool

	// The item main text color.
	mainTextColor tcell.Color

	// The text color for selected items.
	selectedTextColor tcell.Color

	// The background color for selected items.
	selectedBackgroundColor tcell.Color

	// If true, the entire row is highlighted when selected.
	highlightFullLine bool

	// The number of list items skipped at the top before the first item is drawn.
	offset int

	// Whether or not this list has focus.
	hasFocus bool
}

// NewList returns a new form.
func NewList() *List {
	return &List{
		Box:                     NewBox(),
		highlightFullLine:       true,
		selected:                false,
		mainTextColor:           tcell.ColorWhite,
		selectedTextColor:       tcell.ColorBlack,
		selectedBackgroundColor: tcell.ColorWhite,
		hasFocus:                true,
	}
}

// SetCurrentItem sets the currently selected item by its index, starting at 0
// for the first item. If a negative index is provided, items are referred to
// from the back (-1 = last item, -2 = second-to-last item, and so on). Out of
// range indices are clamped to the beginning/end.
//
// Calling this function triggers a "changed" event if the selection changes.
func (l *List) SetCurrentItem(index int) *List {
	if index < 0 {
		index = len(l.items) + index
	}

	if index >= len(l.items) {
		index = 0
	}
	if index < 0 {
		index = 0
	}

	l.currentItem = index

	return l
}

// GetCurrentItem returns the index of the currently selected list item,
// starting at 0 for the first item.
func (l *List) GetCurrentItem() int {
	return l.currentItem
}

// RemoveItem removes the item with the given index (starting at 0) from the
// list. If a negative index is provided, items are referred to from the back
// (-1 = last item, -2 = second-to-last item, and so on). Out of range indices
// are clamped to the beginning/end, i.e. unless the list is empty, an item is
// always removed.
//
// The currently selected item is shifted accordingly. If it is the one that is
// removed, a "changed" event is fired.
func (l *List) RemoveItem(index int) *List {
	if len(l.items) == 0 {
		return l
	}

	// // Adjust index.
	// if index < 0 {
	// 	index = len(l.items) + index
	// }
	// if index >= len(l.items) {
	// 	index = len(l.items) - 1
	// }
	// if index < 0 {
	// 	index = 0
	// }

	// Remove item.
	l.items = append(l.items[:index], l.items[index+1:]...)

	// If there is nothing left, we're done.
	if len(l.items) == 0 {
		return l
	}

	// Shift current item.
	// previousCurrentItem := l.currentItem
	if l.currentItem >= index {
		if l.currentItem != 0 {
			l.currentItem--
		}
	}

	return l
}

// SetMainTextColor sets the color of the items' main text.
func (l *List) SetMainTextColor(color tcell.Color) *List {
	l.mainTextColor = color
	return l
}

// SetSelectedTextColor sets the text color of selected items.
func (l *List) SetSelectedTextColor(color tcell.Color) *List {
	l.selectedTextColor = color
	return l
}

// SetSelectedBackgroundColor sets the background color of selected items.
func (l *List) SetSelectedBackgroundColor(color tcell.Color) *List {
	l.selectedBackgroundColor = color
	return l
}

// SetHighlightFullLine sets a flag which determines whether the colored
// background of selected items spans the entire width of the view. If set to
// true, the highlight spans the entire view. If set to false, only the text of
// the selected item from beginning to end is highlighted.
func (l *List) SetHighlightFullLine(highlight bool) *List {
	l.highlightFullLine = highlight
	return l
}

// AddItem calls InsertItem() with an index of -1.
func (l *List) AddItem(mainText string, selected func()) *List {
	l.InsertItem(-1, mainText, selected)
	return l
}

// InsertItem adds a new item to the list at the specified index. An index of 0
// will insert the item at the beginning, an index of 1 before the second item,
// and so on. An index of GetItemCount() or higher will insert the item at the
// end of the list. Negative indices are also allowed: An index of -1 will
// insert the item at the end of the list, an index of -2 before the last item,
// and so on. An index of -GetItemCount()-1 or lower will insert the item at the
// beginning.
//
// An item has a main text which will be highlighted when selected. It also has
// a secondary text which is shown underneath the main text (if it is set to
// visible) but which may remain empty.
//
// The shortcut is a key binding. If the specified rune is entered, the item
// is selected immediately. Set to 0 for no binding.
//
// The "selected" callback will be invoked when the user selects the item. You
// may provide nil if no such callback is needed or if all events are handled
// through the selected callback set with SetSelectedFunc().
//
// The currently selected item will shift its position accordingly. If the list
// was previously empty, a "changed" event is fired because the new item becomes
// selected.
func (l *List) InsertItem(index int, mainText string, selected func()) *List {
	item := &listItem{
		MainText: mainText,
	}

	// Shift index to range.
	if index < 0 {
		index = len(l.items) + index + 1
	}
	if index < 0 {
		index = 0
	} else if index > len(l.items) {
		index = len(l.items)
	}

	// Shift current item.
	if l.currentItem < len(l.items) && l.currentItem >= index {
		l.currentItem++
	}

	// Insert item (make space for the new item, then shift and insert).
	l.items = append(l.items, nil)
	if index < len(l.items)-1 { // -1 because l.items has already grown by one item.
		copy(l.items[index+1:], l.items[index:])
	}
	l.items[index] = item

	return l
}

// GetItemCount returns the number of items in the list.
func (l *List) GetItemCount() int {
	return len(l.items)
}

// GetItemText returns an item's texts (main and secondary). Panics if the index
// is out of range.
func (l *List) GetItemText(index int) (main string) {
	return l.items[index].MainText
}

// SetItemText sets an item's main and secondary text. Panics if the index is
// out of range.
func (l *List) SetItemText(index int, main string) *List {
	item := l.items[index]
	item.MainText = main
	return l
}

// FindItems searches the main and secondary texts for the given strings and
// returns a list of item indices in which those strings are found. One of the
// two search strings may be empty, it will then be ignored. Indices are always
// returned in ascending order.
//
// If mustContainBoth is set to true, mainSearch must be contained in the main
// text AND secondarySearch must be contained in the secondary text. If it is
// false, only one of the two search strings must be contained.
//
// Set ignoreCase to true for case-insensitive search.
func (l *List) FindItems(mainSearch string, mustContainBoth, ignoreCase bool) (indices []int) {
	if mainSearch == "" {
		return
	}

	if ignoreCase {
		mainSearch = strings.ToLower(mainSearch)
	}

	for index, item := range l.items {
		mainText := item.MainText
		if ignoreCase {
			mainText = strings.ToLower(mainText)
		}

		// strings.Contains() always returns true for a "" search.
		mainContained := strings.Contains(mainText, mainSearch)
		if mustContainBoth && mainContained ||
			!mustContainBoth && (mainText != "" && mainContained) {
			indices = append(indices, index)
		}
	}

	return
}

// Clear removes all items from the list.
func (l *List) Clear() *List {
	l.items = nil
	l.currentItem = 0
	return l
}

func (l *List) HasFocus() bool {
	return l.hasFocus
}

func (l *List) SetFocus(focus bool) {
	l.hasFocus = focus
}

// Draw draws this primitive onto the screen.
func (l *List) Draw(screen tcell.Screen) {
	l.Box.Draw(screen)

	// Determine the dimensions.
	x, y, width, height := l.GetInnerRect()
	bottomLimit := y + height

	// Adjust offset to keep the current selection in view.
	if l.currentItem < l.offset {
		l.offset = l.currentItem
	} else {
		if l.currentItem-l.offset >= height {
			l.offset = l.currentItem + 1 - height
		}
	}

	// Draw the list items.
	for index, item := range l.items {
		if index < l.offset {
			continue
		}

		if y >= bottomLimit {
			break
		}

		// Main text.
		defStyle := tcell.StyleDefault.
			Background(tcell.ColorBlack).
			Foreground(l.mainTextColor)
		emitStr(screen, x+1, y, defStyle, item.MainText)

		// Background color of selected text.
		if index == l.currentItem && l.HasFocus() {
			textWidth := width
			if !l.highlightFullLine {
				if w := len(item.MainText); w < textWidth {
					textWidth = w
				}
			}

			for bx := 0; bx < textWidth; bx++ {
				m, c, style, _ := screen.GetContent(x+bx, y)
				style = style.Background(l.selectedBackgroundColor).Foreground(l.selectedTextColor)
				screen.SetContent(x+bx, y, m, c, style)
			}
		}

		y++

		if y >= bottomLimit {
			break
		}
	}
}
