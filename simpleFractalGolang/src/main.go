package main

import (
	"flag"
	"fmt"

	"github.com/fogleman/gg"
)

var (
	xy     bool
	help   bool
	width  int
	height int
	div    int
	rad    float64
)

func drawCircle(x float64, y float64, radius float64, dc *gg.Context) {
	dc.SetRGB255(int(radius)%8*32, 0, int(radius)%64*128)
	dc.DrawCircle(x, y, radius)
	dc.Stroke()
	if radius > 1 {
		radius /= float64(div)
		//x
		drawCircle((x + radius), y, radius, dc)
		drawCircle((x - radius), y, radius, dc)
		if xy {
			//y
			drawCircle(x, (y + radius), radius, dc)
			drawCircle(x, (y - radius), radius, dc)
		}

	}
}
func draw() {
	dc := gg.NewContext(width, height)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	drawCircle(float64(width)/2, float64(height)/2, rad, dc)
	if xy {
		dc.SavePNG("../images/circleFractalXY.png")
	} else {
		dc.SavePNG("../images/circleFractal.png")
	}

}

func init() {
	//commands
	flag.BoolVar(&xy, "xy", false, "xy")
	flag.BoolVar(&help, "help", false, "h")

	flag.IntVar(&div, "div", 2, "division")
	flag.IntVar(&width, "w", 1000, "width")
	flag.IntVar(&height, "h", 500, "height")

	flag.Float64Var(&rad, "r", 100, "radius")

	flag.Parse()

}
func main() {
	if !help {
		fmt.Println("-help")
		fmt.Printf("the radius is \033[31m%f\n\033[0m", rad)
		draw()
		fmt.Println("\033[34mfractal circle finished\033[0m")
	} else {
		fmt.Println("	\033[34m-help         \033[31muse this command for get help\n	\033[34m-xy           \033[31mthis will be execute x & y\n	\033[34m-r[number]    \033[31mis the radius of the circle\n	\033[34m-div[number]  \033[31mdivide the radius\n	\033[34m-w[number]    \033[31mwidth of the image\n	\033[34m-h[number]    height of the image \033[0m")
	}

}
