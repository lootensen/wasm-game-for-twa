package component

import (
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
	OnClick()
	Update()
	Draw()
}

func (b Button) Update(x int, y int) {

	// Problem, nur fuer mobile, fuer click funzt das komplett nicht,
	// weil keine touchIDs

	b.OnClick(x, y)

}

func (b Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Width, b.Height, color.Black, false)
}

func (b Button) OnClick(x int, y int) {
	if b.OnClickHandler != nil {
		// cursorPosX, cursorPosY := utils.ClickOrTouchPosition(touchId)
		if utils.PositionInRectangle(x, y, int(b.X), int(b.Y), int(b.Width), int(b.Height)) {
			b.OnClickHandler()
			// if b.JSInterface != nil {

			//
			// }
		}
	}
}
