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
	l := len(b)
	p := make([]byte, l)

	n, err := r13.r.Read(p)
	if err != nil {
		return 0, err
	}

	for i := range n {
		code := p[i]

		if code < byte('A') || code > byte('z') {
			b[i] = code
		}

		if code > byte('A') && code < byte('Z') {
			code -= byte('A')
		} else {
			code -= byte('a')
		}

		if code > 12 {
			b[i] = p[i] - 13
		} else {
			b[i] = p[i] + 13
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
