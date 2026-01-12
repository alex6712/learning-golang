package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (myReader MyReader) Read(b []byte) (int, error) {
	l := len(b)
	for i := range l {
		b[i] = 'A'
	}
	return l, nil
}

func main() {
	reader.Validate(MyReader{})
}
