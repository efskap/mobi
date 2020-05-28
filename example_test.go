package mobi_test

import (
	"fmt"
	"github.com/efskap/mobi"
	"strings"
	"os"
	"image/jpeg"

)

func ExampleNewWriter() {
	w, err := mobi.NewWriter("/tmp/example.mobi")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	w.Title("Book Title")
	w.Compression(mobi.CompressionNone) // LZ77 compression is also possible using  mobi.CompressionPalmDoc

	// Meta data
	w.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	w.NewExthRecord(mobi.EXTH_AUTHOR, "Book Author Name")
	// See exth.go for additional EXTH record IDs

	// Add chapters and subchapters
	ch1 := w.NewChapter("Chapter 1", []byte("Some text here"))
	ch1.AddSubChapter("Chapter 1-1", []byte("Some text here"))
	ch1.AddSubChapter("Chapter 1-2", []byte("Some text here"))

	w.NewChapter("Chapter 2", []byte("Some text here")).AddSubChapter("Chapter 2-1", []byte("Some text here")).AddSubChapter("Chapter 2-2", []byte("Some text here"))
	w.NewChapter("Chapter 3", []byte("Some text here")).AddSubChapter("Chapter 3-1", []byte("Some text here"))
	w.NewChapter("Chapter 4", []byte("Some text here")).AddSubChapter("Chapter 4-1", []byte("Some text here"))

	// Output MOBI File
	w.Write()

	// Output:
}

func ExampleNewReader() {
	r, err := mobi.NewReader("/tmp/example.mobi")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	fmt.Printf("Document type: %s\n", r.DocType())
	fmt.Printf("Author(s): %s\n", strings.Join(r.Authors(), ", "))
	fmt.Printf("Title: %s\n", r.BestTitle())
	// See reader.go for additional record types

	if r.HasCover() {
		img, _ := r.Cover()
		f, _ := os.Create("/tmp/mobi_example.jpg")
		jpeg.Encode(f, img, nil)
		f.Close()
	}

	// Output: Document type: EBOK
	//Author(s): Book Author Name
	//Title: Book Title
}
