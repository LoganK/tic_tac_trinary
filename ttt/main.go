package main

import (
	"flag"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"image/color"
	"io/ioutil"
	ttt "tic_tac_trinary"
)

var (
	canvas_x = flag.Int("canvas-x", 2300, "canvas width")
	canvas_y = flag.Int("canvas-y", 3500, "canvas height")
	fontFile = flag.String("font", "", "TTF font file")
	width    = flag.Int("width", 130, "board width")
)

func loadFontFace(size float64) (font.Face, error) {
	fontBytes := gobold.TTF
	if *fontFile != "" {
		fileBytes, err := ioutil.ReadFile(*fontFile)
		if err != nil {
			return nil, err
		}
		fontBytes = fileBytes
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(f, &truetype.Options{
		Size: size,
	}), nil
}

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

	glyphWidth := float64(*width) / 7
	face, err := loadFontFace(glyphWidth)
	if err != nil {
		panic("Failed to load font: " + err.Error())
	}

	dc.SetFontFace(face)
	dc.SetLineWidth(3)
	dc.SetRGB255(255, 0, 0)
	dc.SetStrokeStyle(gg.NewSolidPattern(color.NRGBA{255, 0, 0, 255}))

	if err := ttt.RenderMessage(msg, *width, dc); err != nil {
		panic(err)
	}
	dc.SavePNG(outFile)
}
