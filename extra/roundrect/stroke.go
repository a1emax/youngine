package roundrect

import (
	"math"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/extra/bitmap"
)

// Stroke for rounded rectangle of given width and height with given axes of corner ellipses
// builds and returns new stroke mask of given thickness.
func Stroke(width, height, a, b, thickness basic.Float) bitmap.Bitmap {

	// Modified Xiaolin Wu's algorithm.

	sx := int(basic.FloorPoz(width))
	sy := int(basic.FloorPoz(height))
	rxHypot := int(basic.FloorPoz(a))
	ryHypot := int(basic.FloorPoz(b))

	result := bitmap.New(sx, sy)

	d := int(basic.FloorPoz(thickness)) - 1

	if d < 0 {
		return result
	}

	color := grayColor(0xFF)

	if sx < 3 || sy < 3 {
		result.Fill(color)

		return result
	}

	rxMax := (sx - 3) / 2
	ryMax := (sy - 3) / 2

	r1x := min(rxHypot, rxMax)
	r1y := min(ryHypot, ryMax)

	if r1x == 0 || r1y == 0 {
		result.Rect(0, 0, sx, d, color)
		result.Rect(0, sy-1-d, sx-1, sy-1, color)
		result.Rect(0, d+1, d, sy-1-d-1, color)
		result.Rect(sx-1-d, d+1, sx-1, sy-1-d-1, color)

		return result
	}

	r2x := basic.Poz(r1x - d)
	r2y := basic.Poz(r1y - d)

	cx := (sx - 2) / 2
	cy := (sy - 2) / 2
	xOff := rxMax - r1x
	yOff := ryMax - r1y

	r1xSqr := float64(r1x * r1x)
	r1ySqr := float64(r1y * r1y)
	r1Len := math.Sqrt(r1xSqr + r1ySqr)

	r2Zero := r2x == 0 || r2y == 0
	var r2xSqr, r2ySqr, r2Len float64
	if !r2Zero {
		r2xSqr = float64(r2x * r2x)
		r2ySqr = float64(r2y * r2y)
		r2Len = math.Sqrt(r2xSqr + r2ySqr)
	}

	qx1 := int(math.Round(r1xSqr / r1Len))
	for x1 := 0; x1 <= qx1; x1++ {
		fy1 := float64(r1y) * math.Sqrt(1-float64(x1*x1)/r1xSqr)
		iy1 := math.Floor(fy1)
		y1 := int(iy1)
		e1 := fy1 - iy1

		outerColor := grayColor(uint8(math.Round(0xFF * e1)))

		draw4Points(result, cx, cy, xOff+x1, yOff+y1+1, outerColor)

		var x2 int
		if r2Zero || y1 > r2y {
			x2 = -1
		} else {
			x2 = int(math.Floor(float64(r2x) * math.Sqrt(1-float64(y1*y1)/r2ySqr)))
		}
		draw4HLines(result, cx, cy, xOff+x2+1, xOff+x1, yOff+y1, color)

		if x1 == 0 {
			draw4HLines(result, cx, cy, 0, xOff-1, cy, outerColor)
		}
	}

	qy1 := int(math.Round(r1ySqr / r1Len))
	for y1 := 0; y1 <= qy1; y1++ {
		fx1 := float64(r1x) * math.Sqrt(1-float64(y1*y1)/r1ySqr)
		ix1 := math.Floor(fx1)
		x1 := int(ix1)
		e1 := fx1 - ix1

		outerColor := grayColor(uint8(math.Round(0xFF * e1)))

		draw4Points(result, cx, cy, xOff+x1+1, yOff+y1, outerColor)

		var x2 int
		if r2Zero || y1 > r2y {
			x2 = -1
		} else {
			x2 = int(math.Floor(float64(r2x) * math.Sqrt(1-float64(y1*y1)/r2ySqr)))
		}
		draw4HLines(result, cx, cy, xOff+x2+1, xOff+x1, yOff+y1, color)

		if y1 == 0 {
			draw4VLines(result, cx, cy, cx, 0, yOff-1, outerColor)
		}
	}

	if !r2Zero {
		qx2 := int(math.Round(r2xSqr / r2Len))
		for x2 := 0; x2 <= qx2; x2++ {
			fy2 := float64(r2y) * math.Sqrt(1-float64(x2*x2)/r2xSqr)
			iy2 := math.Floor(fy2)
			y2 := int(iy2)
			e2 := fy2 - iy2

			innerColor := grayColor(uint8(math.Round(0xFF * (1 - e2))))

			draw4Points(result, cx, cy, xOff+x2, yOff+y2, innerColor)

			if x2 == 0 {
				draw4HLines(result, cx, cy, 0, xOff-1, cy-d-1, innerColor)
			}
		}

		qy2 := int(math.Round(r2ySqr / r2Len))
		for y2 := 0; y2 <= qy2; y2++ {
			fx2 := float64(r2x) * math.Sqrt(1-float64(y2*y2)/r2ySqr)
			ix2 := math.Floor(fx2)
			x2 := int(ix2)
			e2 := fx2 - ix2

			innerColor := grayColor(uint8(math.Round(0xFF * (1 - e2))))

			draw4Points(result, cx, cy, xOff+x2, yOff+y2, innerColor)

			if y2 == 0 {
				draw4VLines(result, cx, cy, cx-d-1, 0, yOff-1, innerColor)
			}
		}
	}

	result.Rect(cx-xOff+1, 1, cx+xOff-1, d, color)
	result.Rect(cx-xOff+1, sy-2-d, cx+xOff-1, sy-3, color)
	result.Rect(1, cy-yOff+1, d, cy+yOff-1, color)
	result.Rect(sx-2-d, cy-yOff+1, sx-3, cy+yOff-1, color)

	return result
}
