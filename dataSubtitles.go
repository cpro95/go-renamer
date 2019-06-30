package main

type Subtitles struct {
	items []string
}

func NewSubtitles() *Subtitles {
	s := &Subtitles{}
	return s
}

func (s *Subtitles) InsertItem(name string) *Subtitles {
	s.items = append(s.items, name)
	return s
}

func (s *Subtitles) GetItemCount() int {
	return len(s.items)
}
