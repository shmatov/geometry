package geometry

import (
	"testing"
)

func TestRightRayIntersection(t *testing.T) {
	testCases := []struct {
		origin   Point2D
		begin    Point2D
		end      Point2D
		expected Position
	}{
		// If one of the segment's points equals to origin or
		// segment contains origin, then result should be equal 0.
		{Point2D{0, 0}, Point2D{0, 0}, Point2D{1, -1}, Boundary},
		{Point2D{0, 0}, Point2D{0, 0}, Point2D{-1, 1}, Boundary},
		{Point2D{0, 0}, Point2D{2, -1}, Point2D{0, 0}, Boundary},
		{Point2D{0, 0}, Point2D{-1, 1}, Point2D{0, 0}, Boundary},

		{Point2D{0, 0}, Point2D{0, 1}, Point2D{0, -1}, Boundary},
		{Point2D{0, 0}, Point2D{1, 1}, Point2D{-1, -1}, Boundary},

		// If right-ray from the origin doesn't intersect segment or
		// segment lies on right-ray or segment touches right-ray from the top,
		// then result should be equal 1.
		{Point2D{0, 0}, Point2D{-1, 1}, Point2D{-1, 1}, Inside},
		{Point2D{0, 0}, Point2D{1, 1}, Point2D{1, 2}, Inside},
		{Point2D{0, 0}, Point2D{1, -1}, Point2D{1, -2}, Inside},

		{Point2D{0, 0}, Point2D{1, 0}, Point2D{2, 0}, Inside},
		{Point2D{0, 0}, Point2D{5, 0}, Point2D{2, 0}, Inside},

		{Point2D{0, 0}, Point2D{1, 1}, Point2D{2, 0}, Inside},
		{Point2D{0, 0}, Point2D{2, 1}, Point2D{2, 0}, Inside},
		{Point2D{0, 0}, Point2D{3, 1}, Point2D{2, 0}, Inside},

		// If right-ray from the origin intersects segment or segment touches
		// right-ray from the bottom, then result should be equal -1.
		{Point2D{0, 0}, Point2D{1, 1}, Point2D{1, -1}, Outside},
		{Point2D{0, 0}, Point2D{1, -1}, Point2D{1, 1}, Outside},
		{Point2D{0, 0}, Point2D{1, -1}, Point2D{1, 1}, Outside},
		{Point2D{0, 0}, Point2D{-100, -200}, Point2D{500, 1}, Outside},

		{Point2D{0, 0}, Point2D{1, -1}, Point2D{2, 0}, Outside},
		{Point2D{0, 0}, Point2D{2, -1}, Point2D{2, 0}, Outside},
		{Point2D{0, 0}, Point2D{3, -1}, Point2D{2, 0}, Outside},
	}
	for i, tc := range testCases {
		result := RightRayIntersection(tc.origin, Segment{tc.begin, tc.end})
		if result != tc.expected {
			t.Errorf("Test case #%v. Origin: %v Segment: %v %v Expected: %v Result: %v",
				i, tc.origin, tc.begin, tc.end, tc.expected, result)
		}
	}
}

func BenchmarkRightRayIntersection(b *testing.B) {
	benches := []struct {
		begin Point2D
		end   Point2D
		name  string
	}{
		{Point2D{-1, -1}, Point2D{-1, 1}, "outside"},
		{Point2D{0, -1}, Point2D{0, 1}, "touches"},
		{Point2D{1, 0}, Point2D{1, 1}, "above"},
		{Point2D{1, 0}, Point2D{1, -1}, "below"},
		{Point2D{1, -1}, Point2D{1, 1}, "intersects"},
	}

	origin := Point2D{0, 0}
	for _, bench := range benches {
		segment := Segment{bench.begin, bench.end}

		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RightRayIntersection(origin, segment)
			}
		})
	}
}
