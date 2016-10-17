package geometry

func Simplify(points []Point2D, threshold Distance) []Point2D {
	mask, found := applyDouglasPeucker(points, threshold.ToComparableDistance())
	selectedPoints := make([]Point2D, 0, found)

	for i, selected := range mask {
		if selected {
			selectedPoints = append(selectedPoints, points[i])
		}
	}
	return selectedPoints
}

func applyDouglasPeucker(points []Point2D, threshold ComparableDistance) ([]bool, int) {
	mask := make([]bool, len(points))
	mask[0] = true
	mask[len(mask)-1] = true

	type interval struct {
		start int
		end   int
	}
	stack := []interval{{0, len(points) - 1}}
	found := 0

	for len(stack) > 0 {
		currentInterval := stack[len(stack)-1]
		stack = stack[:len(stack)-1] 

		segment := Segment{points[currentInterval.start], points[currentInterval.end]}

		maxDistance := ComparableDistance(-1.0)
		maxIndex := 0
		for i := currentInterval.start + 1; i < currentInterval.end; i++ {
			distance := segment.ComparableDistanceToPoint(points[i])

			if distance > maxDistance {
				maxDistance = distance
				maxIndex = i
			}
		}

		if maxDistance > threshold {
			found++
			mask[maxIndex] = true

			stack = append(
				stack,
				interval{currentInterval.start, maxIndex},
				interval{maxIndex, currentInterval.end},
			)
		}
	}

	return mask, found
}
