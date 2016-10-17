package geometry

import (
	"math"
)

type BoundingBox struct {
	min Point2D
	max Point2D
}

func (bbox BoundingBox) Contains(point Point2D) bool {
	return (bbox.min.X <= point.X && point.X <= bbox.max.X &&
		bbox.min.Y <= point.Y && point.Y <= bbox.max.Y)
}

func (bbox BoundingBox) ComparableDistanceToPoint(point Point2D) ComparableDistance {
	if bbox.Contains(point) {
		return 0
	}
	return Path([]Point2D{
		bbox.min, bbox.max, {bbox.min.X, bbox.max.Y}, {bbox.max.X, bbox.min.Y},
	}).ComparableDistanceToPoint(point)
}

func (bbox BoundingBox) DistanceToPoint(point Point2D) Distance {
	return bbox.ComparableDistanceToPoint(point).ToDistance()
}

func FindBoundingBox(points []Point2D) BoundingBox {
	minX, minY := points[0].X, points[0].Y
	maxX, maxY := points[0].X, points[0].Y
	for _, point := range points {
		minX = math.Min(minX, point.X)
		minY = math.Min(minY, point.Y)
		maxX = math.Max(maxX, point.X)
		maxY = math.Max(maxY, point.Y)
	}
	return BoundingBox{
		min: Point2D{X: minX, Y: minY},
		max: Point2D{X: maxX, Y: maxY},
	}
}
