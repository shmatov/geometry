package geometry

import "testing"

func TestSegmentDistanceToPoint(t *testing.T) {
	type pointWithDistance struct {
		point    Point2D
		distance Distance
	}
	testCases := []struct {
		segment            Segment
		pointsWithDistanes []pointWithDistance
	}{
		{
			// horizontal segment
			Segment{Point2D{0, 0}, Point2D{4, 0}},
			[]pointWithDistance{
				// on the segment
				{Point2D{0, 0}, 0},
				{Point2D{4, 0}, 0},
				{Point2D{2, 0}, 0},

				// above the segment
				{Point2D{0, 2}, 2},
				{Point2D{2, 2}, 2},
				{Point2D{4, 2}, 2},

				// below the segment
				{Point2D{0, -2}, 2},
				{Point2D{2, -2}, 2},
				{Point2D{4, -2}, 2},

				// on the same line
				{Point2D{-2, 0}, 2},
				{Point2D{6, 0}, 2},

				// above/below the line outside of the segment
				{Point2D{7, 4}, 5},
				{Point2D{7, -4}, 5},
				{Point2D{-4, 3}, 5},
				{Point2D{-4, -3}, 5},
			},
		},
		{
			// vertical segment
			Segment{Point2D{0, 0}, Point2D{0, 4}},
			[]pointWithDistance{
				// on the segment
				{Point2D{0, 0}, 0},
				{Point2D{0, 4}, 0},
				{Point2D{0, 2}, 0},

				// to the right of the segment
				{Point2D{2, 0}, 2},
				{Point2D{2, 2}, 2},
				{Point2D{2, 4}, 2},

				// to the left of the segment
				{Point2D{-2, 0}, 2},
				{Point2D{-2, 2}, 2},
				{Point2D{-2, 4}, 2},

				// on the same line
				{Point2D{0, -2}, 2},
				{Point2D{0, 6}, 2},

				// to the left/right of the line outside of the segment
				{Point2D{4, 7}, 5},
				{Point2D{-4, 7}, 5},
				{Point2D{3, -4}, 5},
				{Point2D{-3, -4}, 5},
			},
		},
		{
			// zero-sized segment
			Segment{Point2D{0, 0}, Point2D{0, 0}},
			[]pointWithDistance{
				{Point2D{0, 0}, 0},
				{Point2D{2, 0}, 2},
				{Point2D{-2, 0}, 2},
				{Point2D{0, 2}, 2},
				{Point2D{0, -2}, 2},

				{Point2D{3, -4}, 5},
				{Point2D{-4, 3}, 5},
			},
		},
	}

	for testCaseIndex, tc := range testCases {
		for pointIndex, pd := range tc.pointsWithDistanes {
			result := tc.segment.DistanceToPoint(pd.point)

			if result != pd.distance {
				t.Errorf("Test case #%v.%v. Segment: %v Point: %v Expected: %v Result: %v",
					testCaseIndex, pointIndex, tc.segment, pd.point, pd.distance, result)
			}
		}
	}
}
