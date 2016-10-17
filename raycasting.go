package geometry

type Position int8

const (
	Inside   Position = 1
	Boundary Position = 0
	Outside  Position = -1
)

func RayCasting(origin Point2D, polygon []Point2D) Position {
	position := Outside
	for i := 0; i < len(polygon)-1; i++ {
		position *= RightRayIntersection(origin, Segment{polygon[i], polygon[i+1]})
	}
	return position
}

func RightRayIntersection(origin Point2D, segment Segment) Position {
	begin := segment.begin
	end := segment.end

	ax := begin.X - origin.X
	ay := begin.Y - origin.Y

	bx := end.X - origin.X
	by := end.Y - origin.Y

	sv := ax*by - ay*bx
	if sv == 0 && (ay == 0 || by == 0) && ax*bx <= 0 {
		return 0
	}
	if (ay < 0) != (by < 0) {
		if sv == 0 {
			return 0
		}
		if by < 0 {
			return sign(sv)
		}
		return -sign(sv)
	}
	return 1
}

func sign(num float64) Position {
	if num < 0 {
		return -1
	}
	return 1
}
