package main

func main() {

}

func isBoomerang(points [][]int) bool {
	x1, y1 := points[1][0]-points[0][0], points[1][1]-points[0][1]
	x2, y2 := points[2][0]-points[1][0], points[2][1]-points[1][1]
	return x1*y2 != x2*y1
}
