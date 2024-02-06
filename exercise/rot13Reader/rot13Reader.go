package main

import (
	"io"
	"os"
	"strings"
)

// rot13 ... アルファベットの文字を13文字分ずらすこと

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := range b[:n] {
		b[i] = move(b[i])
	}
	return n, err

}

func move(b byte) byte {
	switch {
		case 'a' <= b && b <= 'z':
			return 'a' + (b - 'a' + 13) % 26
		case 'A' <= b && b <= 'Z':
			return 'A' + (b-'A'+13)%26
		default:
			return b
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // ここで Read() が呼び出される
}
