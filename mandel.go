package main

import "fmt"

func main() {
	fmt.Printf("%s", string(Mandel(1000000)))
}

type Point struct {
	x float64
	y float64
}

const width = 80
const height = 40

func Mandel(maxIter int32) []rune {
	xMin := -2.0
	xMax := 1.0
	yMin := -1.5
	yMax := 1.5
	xStep := (xMax - xMin) / float64(width)
	yStep := (yMax - yMin) / float64(height)

	result := make([]rune, (width+1)*height)
	var y float64
	var x float64
	c := &Point{0, 0}
	i := 0
	for y = 0; y < float64(height); y++ {
		for x = 0; x < float64(width); x++ {
			c.x, c.y = xMin+xStep*x, yMin+yStep*y
			result[i] = pixel2rune(mandelbrotSet(c, maxIter))
			i++
		}
		result[i] = '\n'
		i++
	}
	return result
}

func pixel2rune(n int32) rune {
	switch {
	case 0 <= n && n <= 10:
		return ' '
	case 11 <= n && n <= 20:
		return '.'
	case 21 <= n && n <= 30:
		return '+'
	case 31 <= n && n <= 40:
		return '='
	case 41 <= n && n <= 50:
		return '?'
	case 51 <= n && n <= 60:
		return '#'
	case 61 <= n && n <= 70:
		return ':'
	default:
		return '*'
	}
}

func mandelbrotSet(c *Point, maxIter int32) int32 {
	z := &Point{0.0, 0.0}
	var n int32 = 0
	for z.x*z.x+z.y*z.y < 4.0 && n < maxIter {
		z.x, z.y = z.x*z.x-z.y*z.y+c.x, 2.0*z.x*z.y+c.y
		n++
	}
	return n
}
