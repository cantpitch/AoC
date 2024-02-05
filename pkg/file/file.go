package file

import (
	"bufio"
	"io"
	"log"
	"os"
)

type AOCFile interface {
	LoadFile() error
	Reset() error
	Next() (string, error)
}

type aocFile struct {
	fileName string
	first    *aocLine
	curr     *aocLine
}

type aocLine struct {
	text string
	next *aocLine
}

func NewFile(fileName string) AOCFile {
	f := &aocFile{
		fileName: fileName,
		first:    nil,
	}

	f.LoadFile()

	return f
}

func newLine(text string) *aocLine {
	return &aocLine{
		text: text,
		next: nil,
	}
}

func (f *aocFile) Reset() error {
	f.curr = f.first
	return nil
}

func (f *aocFile) Next() (string, error) {
	if f.curr == nil {
		return "", io.EOF
	}

	return f.curr.text, nil
}

func (f *aocFile) LoadFile() error {
	file, err := os.Open(f.fileName)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(bufio.NewReader(file))

	for scanner.Scan() {

		l := newLine(string(scanner.Text()))

		if f.first == nil {
			f.first = l
			f.curr = l
		} else {
			f.curr.next = l
			f.curr = l
		}
	}
	f.curr = f.first
	return nil
}
