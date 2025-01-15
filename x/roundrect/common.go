package roundrect

import (
	"github.com/a1emax/youngine/x/bitmap"
	"github.com/a1emax/youngine/x/colors"
)

// grayColor returns RGBA color with all channels set to c.
func grayColor(c uint8) colors.RGBA {
	return colors.RGBA{c, c, c, c}
}

// draw4Points draws points distanced by (x, y) from (cx, cy) in all four directions.
func draw4Points(dst bitmap.Bitmap, cx, cy, x, y int, color colors.RGBA) {
	dst.Set(cx-x, cy-y, color)
	dst.Set(cx+x, cy-y, color)
	dst.Set(cx-x, cy+y, color)
	dst.Set(cx+x, cy+y, color)
}

// draw4HLines draws lines between points distanced by (x1, y) and (x2, y) from (cx, cy) in all four directions.
func draw4HLines(dst bitmap.Bitmap, cx, cy, x1, x2, y int, color colors.RGBA) {
	dst.Lines(x2-x1+1, [][2]int{
		{cx - x2, cy - y},
		{cx + x1, cy - y},
		{cx - x2, cy + y},
		{cx + x1, cy + y},
	}, color)
}

// draw4VLines draws lines between points distanced by (x, y1) and (x, y2) from (cx, cy) in all four directions.
func draw4VLines(dst bitmap.Bitmap, cx, cy, x, y1, y2 int, color colors.RGBA) {
	for y := y1; y <= y2; y++ {
		draw4Points(dst, cx, cy, x, y, color)
	}
}
