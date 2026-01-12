package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := range n {
		c := b[i]
		switch {
		case 'A' <= c && c <= 'Z':
			b[i] = 'A' + (c-'A'+13)%26
		case 'a' <= c && c <= 'z':
			b[i] = 'a' + (c-'a'+13)%26
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
