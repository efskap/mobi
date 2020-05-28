package mobi

import (
	"bytes"
	"fmt"
)

type MobiChapter struct {
	Id           int
	Parent       int
	Title        string
	RecordOffset int
	LabelOffset  int
	Len          int
	Html         []uint8
	SubChapters  []*MobiChapter
	Depth int
}

func (w *MobiWriter) NewChapter(title string, text []byte) *MobiChapter {
	w.chapters = append(w.chapters, MobiChapter{Id: w.chapterCount, Title: title, Html: minimizeHTML(text), Depth: 1})
	w.chapterCount++
	return &w.chapters[len(w.chapters)-1]
}

func (w *MobiChapter) AddSubChapter(title string, text []byte) *MobiChapter {
	subChapter := &MobiChapter{Parent: w.Id, Title: title, Html: minimizeHTML(text), Depth: w.Depth + 1}
	w.SubChapters = append(w.SubChapters, subChapter)
	return subChapter
}

func (w *MobiChapter) SubChapterCount() int {
	return len(w.SubChapters)
}

func (w *MobiChapter) generateHTML(out *bytes.Buffer) {
	//Add check for unsupported HTML tags, characters, clean up HTML
	w.RecordOffset = out.Len()
	Len0 := out.Len()
	//fmt.Printf("Offset: --- %v %v \n", Len0, w.Title)

	out.WriteString(fmt.Sprintf("<h%d>%s</h%d>", w.Depth, w.Title, w.Depth))
	out.Write(w.Html)
	if len(w.Html) > 2000 {
		out.WriteString("<mbp:pagebreak/>")
	}
	w.Len = out.Len() - Len0
	for i, _ := range w.SubChapters {
		w.SubChapters[i].generateHTML(out)
	}
}
