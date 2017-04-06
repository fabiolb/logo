package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	var (
		// box dimensions
		bW   = 80     // box width
		bH   = bW / 2 // box height
		bRad = 10     // box corner radius
		bPad = 16     // box padding

		// other dimensions
		botY    = bH + 4*bPad      // top of bottom row
		centerX = bW + bPad + bW/2 // canvas center

		// canvas dimensions
		width  = 749 // by try and error
		height = 145

		// colors
		darkBlue  = "#4294d3"
		lightBlue = "#21bdf0"

		// styles
		darkBoxStyle  = "fill:" + darkBlue
		lightBoxStyle = "fill:" + lightBlue
		lineStyle     = "stroke-width:6;stroke-linecap:round;stroke:" + lightBlue
		textStyle     = "text-anchor:start;font-size:160px;font-family:MicrogrammaDMedExt;fill:" + darkBlue
	)

	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Title("Fabio")

	// use this to determine the width
	// canvas.Rect(0, 0, width, height, "fill:red")

	// fabio
	tX := 3*bW + 4*bPad
	tY := botY + bH
	canvas.Text(tX, tY, "Fabio", textStyle)

	// boxes

	// box draws a box
	box := func(x, y int, style string) {
		canvas.Roundrect(x, y, bW, bH, bRad, bRad, style)
	}

	// top row box
	box(bW+bPad, 0, darkBoxStyle)

	// bottom row
	box(0, botY, darkBoxStyle)
	box(bW+bPad, botY, darkBoxStyle)
	box(2*bW+2*bPad, botY, lightBoxStyle)

	// lines

	// center vertical row
	cW, cH := 0, 6*bPad/2
	cX1, cY1 := centerX, bH+bPad/2
	cX2, cY2 := cX1+cW, cY1+cH
	canvas.Line(cX1, cY1, cX2, cY2, lineStyle)

	// left vertical row
	lvW, lvH := 0, 2*bPad/2
	lvX1, lvY1 := bW/2, bH+5*bPad/2
	lvX2, lvY2 := lvX1+lvW, lvY1+lvH
	canvas.Line(lvX1, lvY1, lvX2, lvY2, lineStyle)

	// right vertical row
	rvW, rvH := 0, lvH
	rvX1, rvY1 := 2*bW+2*bPad+bW/2, lvY1
	rvX2, rvY2 := rvX1+rvW, rvY1+rvH
	canvas.Line(rvX1, rvY1, rvX2, rvY2, lineStyle)

	// top cross row
	tcW, tcH := bW+2*bPad, 0
	tcX1, tcY1 := bW, bH+3*bPad/2
	tcX2, tcY2 := tcX1+tcW, tcY1+tcH
	canvas.Line(tcX1, tcY1, tcX2, tcY2, lineStyle)

	// diag left
	canvas.Line(tcX1, tcY1, lvX1, lvY1, lineStyle)

	// diag right
	canvas.Line(tcX2, tcY2, rvX1, rvY1, lineStyle)

	canvas.End()
}
