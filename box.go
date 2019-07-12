package main

import (
	"github.com/gdamore/tcell"
)

// Box implements gui design with a background and optional elements such as a title.
//
// See https://github.com/rivo/tview/wiki/Box for original Box implementation

type Box struct {
	// The position of the rect.
	x, y, width, height int

	// The inner rect reserved for the box's content.
	innerX, innerY, innerWidth, innerHeight int

	// Border padding.
	paddingTop, paddingBottom, paddingLeft, paddingRight int

	// The box's background color.
	backgroundColor tcell.Color

	// The title. Only visible if there is a border, too.
	title string

	// The color of the title.
	titleColor tcell.Color

	// default style
	defStyle tcell.Style
}

// NewBox returns a Box without a border.
func NewBox() *Box {
	b := &Box{
		width:           15,
		height:          10,
		innerX:          -1, // Mark as uninitialized.
		backgroundColor: tcell.ColorBlack,
		titleColor:      tcell.ColorYellow,
	}
	return b
}

// SetBorderPadding sets the size of the borders around the box content.
func (b *Box) SetBorderPadding(top, bottom, left, right int) *Box {
	b.paddingTop, b.paddingBottom, b.paddingLeft, b.paddingRight = top, bottom, left, right
	return b
}

// GetRect returns the current position of the rectangle, x, y, width, and
// height.
func (b *Box) GetRect() (int, int, int, int) {
	return b.x, b.y, b.width, b.height
}

// GetInnerRect returns the position of the inner rectangle (x, y, width,
// height), without the border and without any padding. Width and height values
// will clamp to 0 and thus never be negative.
func (b *Box) GetInnerRect() (int, int, int, int) {
	// default border true
	border := true
	if b.innerX >= 0 {
		return b.innerX, b.innerY, b.innerWidth, b.innerHeight
	}
	x, y, width, height := b.GetRect()
	if border {
		x++
		y++
		width -= 2
		height -= 2
	}
	x, y, width, height = x+b.paddingLeft,
		y+b.paddingTop,
		width-b.paddingLeft-b.paddingRight,
		height-b.paddingTop-b.paddingBottom
	if width < 0 {
		width = 0
	}
	if height < 0 {
		height = 0
	}
	return x, y, width, height
}

// SetRect sets a new position of the primitive. Note that this has no effect
// if this primitive is part of a layout (e.g. Flex, Grid) or if it was added
// like this:
//
//   application.SetRoot(b, true)
func (b *Box) SetRect(x, y, width, height int) {
	b.x = x
	b.y = y
	b.width = width
	b.height = height
	b.innerX = -1 // Mark inner rect as uninitialized.
}

// SetStyle sets a default style to style
func (b *Box) SetStyle(style tcell.Style) {
	b.defStyle = style
}

// SetBackgroundColor sets the box's background color.
func (b *Box) SetBackgroundColor(color tcell.Color) *Box {
	b.backgroundColor = color
	return b
}

// SetTitle sets the box's title.
func (b *Box) SetTitle(title string) *Box {
	b.title = title
	return b
}

// SetTitleColor sets the box's title color.
func (b *Box) SetTitleColor(color tcell.Color) *Box {
	b.titleColor = color
	return b
}

// Draw draws this primitive onto the screen.
func (b *Box) Draw(s tcell.Screen) {
	x2 := b.x + b.width - 1
	y2 := b.y + b.height - 1

	if y2 < b.y {
		b.y, y2 = y2, b.y
	}
	if x2 < b.x {
		b.x, x2 = x2, b.x
	}

	// Fill background
	for row := b.y + 1; row < y2; row++ {
		for col := b.x + 1; col < x2; col++ {
			s.SetContent(col, row, ' ', nil, b.defStyle)
		}
	}

	for col := b.x; col <= x2; col++ {
		s.SetContent(col, b.y, BoxDrawingsHeavyDoubleDashHorizontal, nil, b.defStyle)
		s.SetContent(col, y2, BoxDrawingsHeavyDoubleDashHorizontal, nil, b.defStyle)
	}

	for row := b.y + 1; row < y2; row++ {
		s.SetContent(b.x, row, BoxDrawingsDoubleVertical, nil, b.defStyle)
		s.SetContent(x2, row, BoxDrawingsDoubleVertical, nil, b.defStyle)
	}

	if b.y != y2 && b.x != x2 {
		// Only add corners if we need to
		s.SetContent(b.x, b.y, BoxDrawingsDoubleDownAndRight, nil, b.defStyle)
		s.SetContent(x2, b.y, BoxDrawingsDoubleDownAndLeft, nil, b.defStyle)
		s.SetContent(b.x, y2, BoxDrawingsDoubleUpAndRight, nil, b.defStyle)
		s.SetContent(x2, y2, BoxDrawingsDoubleUpAndLeft, nil, b.defStyle)
	}

	// Draw title.
	if b.title != "" && x2 >= 4 {
		//this is for 2 byte character, insert " " at the end of title and first of line
		title := " " + b.title + " "
		emitStr(s, b.x+1, b.y, tcell.StyleDefault.Foreground(b.titleColor), title)

	}
}
