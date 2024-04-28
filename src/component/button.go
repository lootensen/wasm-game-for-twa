package component

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lootensen/wasm-demo-game/src/utils"
)

type Button struct {
	X              float32
	Y              float32
	Width          float32
	Height         float32
	JSInterface    *JSInterface
	OnClickHandler func()
}

type IButton interface {
	Update()
	Draw()
}

func (b Button) Update(x int, y int) {
	// Problem, nur fuer mobile, fuer click funzt das komplett nicht,
	// weil keine touchIDs
	if utils.PositionInRectangle(x, y, int(b.X), int(b.Y), int(b.Width), int(b.Height)) {
		fmt.Println("in Rect")
		b.OnClickHandler()
	}
}

func (b Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Width, b.Height, color.Black, false)
}
