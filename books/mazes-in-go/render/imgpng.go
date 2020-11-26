package render

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"mazes/maze"
)

type PngOptions struct {
	ww, pw int
	c      color.Color
}

//todo refactor
func ToPng(grid *maze.Grid, wr io.Writer) error {
	return ToPngOpt(grid, wr, PngOptions{4, 2, color.RGBA{R: 150, A: 255}})
}

func ToPngOpt(grid *maze.Grid, wr io.Writer, o PngOptions) error {
	w := grid.Width*(o.pw+o.ww) + o.ww
	h := grid.Height*(o.pw+o.ww) + o.ww
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	grid.EachRow(func(y int, row []*maze.Cell) {
		if y == 0 {
			renderTop(grid, row, 0, img, o)
		}
		renderMiddle(grid, row, y*o.pw+y*o.ww+o.ww, img, o)
		renderBottom(grid, row, y*o.pw+o.pw+y*o.ww+o.ww, img, o)
	})

	err := png.Encode(wr, img)
	if err != nil {
		return err
	}
	return nil
}

func renderTop(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA, o PngOptions) {
	for a := 0; a < o.ww; a++ {
		for b := 0; b < o.ww; b++ {
			img.Set(b, y+a, o.c)
		}
	}
	for i, cell := range row {
		if !cell.Linked(grid.North(cell)) {
			for a := 0; a < o.ww; a++ {
				img.Set(o.pw*i+o.ww*i+o.ww+0, y+a, o.c)
				img.Set(o.pw*i+o.ww*i+o.ww+1, y+a, o.c)
				img.Set(o.pw*i+o.ww*i+o.ww+2, y+a, o.c)
			}
		}
		for a := 0; a < o.ww; a++ {
			for b := 0; b < o.ww; b++ {
				img.Set((i*o.pw+o.pw)+(i*o.ww+o.ww)+b, y+a, o.c)
			}
		}
	}
}

func renderBottom(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA, o PngOptions) {
	for a := 0; a < o.ww; a++ {
		for b := 0; b < o.ww; b++ {
			img.Set(b, y+a, o.c)
		}
	}
	for i, cell := range row {
		if !cell.Linked(grid.South(cell)) {
			for a := 0; a < o.ww; a++ {
				img.Set(i*o.pw+i*o.ww+o.ww+0, y+a, o.c)
				img.Set(i*o.pw+i*o.ww+o.ww+1, y+a, o.c)
				img.Set(i*o.pw+i*o.ww+o.ww+2, y+a, o.c)
			}
		}
		for a := 0; a < o.ww; a++ {
			for b := 0; b < o.ww; b++ {
				img.Set((i*o.pw+o.pw)+(i*o.ww+o.ww)+b, y+a, o.c)
			}
		}
	}
}
func renderMiddle(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA, o PngOptions) {
	for a := 0; a < o.ww; a++ {
		img.Set(a, y+0, o.c)
		img.Set(a, y+1, o.c)
		img.Set(a, y+2, o.c)
	}
	for i, cell := range row {
		if !cell.Linked(grid.East(cell)) {
			for a := 0; a < o.ww; a++ {
				img.Set(i*o.pw+o.pw+i*o.ww+o.ww+a, y+0, o.c)
				img.Set(i*o.pw+o.pw+i*o.ww+o.ww+a, y+1, o.c)
				img.Set(i*o.pw+o.pw+i*o.ww+o.ww+a, y+2, o.c)
			}
		}
	}

}
