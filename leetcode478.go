package main

import "math/rand"

func main() {

}

type Solution struct {
	radius float64
	x      float64
	y      float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius: radius,
		x:      x_center,
		y:      y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	dx := rand.Float64()
	dy := rand.Float64()
	if dx*dx+dy*dy > 1 {
		return this.RandPoint()
	}
	ne := rand.Intn(2)
	if ne == 0 {
		dx = -dx
	}
	ne = rand.Intn(2)
	if ne == 0 {
		dy = -dy
	}
	return []float64{this.x + dx*this.radius, this.y + dy*this.radius}

}
