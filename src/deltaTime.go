package deltaTime

import (
	"fmt"
	"time"
)

var oldTime float64 = float64(time.Now().Local().UTC().UnixMicro())

var deltaTime float64 = 0

func currTimeToUnix() float64 {
	return float64(time.Now().Local().UTC().UnixMicro())
}

func calculateDelta() float64 {
	deltaTime = oldTime - currTimeToUnix()

	fmt.Println(deltaTime)

	return deltaTime
}
