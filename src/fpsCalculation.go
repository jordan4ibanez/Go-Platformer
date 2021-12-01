package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var averageFPS [100]float64 = [100]float64{}

var fpsIndex int = 0

var fpsCalc float64 = 0

func CalculateFPS() float64 {
	averageFPS[fpsIndex] = ebiten.CurrentFPS()

	fpsIndex++

	if fpsIndex >= 100 {
		fpsIndex = 0
	}

	fpsCalc = 0
	for i := 0; i < 100; i++ {
		fpsCalc += averageFPS[i]
	}

	fpsCalc /= 100

	//fmt.Println("-----------------------")

	//fmt.Println("Average:", fpsCalc)

	//fmt.Println("Current:", ebiten.CurrentFPS())

	return fpsCalc
}
