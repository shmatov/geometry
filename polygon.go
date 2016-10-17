package geometry

import "math"

type Polygon interface {
	Contains(Point2D) bool
	BBox() BoundingBox
	ComparableDistanceToPoint(Point2D) ComparableDistance
	DistanceToPoint(Point2D) Distance
}

func NewPolygon(shell Path, holes ...Path) simplePolygon {
	return simplePolygon{shell: shell, holes: holes, bbox: FindBoundingBox(shell)}
}

type simplePolygon struct {
	shell Path
	holes []Path
	bbox  BoundingBox
}

func (p simplePolygon) Contains(point Point2D) bool {
	if !p.BBox().Contains(point) {
		return false
	}
	if RayCasting(point, p.shell) != Inside {
		return false
	}
	for _, hole := range p.holes {
		if RayCasting(point, hole) == Inside {
			return false
		}
	}
	return true
}

func (p simplePolygon) BBox() BoundingBox {
	return p.bbox
}

func (p simplePolygon) Prepare() preparedPolygon {
	return preparedPolygon{bbox: p.bbox}
}

type preparedPolygon struct {
	bbox BoundingBox
}

func (polygon simplePolygon) DistanceToPoint(point Point2D) Distance {
	return polygon.ComparableDistanceToPoint(point).ToDistance()
}

func (polygon simplePolygon) ComparableDistanceToPoint(point Point2D) ComparableDistance {
	if polygon.Contains(point) {
		return 0
	}

	minDistance := Path(polygon.shell).ComparableDistanceToPoint(point)
	for _, hole := range polygon.holes {
		distance := Path(hole).ComparableDistanceToPoint(point)
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func FindNearestPolygon(point Point2D, polygons []Polygon) int {
	minDistance := ComparableDistance(math.MaxFloat64)
	closestPolygonIndex := -1
	for index, polygon := range polygons {
		if polygon.BBox().ComparableDistanceToPoint(point) > minDistance {
			continue
		}
		distance := polygon.ComparableDistanceToPoint(point)
		if distance < minDistance {
			minDistance = distance
			closestPolygonIndex = index
		}
	}
	return closestPolygonIndex
}
