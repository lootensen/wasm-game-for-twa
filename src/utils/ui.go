package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func PositionInRectangle(posX int, posY int, rectX int, rectY int, rectWidth int, rectHeight int) bool {
	if (posX >= rectX && posX <= rectX+rectWidth) && (posY >= rectY && posY <= rectY+rectHeight) {
		return true
	}
	return false
}

func ClickedOrTouched(touchID ebiten.TouchID) bool {
	if inpututil.IsTouchJustReleased(touchID) || inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		return true
	}
	return false
}

func ClickOrTouchPosition(touchID ebiten.TouchID) (x, y int) {
	if inpututil.IsTouchJustReleased(touchID) {
		return ebiten.TouchPosition(touchID)
	}
	return ebiten.CursorPosition()
}
