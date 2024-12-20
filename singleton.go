package kar

import (
	"image/color"
	"kar/items"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/mlange-42/arche/ecs"
	"github.com/setanarut/kamera/v2"
)

type ISystem interface {
	Init()
	Update()
	Draw()
}

const TimerTick time.Duration = time.Second / 60
const DeltaTime float64 = 1.0 / 60.0

var (
	ScreenW, ScreenH = 854.0, 480.0
	Screen           *ebiten.Image
	Camera           = kamera.NewCamera(400, 450, ScreenW, ScreenH)
	WorldECS         = ecs.NewWorld()
	// DesktopPath      string
	GlobalDIO         = &colorm.DrawImageOptions{}
	GlobalColorM      = colorm.ColorM{}
	BackgroundColor   = rgb(38, 0, 121)
	DrawDebugHitboxes = false
	ItemScale         = 1.0
	PlayerScale       = 1.0
)
var ItemColorMap = map[uint16]color.RGBA{
	items.Air:        rgb(1, 1, 1),
	items.GrassBlock: rgb(0, 186, 53),
	items.Dirt:       rgb(133, 75, 54),
	items.Sand:       rgb(199, 193, 158),
	items.Stone:      rgb(139, 139, 139),
	items.CoalOre:    rgb(0, 0, 0),
	items.GoldOre:    rgb(255, 221, 0),
	items.IronOre:    rgb(171, 162, 147),
	items.DiamondOre: rgb(0, 247, 255),
}

func init() {
	// GlobalColorM.ChangeHSV(1, 0, 1)
	// Camera.SmoothOptions.LerpSpeedX = 0.09
	// Camera.SmoothOptions.LerpSpeedY = 0.02
	// Camera.SmoothingOptions.SmoothDampTimeX = 0.3
	// Camera.SmoothingOptions.SmoothDampTimeY = 1
	// Camera.SmoothingOptions.SmoothDampMaxSpeedX = 1000
	// Camera.SmoothingOptions.SmoothDampMaxSpeedY = 1000

	// homePath, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// DesktopPath = homePath + "/Desktop/"
}

func rgb(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 255}
}
