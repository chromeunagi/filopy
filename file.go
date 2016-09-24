package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

/*
	TODO: optimization. for blobify, we can skip the intermediate "chunk"
	  step. we can also parallelize the blobify method, maybe. unless
	  a blockchain encryption scheme ends up being used. TBD.
*/

const (
	chunkSize = 2048
)

type (
	File struct {
		Locked       bool
		AbsolutePath string
		BlobVector   []string
	}
)

// Break the given file down into blobs.
func (f *File) blobify() ([]*Blob, error) {
	var reader *bufio.Reader
	var file *os.File
	var blobs []*Blob
	var b *Blob
	var chunks [][]byte
	var buf []byte
	var err error
	var n int

	file, err = os.Open(f.AbsolutePath)
	if err != nil {
		return nil, err
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

	blobs = make([]*Blob, 0)
	for _, c := range chunks {
		b, err = NewBlob(c)
		if err != nil {
			return nil, err
		}
		blobs = append(blobs, b)
	}

	return blobs, nil
}

func (f *File) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("%t%s", f.Locked, f.AbsolutePath))
	for _, id := range f.BlobVector {
		buf.WriteString(id)
	}
	return buf.String()
}

func (f *File) unlock() error {
	return nil
}

func (f *File) nextBlob() string {
	return ""
}
