package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

type filesTodo struct {
	m Movies
	s Subtitles
}

func NewFilesTodo() *filesTodo {
	// log.Warn("==== NewFilesTodo Function ====")
	f := &filesTodo{}
	return f
}

func (f *filesTodo) FindData(path string) {
	// log.Warn("==== FindData Method ====")

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		ext_name := filepath.Ext(file.Name())
		if ext_name == ".mkv" || ext_name == ".avi" || ext_name == ".mp4" {
			f.m.InsertItem(file.Name())
		}
		if ext_name == ".srt" || ext_name == ".smi" || ext_name == ".smil" || ext_name == ".ass" {
			f.s.InsertItem(file.Name())
		}
	}
}

func (f *filesTodo) RenameIt() {
	// log.Warn("==== RenameIt Method ====")

	var movieNameSlice []string
	var subtitleExtSlice []string

	for _, file := range f.m.items {
		name := strings.TrimSuffix(file, filepath.Ext(file))
		movieNameSlice = append(movieNameSlice, name)
	}

	for _, file := range f.s.items {
		ext := filepath.Ext(file)
		subtitleExtSlice = append(subtitleExtSlice, ext)
	}

	var newName2 []string
	for i := 0; i < len(movieNameSlice); i++ {
		newName2 = append(newName2, movieNameSlice[i]+subtitleExtSlice[i])
	}

	// log.Info(newName2)

	for i := 0; i < len(f.s.items); i++ {
		os.Rename(f.s.items[i], newName2[i])
	}
}

// func main() {
// 	f := NewFilesTodo()
// 	f.FindData(".")

// 	log.Info(f.m.items)
// 	log.Info(f.s.items)

// 	f.RenameIt()

// 	log.Warn("==== test of TrimSuffix ====")
// 	basename := "hello.bash"
// 	name := strings.TrimSuffix(basename, filepath.Ext(basename))
// 	log.Info(name)
// }
