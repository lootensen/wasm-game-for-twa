package app

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lootensen/wasm-demo-game/component"
)

type Game struct {
	btn component.Button
}

func (g *Game) Update() error {
	g.btn.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0x1f, 0xa0, 0xff})
	g.btn.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
func RunApp(jsInterface *component.JSInterface) {
	btn := &component.Button{
		X:           110,
		Y:           100,
		Width:       100,
		Height:      60,
		JSInterface: jsInterface,
	}
	game := &Game{
		btn: *btn,
	}
	ebiten.SetWindowSize(720, 1280)
	ebiten.SetWindowTitle("Fill")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
