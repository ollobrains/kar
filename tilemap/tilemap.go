package tilemap

import (
	"fmt"
	"image"
	"math"
)

type TileMap struct {
	Grid         [][]uint16
	W, H         int
	TileW, TileH int
}

func MakeTileMap(w, h, tileW, tileH int) *TileMap {
	return &TileMap{
		Grid:  MakeGrid(w, h),
		W:     w,
		H:     h,
		TileW: tileW,
		TileH: tileH,
	}
}

func NewTileMap(tm [][]uint16, tileW, tileH int) *TileMap {
	return &TileMap{
		Grid:  tm,
		W:     len(tm[0]),
		H:     len(tm),
		TileW: tileW,
		TileH: tileH,
	}
}

func (t *TileMap) String() string {
	s := ""
	for _, row := range t.Grid {
		for _, cell := range row {
			s += fmt.Sprintf("%d ", cell)
		}
		s += "\n"
	}
	return s
}

func MakeGrid(width, height int) [][]uint16 {
	var tm [][]uint16
	for i := 0; i < height; i++ {
		tm = append(tm, make([]uint16, width))
	}
	return tm
}

// IsSingleAxis checks if only one component of image.Point is non-zero
func IsSingleAxis(axis image.Point) bool {
	// True if exactly one of the components is non-zero
	return (axis.X != 0 && axis.Y == 0) || (axis.X == 0 && axis.Y != 0)
}

func (t *TileMap) Raycast(pos, dir image.Point, dist int) (image.Point, bool) {
	if IsSingleAxis(dir) {
		for range dist {
			pos = pos.Add(dir)
			if t.TileID(pos) != 0 {
				return pos, true
			}
		}
	} else {
		return image.Point{}, false
	}
	return image.Point{}, false
}

func (t *TileMap) WorldToTile(x, y float64) image.Point {
	return image.Point{int(math.Floor(x / float64(t.TileW))), int(math.Floor(y / float64(t.TileH)))}
}

func (t *TileMap) TileToWorld(pos image.Point) (float64, float64) {
	a := float64((pos.X * t.TileW) + t.TileW/2)
	b := float64((pos.Y * t.TileH) + t.TileH/2)
	return a, b
}

// func (t *TileMap) NearestBlockCenter(x, y float64) (float64, float64) {
// 	xCenter := math.Floor(x/float64(t.TileW)) + float64(t.TileW)/2
// 	yCenter := math.Floor(y/float64(t.TileH)) + float64(t.TileH)/2
// 	return xCenter, yCenter
// }

func (t *TileMap) SetTile(pos image.Point, id uint16) {
	if pos.X < 0 || pos.X >= t.W || pos.Y < 0 || pos.Y >= t.H {
		return
	}
	t.Grid[pos.Y][pos.X] = id
}
func (t *TileMap) TileID(pos image.Point) uint16 {
	if pos.X < 0 || pos.X >= t.W || pos.Y < 0 || pos.Y >= t.H {
		return 0
	}
	return t.Grid[pos.Y][pos.X]
}

func (t *TileMap) GetTileRect(pos image.Point) (x, y, w, h float64) {
	return float64(pos.X * t.TileW), float64(pos.Y * t.TileH), float64(t.TileW), float64(t.TileH)
}
