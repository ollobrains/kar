package tilemap

import (
	"fmt"
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

func Raycast(tm [][]uint16, x, y int, dirX, dirY int) (pos [2]int, id uint16, ok bool) {
	cursorX, cursorY := x, y
	for range 3 {
		cursorX += dirX
		cursorY += dirY
		if tm[cursorY][cursorX] != 0 {
			return [2]int{cursorX, cursorY}, tm[cursorY][cursorX], true
		}
	}
	return [2]int{cursorX, cursorY}, 0, false
}

// GetTileCoords, verilen piksel koordinatlarına karşılık gelen döşeme koordinatlarını döndürür
func (t *TileMap) GetTileCoords(x, y float64) (int, int) {
	return int(math.Floor(x / float64(t.TileW))), int(math.Floor(y / float64(t.TileH)))
}

// GetTileCoordsFromCenter, verilen karakter boyutlarını kullanarak merkez noktadan döşeme koordinatlarını hesaplar
func (t *TileMap) GetTileCoordsFromCenter(x, y, w, h float64) (int, int) {
	return int(math.Floor((x + w/2) / float64(t.TileW))), int(math.Floor((y + h/2) / float64(t.TileH)))
}

func (t *TileMap) SetTile(x, y int, id uint16) {
	if x < 0 || x >= t.W || y < 0 || y >= t.H {
		return
	}
	t.Grid[y][x] = id
}
