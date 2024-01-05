package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)
	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') {
			switch {
			case (p[i] >= 'A' && p[i] <= 'Z'):
				p[i] = 'A' + (p[i]-'A'+13)%26
			case (p[i] >= 'a' && p[i] <= 'z'):
				p[i] = 'a' + (p[i]-'a'+13)%26
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
