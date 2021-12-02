package engine

const NE int = 4

var static [NE]bool = [NE]bool{false, true, false, false} //if entity is static, it cannot move

var position [NE][2]float64 = [NE][2]float64{{0.0, 0.0}, {26.0, 0.0}, {256.0, 32.0}, {300.0, 32.0}} //this must be defined inline for some reason

var inertia [NE][2]float64 = [NE][2]float64{{0.0, 0.0}, {0.0, 0.0}, {0.0, 0.0}, {0.0, 0.0}} //this must be defined inline for some reason

var size [NE][2]float64 = [NE][2]float64{{20.0, 32.0}, {32.0, 32.0}, {32.0, 32.0}, {32.0, 32.0}} // X and Y - width and height

/*
Todo:

add ability to remove entities

- auto shrink array to fit entities

add ability for array to extend to accomodate more entities
*/

// a simple hack because these will be synced

func GetNumberOfEntities() int {
	return len(static)
}

//static getters and setters

func GetStatic(index int) bool {
	return static[index]
}

func SetStatic(index int, state bool) {
	static[index] = state
}

//position getters

func GetPosition(index int) [2]float64 {
	return position[index]
}

func GetPositionX(index int) float64 {
	return position[index][0]
}

func GetPositionY(index int) float64 {
	return position[index][1]
}

//position setters

func SetPosition(index int, x float64, y float64) {
	position[index] = [2]float64{x, y}
}

func SetPositionX(index int, x float64) {
	position[index][0] = x
}

func SetPositionY(index int, y float64) {
	position[index][1] = y
}

//inertia getters

func GetInertia(index int) [2]float64 {
	return inertia[index]
}

func GetInertiaX(index int) float64 {
	return inertia[index][0]
}

func GetInertiaY(index int) float64 {
	return inertia[index][1]
}

//inertia setters

func SetInertia(index int, x float64, y float64) {
	inertia[index] = [2]float64{x, y}
}

func SetInertiaX(index int, x float64) {
	inertia[index][0] = x
}

func SetInertiaY(index int, y float64) {
	inertia[index][1] = y
}

//size getters

func GetSize(index int) [2]float64 {
	return size[index]
}

func GetWidth(index int) float64 {
	return size[index][0]
}

func GetHeight(index int) float64 {
	return size[index][1]
}

//size setters

func SetSize(index int, x float64, y float64) {
	size[index] = [2]float64{x, y}
}

func SetWidth(index int, x float64) {
	size[index][0] = x
}

func SetHeight(index int, y float64) {
	size[index][1] = y
}

//this is a bolt on hack to move the camera with the player

func getCameraPositionX() float64 {
	return -(position[0][0] - (640 / 2) - (size[0][0] / 2))
}

func getCameraPositionY() float64 {
	return -(position[0][1] - (480 / 2) - (size[0][1] / 2))
}
