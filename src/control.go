package engine

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//entity 0 will always be the player

const maxPlayerSpeed = 10

var upKey bool = false
var downKey bool = false
var leftKey bool = false
var rightKey bool = false

var playerInertia [2]float64 = [2]float64{}

func PlayerControlInput(dtime float64) {
	//poll inputs
	upKey = ebiten.IsKeyPressed(ebiten.KeyUp)
	downKey = ebiten.IsKeyPressed(ebiten.KeyDown)
	leftKey = ebiten.IsKeyPressed(ebiten.KeyLeft)
	rightKey = ebiten.IsKeyPressed(ebiten.KeyRight)

	playerInertia = GetInertia(0)

	//X axis
	if leftKey && !rightKey {
		if playerInertia[0] > -maxPlayerSpeed {
			playerInertia[0] -= (dtime / 2)
		}

		if playerInertia[0] < -maxPlayerSpeed {
			playerInertia[0] = -maxPlayerSpeed
		}
	} else if !leftKey && rightKey {
		if playerInertia[0] < maxPlayerSpeed {
			playerInertia[0] += (dtime / 2)
		}

		if playerInertia[0] > maxPlayerSpeed {
			playerInertia[0] = maxPlayerSpeed
		}
	} else {

		//slow down - smoothly
		playerInertia[0] += -playerInertia[0] * (deltaTime / 4)

		//stop jittering
		if math.Abs(playerInertia[0]) < 0.0005 {
			playerInertia[0] = 0
		}
	}

	//Y axis
	if upKey && !downKey {
		if playerInertia[1] > -maxPlayerSpeed {
			playerInertia[1] -= (dtime / 2)
		}

		if playerInertia[1] < -maxPlayerSpeed {
			playerInertia[1] = -maxPlayerSpeed
		}
	} else if !upKey && downKey {
		if playerInertia[1] < maxPlayerSpeed {
			playerInertia[1] += (dtime / 2)
		}

		if playerInertia[1] > maxPlayerSpeed {
			playerInertia[1] = maxPlayerSpeed
		}
	} else {

		//slow down - smoothly
		playerInertia[1] += -playerInertia[1] * (deltaTime / 4)

		//stop jittering
		if math.Abs(playerInertia[1]) < 0.0005 {
			playerInertia[1] = 0
		}
	}

	SetInertia(0, playerInertia[0], playerInertia[1])
}
