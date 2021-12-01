package engine

var static [2]bool = [2]bool{false, true}

var position [2][2]float64 = [2][2]float64{{0.0, 0.0}, {50.0, 0.0}} //this must be defined inline for some reason

func getStatic(index int) bool {
	return static[index]
}

func setStatic(index int, state bool) {
	static[index] = state
}

func getPosition(index int) [2]float64 {
	return position[index]
}

func getX(index int) float64 {
	return position[index][0]
}

func getY(index int) float64 {
	return position[index][1]
}

func setPosition(index int, x float64, y float64) {
	position[index] = [2]float64{x, y}
}
