package engine

const gravity float64 = 10

var workerPosition [2]float64 = [2]float64{}
var workerInertia [2]float64 = [2]float64{}

func applyInertia(index int, dtime float64) {
	workerPosition = GetPosition(index)

	workerInertia = GetInertia(index)

	workerPosition[0] = workerPosition[0] + (workerInertia[0] * dtime)
	//collision detect X goes here
	workerPosition[1] = workerPosition[1] + (workerInertia[1] * dtime)
	//collision detect Y goes here

	SetPosition(index, workerPosition[0], workerPosition[1])
}

func RunPhysics(delta float64) {

	for i := 0; i < GetNumberOfEntities(); i++ {
		if !GetStatic(i) {
			applyInertia(0, delta)
		}
	}

	//X = rand.Float64() * 300
	//Y = rand.Float64() * 300

	/*
		if right && X <= 100 {
			X++
			Y += 1.25

			if X > 100 {
				right = !right
			}
		} else if !right && X >= 0 {
			X--
			Y -= 1.2

			if X < 0 {
				right = !right
			}
		}
	*/
}
