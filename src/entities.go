package engine

var static [2]bool = [2]bool{false, true}

var position [2][2]float64 = [2][2]float64{{0.0, 0.0}, {50.0, 0.0}} //this must be defined inline for some reason

var inertia [2][2]float64 = [2][2]float64{{0.0, 0.0}, {0.0, 0.0}} //this must be defined inline for some reason

//static getters and setters

func getStatic(index int) bool {
	return static[index]
}

func setStatic(index int, state bool) {
	static[index] = state
}

//position getters

func getPosition(index int) [2]float64 {
	return position[index]
}

func getPositionX(index int) float64 {
	return position[index][0]
}

func getPositionY(index int) float64 {
	return position[index][1]
}

//position setters

func setPosition(index int, x float64, y float64) {
	position[index] = [2]float64{x, y}
}

func setPositionX(index int, x float64) {
	position[index][0] = x
}

func setPositionY(index int, y float64) {
	position[index][1] = y
}
