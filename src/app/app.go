package app

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lootensen/wasm-demo-game/src/component"
)

type Game struct {
	btn []*component.Button
}

func (g *Game) Update() error {
	for _, btn := range g.btn {
		btn.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0x1f, 0xa0, 0xff})
	for _, btn := range g.btn {
		btn.Draw(screen)
	}
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
		OnClickHandler: func() {
			jsInterface.CallFunction("Telegram.WebApp.expand", nil)
		},
	}

	closeBtn := &component.Button{
		X:           300,
		Y:           60,
		Width:       60,
		Height:      60,
		JSInterface: jsInterface,
		OnClickHandler: func() {
			jsInterface.CallFunction("Telegram.WebApp.close", nil)
		},
	}
	game := &Game{
		btn: []*component.Button{btn, closeBtn},
	}
	ebiten.SetWindowSize(720, 1280)
	ebiten.SetWindowTitle("Fill")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
