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

func (b Button) Update(touchIds []ebiten.TouchID) {

	// Problem, nur fuer mobile, fuer click funzt das komplett nicht,
	// weil keine touchIDs
	for _, tId := range touchIds {
		if utils.ClickedOrTouched(tId) {
			b.OnClick(tId)
		}
	}

}

func (b Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Width, b.Height, color.Black, false)
}

func (b Button) OnClick(touchId ebiten.TouchID) {
	if b.OnClickHandler != nil {
		cursorPosX, cursorPosY := utils.ClickOrTouchPosition(touchId)
		if utils.PositionInRectangle(cursorPosX, cursorPosY, int(b.X), int(b.Y), int(b.Width), int(b.Height)) {
			b.OnClickHandler()
			// if b.JSInterface != nil {

			//
			// }
		}
	}
}
