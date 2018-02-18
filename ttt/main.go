package main

import (
	"flag"
	"github.com/fogleman/gg"
	"image/color"
	ttt "tic_tac_trinary"
)

var (
	canvas_x = flag.Int("canvas-x", 2300, "canvas width")
	canvas_y = flag.Int("canvas-y", 3500, "canvas height")
	font = flag.String("font", "KR Marker Thin.ttf", "TTF font file")
	width = flag.Int("width", 130, "board width")
)


func main() {
	flag.Parse()
	msg := "The only winning move is not to play   "
	if flag.NArg() > 0 {
		msg = flag.Arg(0)
	}
	outFile := "out.png"
	if flag.NArg() > 1 {
		outFile = flag.Arg(1)
	}

	dc := gg.NewContext(*canvas_x, *canvas_y)
	points := (float64(*width) / 7) * 96 / 72
	if err := dc.LoadFontFace(*font, points); err != nil {
		panic(err)
	}
	dc.SetLineWidth(3)
	dc.SetRGB255(255, 0, 0)
	dc.SetStrokeStyle(gg.NewSolidPattern(color.NRGBA { 255, 0, 0, 255 }))

	if err := ttt.RenderMessage(msg, *width, dc); err != nil {
		panic(err)
	}
	dc.SavePNG(outFile)
}
