package main

import "fmt"

func main() {
	fmt.Printf("%s", Mandel(100))
}

type Point struct {
	x float64
	y float64
}

func Mandel(max_iter int32) string {
	width := 80
	height := 40
	x_min := -2.0
	x_max := 1.0
	y_min := -1.5
	y_max := 1.5
	x_step := (x_max - x_min) / float64(width)
	y_step := (y_max - y_min) / float64(height)

	var result = ""
	var y float64
	for y = 0; y < float64(height); y++ {
		var x float64
		for x = 0; x < float64(width); x++ {
			c := &Point{x_min + x_step*x, y_min + y_step*y}
			n := mandelbrot_set(c, max_iter)

			var pixel string
			switch {
			case 0 <= n && n <= 10:
				pixel = " "
			case 11 <= n && n <= 20:
				pixel = "."
			case 21 <= n && n <= 30:
				pixel = "+"
			case 31 <= n && n <= 40:
				pixel = "="
			case 41 <= n && n <= 50:
				pixel = "?"
			case 51 <= n && n <= 60:
				pixel = "#"
			case 61 <= n && n <= 70:
				pixel = ":"
			default:
				pixel = "*"
			}
			result = result + pixel
		}
		result += "\n"
	}
	return result
}

func mandelbrot_set(c *Point, max_iter int32) int32 {
	z := &Point{0.0, 0.0}
	var n int32 = 0
	for z.x*z.x+z.y*z.y < 4.0 && n < max_iter {
		z = &Point{z.x*z.x - z.y*z.y + c.x, 2.0*z.x*z.y + c.y}
		n++
	}
	return n
}
