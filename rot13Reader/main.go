package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	for i := range b {
		rot13(&b[i])
	}
	return
}

func rot13(b *byte) {
	if *b > 'A' && *b < 'Z' {
		*b = 'A' + ((*b - 'A' + 13) % 26)
	} else if *b > 'a' && *b < 'z' {
		*b = 'a' + ((*b - 'a' + 13) % 26)
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
