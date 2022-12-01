package main

import (
	"bytes"
	"io"
)

type Trimmer struct {
	doneTrimming bool
	src          io.Reader
	seekChar     byte
}

func NewTrimmer(src io.Reader, seekChar byte) *Trimmer {
	return &Trimmer{
		src:      src,
		seekChar: seekChar,
	}
}

func (t *Trimmer) Read(p []byte) (n int, err error) {
	n, err = t.src.Read(p)
	if t.doneTrimming {
		return
	}

	buf := p[:n]
	buf = bytes.TrimLeftFunc(buf, func(r rune) bool {
		return r != rune(t.seekChar)
	})

	if len(buf) > 0 {
		t.doneTrimming = true
	}

	n = copy(p, buf)

	return
}
