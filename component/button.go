package component

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
		fmt.Println("clicked")
		b.OnClick()
	}

}

func (b Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Width, b.Height, color.Black, false)
}

func (b Button) OnClick() {
	if b.JSInterface != nil {
		b.JSInterface.CallFunction("console.log", "My New JS Interface works") // Rufe eine JS-Funktion auf
	}
}
