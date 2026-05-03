package utils

import "github.com/gdamore/tcell/v3"

func PrintText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}

func PrintRectangle(screen tcell.Screen, style *tcell.Style, topY, bottomY, leftX, rightX int) {
	for i := leftX + 1; i < rightX; i++ {
		screen.SetContent(i, topY, '-', nil, *style)
		screen.SetContent(i, bottomY, '-', nil, *style)
	}

	for i := topY + 1; i < bottomY; i++ {
		screen.SetContent(leftX, i, '|', nil, *style)
		screen.SetContent(rightX, i, '|', nil, *style)
	}
	screen.SetContent(leftX, topY, '┌', nil, *style)
	screen.SetContent(rightX, topY, '┐', nil, *style)
	screen.SetContent(leftX, bottomY, '└', nil, *style)
	screen.SetContent(rightX, bottomY, '┘', nil, *style)
}

func FillBackground(screen tcell.Screen, color tcell.Color) {
	style := tcell.StyleDefault.Background(color)
	width, height := screen.Size()
	for y := range height {
		for x := range width {
			screen.SetContent(x, y, ' ', nil, style)
		}
	}
}
