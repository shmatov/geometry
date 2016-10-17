package geometry

import "math"

type Path []Point2D

func (path Path) ComparableDistanceToPoint(point Point2D) ComparableDistance {
	minDistance := ComparableDistance(math.MaxFloat64)
	for i := 0; i < len(path)-1; i++ {
		distance := Segment{path[i], path[i+1]}.ComparableDistanceToPoint(point)
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func (path Path) DistanceToPoint(point Point2D) Distance {
	return path.ComparableDistanceToPoint(point).ToDistance()
}
