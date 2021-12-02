package engine

const gravity float64 = 10

//cache hitters
var workerPosition [2]float64 = [2]float64{}
var workerPosition2 [2]float64 = [2]float64{}
var workerSize [2]float64 = [2]float64{}
var workerSize2 [2]float64 = [2]float64{}
var workerInertia [2]float64 = [2]float64{}
var workerTile [2]int = [2]int{}

//basically is the inverse of inclusive AABB - faster
func exclusionAABB(workerX float64, workerY float64) bool {
	//values were already cached - can be automatically reused
	return (
	//X calculation
	workerX+workerSize[0] > workerPosition2[0] &&
		workerX < workerPosition2[0]+workerSize2[0] &&
		//Y calculation
		workerY+workerSize[1] > workerPosition2[1] &&
		workerY < workerPosition2[1]+workerSize2[1])
}

func collisionDetectX(index int, workerX float64, workerY float64, inertiaX float64) float64 {

	workerSize = GetSize(index)

	//this is typed out twice because it is faster to check inertia X once rather than 1x amount of times
	if inertiaX < 0 {
		for i := 0; i < GetNumberOfEntities(); i++ {
			//do not detect against self
			if i != index {
				workerPosition2 = GetPosition(i)
				workerSize2 = GetSize(i)

				if exclusionAABB(workerX, workerY) {
					//fmt.Println("collide - x") //debug
					workerPosition[0] = workerPosition2[0] + workerSize2[0] + 0.005
					inertiaX = 0
				}
			}
		}
	} else if inertiaX > 0 {
		for i := 0; i < GetNumberOfEntities(); i++ {
			//do not detect against self
			if i != index {
				workerPosition2 = GetPosition(i)
				workerSize2 = GetSize(i)

				if exclusionAABB(workerX, workerY) {
					//fmt.Println("collide + x") //debug
					workerPosition[0] = workerPosition2[0] - workerSize[0] - 0.005
					inertiaX = 0
				}
			}
		}
	}

	return inertiaX
}

func collisionDetectY(index int, workerX float64, workerY float64, inertiaY float64) float64 {

	workerSize = GetSize(index)

	onGround = false

	//this is typed out twice because it is faster to check inertia X once rather than 1x amount of times
	if inertiaY < 0 {
		for i := 0; i < GetNumberOfEntities(); i++ {
			//do not detect against self
			if i != index {
				workerPosition2 = GetPosition(i)
				workerSize2 = GetSize(i)

				if exclusionAABB(workerX, workerY) {
					//fmt.Println("collide - x") //debug
					workerPosition[1] = workerPosition2[1] + workerSize2[1] + 0.005
					inertiaY = 0
				}
			}
		}
	} else if inertiaY > 0 {
		for i := 0; i < GetNumberOfEntities(); i++ {
			//do not detect against self
			if i != index {
				workerPosition2 = GetPosition(i)
				workerSize2 = GetSize(i)

				if exclusionAABB(workerX, workerY) {
					//fmt.Println("collide + x") //debug
					workerPosition[1] = workerPosition2[1] - workerSize[1] - 0.005
					inertiaY = 0
				}
			}
		}
	}

	return inertiaY
}

func collisionDetectXMap(index int, workerX float64, workerY float64, inertiaX float64) float64 {

	workerTile = [2]int{int(workerX / 32), int(workerY / 32)}

	for x := workerTile[0] - 1; x <= workerTile[0]+1; x++ {
		for y := workerTile[1] - 1; y <= workerTile[1]+1; y++ {

			//don't allow to go out of bounds
			if x >= 0 && x < GetMapLength() && y >= 0 && y < GetMapHeight() {

				//collide with the tiles
				if GameMap[y][x] > 0 {

					if inertiaX < 0 {
						workerPosition2 = [2]float64{float64(x) * 32, float64(y) * 32}
						workerSize2 = [2]float64{32, 32}

						if exclusionAABB(workerX, workerY) {
							workerPosition[0] = workerPosition2[0] + workerSize2[0] + 0.005
							inertiaX = 0
						}
					} else if inertiaX > 0 {

						workerPosition2 = [2]float64{float64(x) * 32, float64(y) * 32}
						workerSize2 = [2]float64{32, 32}

						if exclusionAABB(workerX, workerY) {
							//fmt.Println("collide + x") //debug
							workerPosition[0] = workerPosition2[0] - workerSize[0] - 0.005
							inertiaX = 0
						}

					}
				}
			}
		}
	}

	return inertiaX
}

func collisionDetectYMap(index int, workerX float64, workerY float64, inertiaY float64) float64 {

	workerTile = [2]int{int(workerX / 32), int(workerY / 32)}

	for x := workerTile[0] - 1; x <= workerTile[0]+1; x++ {
		for y := workerTile[1] - 1; y <= workerTile[1]+1; y++ {

			//don't allow to go out of bounds
			if x >= 0 && x < GetMapLength() && y >= 0 && y < GetMapHeight() {

				//collide with the tiles
				if GameMap[y][x] > 0 {

					if inertiaY < 0 {
						workerPosition2 = [2]float64{float64(x) * 32, float64(y) * 32}
						workerSize2 = [2]float64{32, 32}

						if exclusionAABB(workerX, workerY) {
							workerPosition[1] = workerPosition2[1] + workerSize2[1] + 0.005
							inertiaY = 0
						}
					} else if inertiaY > 0 {

						workerPosition2 = [2]float64{float64(x) * 32, float64(y) * 32}
						workerSize2 = [2]float64{32, 32}

						if exclusionAABB(workerX, workerY) {

							if index == 0 {
								onGround = true
							}
							//fmt.Println("collide + x") //debug
							workerPosition[1] = workerPosition2[1] - workerSize[1] - 0.005
							inertiaY = 0
						}

					}
				}
			}
		}
	}

	return inertiaY
}

func applyInertia(index int, dtime float64) {
	workerPosition = GetPosition(index)

	workerInertia = GetInertia(index)

	if workerInertia[1] < 50 {
		workerInertia[1] += gravity * dtime
	}

	//move X

	workerPosition[0] = workerPosition[0] + (workerInertia[0] * dtime)

	//detect X entities

	workerInertia[0] = collisionDetectX(index, workerPosition[0], workerPosition[1], workerInertia[0])

	//detect X map tiles

	workerInertia[0] = collisionDetectXMap(index, workerPosition[0], workerPosition[1], workerInertia[0])

	//move Y

	workerPosition[1] = workerPosition[1] + (workerInertia[1] * dtime)

	//detect Y entities

	workerInertia[1] = collisionDetectY(index, workerPosition[0], workerPosition[1], workerInertia[1])

	//detect Y map tiles

	workerInertia[1] = collisionDetectYMap(index, workerPosition[0], workerPosition[1], workerInertia[1])

	//set in memory

	SetPosition(index, workerPosition[0], workerPosition[1])

	SetInertia(index, workerInertia[0], workerInertia[1])

}

func RunPhysics(delta float64) {

	for i := 0; i < GetNumberOfEntities(); i++ {
		if !GetStatic(i) {
			applyInertia(i, delta)
		}
	}
}
