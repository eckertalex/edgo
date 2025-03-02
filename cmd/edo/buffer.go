package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

type BufferIO interface {
	io.ReaderFrom
	io.WriterTo
}

// Compile-time check to ensure Buffer implements BufferIO
var _ BufferIO = (*Buffer)(nil)

type Buffer struct {
	lines []string
	index int
}

func NewBuffer() *Buffer {
	return &Buffer{
		lines: make([]string, 0),
	}
}

func (b *Buffer) ReadFrom(r io.Reader) (int64, error) {
	b.lines = b.lines[:0]
	b.index = 0

	var bytesRead int64
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		b.lines = append(b.lines, line)
		b.index++
		bytesRead += int64(len(line) + 1)
	}

	if err := scanner.Err(); err != nil {
		return bytesRead, err
	}

	return bytesRead, nil
}

func (b *Buffer) WriteTo(w io.Writer) (int64, error) {
	var bytesWritten int64
	for _, line := range b.lines {
		n, err := fmt.Fprintln(w, line)
		bytesWritten += int64(n)
		if err != nil {
			return bytesWritten, err
		}
	}
	return bytesWritten, nil
}

func (b *Buffer) Append(line string) {
	b.lines = slices.Insert(b.lines, b.index, line)
	b.index++
}

func (b *Buffer) Current() string {
	if len(b.lines) == 0 {
		return ""
	}

	return b.lines[b.index-1]
}
