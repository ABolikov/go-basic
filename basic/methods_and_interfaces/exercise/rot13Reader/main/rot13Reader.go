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
	if err == io.EOF {
		return 0, err
	}

	for x := range p {
		p[x] = rot13(p[x])
	}
	return
}

// Алфавит: ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz (26 бук в алфавите)
func rot13(x byte) byte {
	isUp := x >= 'A' && x <= 'Z'
	//если символ не заглавный и не входит в рамки прописных букв - это не буква - символ остается таким какой есть
	if !isUp && (x < 'a' || x > 'z') {
		return x
	}
	//сдвиг значения согласно rot13
	x += 13

	//Проверка: вышли ли мы за границы допустимых символов, если да то необходимо отнять количество букв алфавита
	if isUp && x > 'Z' || !isUp && x > 'z' {
		x -= 26
	}
	return x
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
