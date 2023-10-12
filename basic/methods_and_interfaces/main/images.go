package main

/*
package image

type Image interface {
// Возвращает цветовую модель изображения.
    ColorModel() color.Model
// Bounds возвращает домен, для которого At может возвращать ненулевой цвет.
// Границы не обязательно содержат точку (0, 0).
    Bounds() Rectangle
// At возвращает цвет пикселя в точке (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) возвращает верхний левый пиксель сетки.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) возвращает значение в правом нижнем углу.
    At(x, y int) color.Color
}

Изображение представляет собой конечную прямоугольную сетку цветов. Значения цвета взяты из цветовой модели.
*/

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
