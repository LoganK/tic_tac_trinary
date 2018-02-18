package tic_tac_trinary

import (
	"github.com/fogleman/gg"
	"log"
)

func RenderMessage(s string, width int, g *gg.Context) error {
	letters, err := TranslateToTrits(s)
	if err != nil {
		return err
	}
	letters_i := 0

	minPadding := width / 20
	xCount := g.Width() / (width + 2*minPadding)
	yCount := g.Height() / (width + 2*minPadding)

	// Justify the board placement with even padding.
	xPad := (g.Width() - (xCount * width)) / (xCount + 1)
	yPad := (g.Height() - (yCount * width)) / (yCount + 1)

	// Offset accounts for uneven numbers of pixels.
	xOff := (g.Width() - (width+xPad)*xCount) / 2
	yOff := (g.Height() - (width+yPad)*yCount) / 2

	log.Printf("Rendering: xCount=%d, yCount=%d, xPad=%d, yPad=%d", xCount, yCount, xPad, yPad)

	// Intentionally floor the value so we end up with an integer. (Partials are rendered as partially transparent.)
	third := float64(width / 3)

	yPos := float64(yOff)
	for y_i := 0; y_i < yCount; y_i++ {
		xPos := float64(xOff)
		for x_i := 0; x_i < xCount; x_i++ {
			g.DrawLine(xPos+third, yPos, xPos+third, yPos+float64(width))
			g.DrawLine(xPos+2*third, yPos, xPos+2*third, yPos+float64(width))
			g.DrawLine(xPos, yPos+third, xPos+float64(width), yPos+third)
			g.DrawLine(xPos, yPos+2*third, xPos+float64(width), yPos+2*third)

			xLet := xPos + (third / 2)
			for i := 0; i < 3; i++ {
				let := "   "
				if letters_i < len(letters) {
					let = letters[letters_i]
					letters_i++
				}

				yLet := yPos + (third / 2)
				g.DrawStringAnchored(string(let[0]), xLet, yLet, 0.5, 0.5)
				g.DrawStringAnchored(string(let[1]), xLet, yLet+third, 0.5, 0.5)
				g.DrawStringAnchored(string(let[2]), xLet, yLet+2*third, 0.5, 0.5)

				xLet += third
			}

			if letters_i >= len(letters) {
				letters_i = 0
			}

			xPos += float64(width + xPad)
		}
		yPos += float64(width + yPad)
	}
	g.Stroke()

	return nil
}
