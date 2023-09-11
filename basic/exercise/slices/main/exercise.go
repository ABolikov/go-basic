package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println("dy = ", dy)
	//создаем слайс с длинной dy
	y := make([][]uint8, dy)

	//перебираем созданный слайс и в каждое значение пишем новый слайс с длиной dx
	for i := range y {
		fmt.Println("dx = ", dx)
		y[i] = make([]uint8, dx)
	}

	return y
}

func main() {
	pic.Show(Pic)
}
