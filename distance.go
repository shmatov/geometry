package geometry

import "math"

type Distance float64
type ComparableDistance float64

func CalculateComparableDistance(x, y float64) ComparableDistance {
	return ComparableDistance(x*x + y*y)
}

func (distance Distance) ToComparableDistance() ComparableDistance {
	return ComparableDistance(distance * distance)
}

func (comparableDistance ComparableDistance) ToDistance() Distance {
	return Distance(math.Sqrt(float64(comparableDistance)))
} 
