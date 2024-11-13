package main

import (
	"slices"
)

type Buffer struct {
	lines    []string
	modified bool
	index    int
}

func NewBuffer() *Buffer {
	return &Buffer{
		lines: make([]string, 0),
	}
}

func (b *Buffer) Clear() {
	b.lines = b.lines[:0]
	b.modified = false
	b.index = 0
}

func (b *Buffer) Append(line string) {
	b.lines = slices.Insert(b.lines, b.index, line)
	b.index++
}
