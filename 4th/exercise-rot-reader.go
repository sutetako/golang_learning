package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	if err == nil {
		for i := 0; i < len(p); i++ {
			if p[i] >= 'A' && p[i] <= 'Z' {
				p[i] = (p[i]-'A'+13)%26 + 'A'
			} else if p[i] >= 'a' && p[i] <= 'z' {
				p[i] = (p[i]-'a'+13)%26 + 'a'
			}
		}
	}
	return n, err
}

func rot_reader_main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
