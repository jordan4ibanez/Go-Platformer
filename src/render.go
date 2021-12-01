package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//cache reuse variables

var currentPosition [2]float64 = [2]float64{0, 0}

var currentSize [2]float64 = [2]float64{0, 0}

//

func DrawEntities(screen *ebiten.Image) {

	for i := 0; i < GetNumberOfEntities(); i++ {

		currentPosition = GetPosition(i)

		currentSize = GetSize(i)

		// entity debug: static are white, dynamic are red

		if GetStatic(i) {
			ebitenutil.DrawRect(screen, currentPosition[0], currentPosition[1], currentSize[0], currentSize[1], color.White)
		} else {
			ebitenutil.DrawRect(screen, currentPosition[0], currentPosition[1], currentSize[0], currentSize[1], color.RGBA{255, 0, 0, 255})
		}
	}

	//ebiten.SetWindowTitle(fmt.Sprintf("Hello, World! %v", CalculateFPS())) this basically chops the FPS down by 25-ish percent
}
