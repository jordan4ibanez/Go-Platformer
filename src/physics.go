package engine

const gravity float64 = 10

//cache hitters
var workerPosition [2]float64 = [2]float64{}
var workerPosition2 [2]float64 = [2]float64{}
var workerSize [2]float64 = [2]float64{}
var workerSize2 [2]float64 = [2]float64{}
var workerInertia [2]float64 = [2]float64{}

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

func applyInertia(index int, dtime float64) {
	workerPosition = GetPosition(index)

	workerInertia = GetInertia(index)

	//move X

	workerPosition[0] = workerPosition[0] + (workerInertia[0] * dtime)

	//detect X

	workerInertia[0] = collisionDetectX(index, workerPosition[0], workerPosition[1], workerInertia[0])

	//move Y
	workerPosition[1] = workerPosition[1] + (workerInertia[1] * dtime)

	//detect Y

	workerInertia[1] = collisionDetectX(index, workerPosition[0], workerPosition[1], workerInertia[1])

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
