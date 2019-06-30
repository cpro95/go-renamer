package main

import "strings"

type Movies struct {
	items []string
}

func NewMovies() *Movies {
	m := &Movies{}
	return m
}

func (m *Movies) InsertItem(name string) *Movies {
	m.items = append(m.items, name)
	return m
}

func (m *Movies) GetItemCount() int {
	return len(m.items)
}

func (m *Movies) ToString() string {
	tempString := strings.Join(m.items, " ")
	return tempString
}
