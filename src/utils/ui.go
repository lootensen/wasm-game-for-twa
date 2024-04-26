package utils

func PositionInRectangle(posX int, posY int, rectX int, rectY int, rectWidth int, rectHeight int) bool {
	if (posX >= rectX && posX <= rectX+rectWidth) && (posY >= rectY && posY <= rectY+rectHeight) {
		return true
	}
	return false
}
