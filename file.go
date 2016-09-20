package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	chunkSize = 2048
)

type File struct {
	absolutePath string
	blobUUIDs    []string
	blobVector   []string
}

func (f *File) lock() error {
	var reader *bufio.Reader
	var file *os.File
	var chunks [][]byte
	var buf []byte
	var err error
	var n int

	file, err = os.Open(f.absolutePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buf = make([]byte, chunkSize)
	reader = bufio.NewReaderSize(file, chunkSize)
	chunks = make([][]byte, 0)

	for {
		n, err = reader.Read(buf)
		if err != nil {
			break
		}

		chunks = append(chunks, buf[0:n])
	}

	return nil
}

func (f *File) unlock() error {
	return nil
}

func (f *File) nextBlob() string {
	return ""
}

func main() {
	p := "books.txt"

	f := &File{
		absolutePath: p,
	}

	err := f.lock()
	if err != nil {
		fmt.Println(err)
	}

}
