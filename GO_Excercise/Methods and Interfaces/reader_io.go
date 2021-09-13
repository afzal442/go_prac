package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = b[i] - 'A' + 'N'
			b[i] = b[i]%26 + 'A'
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] - 'a' + 'n'
			b[i] = b[i]%26 + 'a'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
