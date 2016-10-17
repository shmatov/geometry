package geometry

type Segment struct {
	begin Point2D
	end   Point2D
}

func (segment Segment) DistanceToPoint(point Point2D) Distance {
	return segment.ComparableDistanceToPoint(point).ToDistance()
}

func (segment Segment) ComparableDistanceToPoint(point Point2D) ComparableDistance {
	x := segment.begin.X
	y := segment.begin.Y
	dx := segment.end.X - x
	dy := segment.end.Y - y

	if dx != 0 || dy != 0 {
		t := ((point.X-x)*dx + (point.Y-y)*dy) / (dx*dx + dy*dy)

		if t > 1 {
			x = segment.end.X
			y = segment.end.Y
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}

	dx = point.X - x
	dy = point.Y - y

	return CalculateComparableDistance(dx, dy)
}
