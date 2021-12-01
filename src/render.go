package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//cache reuse variables

var currentPosition [2]float64 = [2]float64{0, 0}

func DrawEntities(screen *ebiten.Image) {

	for i := 0; i < GetNumberOfEntities(); i++ {

		currentPosition = GetPosition(i)

		if GetStatic(i) {
			ebitenutil.DrawRect(screen, currentPosition[0], currentPosition[1], 20, 20, color.White)
		} else {
			ebitenutil.DrawRect(screen, currentPosition[0], currentPosition[1], 20, 20, color.RGBA{255, 0, 0, 255})
		}
	}
}
