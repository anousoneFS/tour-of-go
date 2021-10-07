package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// automatically Read() function because it implement io.Reader interface
func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if p[i] >= 'A' && p[i] < 'Z' {
			p[i] = 65 + (((p[i] - 65) + 13) % 26)
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 97 + (((p[i] - 97) + 13) % 26)
		}
	}
	return
}

func Rot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s} // type r is io.Reader
	io.Copy(os.Stdout, &r)
}
