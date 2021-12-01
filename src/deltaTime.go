package engine

import (
	"time"
)

var oldTime float64 = float64(time.Now().Local().UTC().UnixMicro())

var currentTime float64 = oldTime

var deltaTime float64 = 0

func currTimeToUnix() float64 {
	return float64(time.Now().Local().UTC().UnixMicro())
}

func CalculateDelta() float64 {

	currentTime = currTimeToUnix()

	deltaTime = currentTime - oldTime

	oldTime = currentTime

	return deltaTime
}
