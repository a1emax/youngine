package roundrect

import (
	"math"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/x/bitmap"
)

// Fill for rounded rectangle of given width and height with given axes of corner ellipses
// builds and returns new fill mask.
func Fill(width, height, a, b basic.Float) bitmap.Bitmap {

	// Modified Xiaolin Wu's algorithm.

	sx := int(basic.FloorPoz(width))
	sy := int(basic.FloorPoz(height))
	rxHypot := int(basic.FloorPoz(a))
	ryHypot := int(basic.FloorPoz(b))

	result := bitmap.New(sx, sy)

	color := grayColor(0xFF)

	if sx < 3 || sy < 3 {
		result.Fill(color)

		return result
	}

	rxMax := (sx - 3) / 2
	ryMax := (sy - 3) / 2

	rx := min(rxHypot, rxMax)
	ry := min(ryHypot, ryMax)

	if rx == 0 || ry == 0 {
		result.Fill(color)

		return result
	}

	cx := (sx - 2) / 2
	cy := (sy - 2) / 2
	xOff := rxMax - rx
	yOff := ryMax - ry

	rxSqr := float64(rx * rx)
	rySqr := float64(ry * ry)
	rLen := math.Sqrt(rxSqr + rySqr)

	qx := int(math.Round(rxSqr / rLen))
	for x := 0; x <= qx; x++ {
		fy := float64(ry) * math.Sqrt(1-float64(x*x)/rxSqr)
		iy := math.Floor(fy)
		y := int(iy)
		e := fy - iy

		outerColor := grayColor(uint8(math.Round(0xFF * e)))

		draw4Points(result, cx, cy, xOff+x, yOff+y+1, outerColor)

		draw4HLines(result, cx, cy, xOff, xOff+x, yOff+y, color)

		if x == 0 {
			draw4HLines(result, cx, cy, 0, xOff-1, cy, outerColor)
		}
	}

	qy := int(math.Round(rySqr / rLen))
	for y := 0; y <= qy; y++ {
		fx := float64(rx) * math.Sqrt(1-float64(y*y)/rySqr)
		ix := math.Floor(fx)
		x := int(ix)
		e := fx - ix

		outerColor := grayColor(uint8(math.Round(0xFF * e)))

		draw4Points(result, cx, cy, xOff+x+1, yOff+y, outerColor)

		draw4HLines(result, cx, cy, xOff, xOff+x, yOff+y, color)

		if y == 0 {
			draw4VLines(result, cx, cy, cx, 0, yOff-1, outerColor)
		}
	}

	result.Rect(cx-xOff+1, 1, cx+xOff-1, cy-yOff, color)
	result.Rect(1, cy-yOff+1, sx-3, cy+yOff-1, color)
	result.Rect(cx-xOff+1, cy+yOff, cx+xOff-1, sy-3, color)

	return result
}
