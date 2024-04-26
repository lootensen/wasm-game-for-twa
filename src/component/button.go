package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lootensen/wasm-demo-game/src/utils"
)

type Button struct {
	X           float32
	Y           float32
	Width       float32
	Height      float32
	JSInterface *JSInterface
}

type IButton interface {
	OnClick()
	Update()
	Draw()
}

func (b Button) Update() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		b.OnClick()
	}

}

func (b Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Width, b.Height, color.Black, false)
}

func (b Button) OnClick() {

	cursorPosX, cursorPosY := ebiten.CursorPosition()
	if utils.PositionInRectangle(cursorPosX, cursorPosY, int(b.X), int(b.Y), int(b.Width), int(b.Height)) {
		if b.JSInterface != nil {
			b.JSInterface.CallFunction("window.Telegram.WebApp.showAlert", "Hey there!")
		}
	}
}
