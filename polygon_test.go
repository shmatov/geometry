package geometry

import (
	"testing"
)

func TestPolygonContains(t *testing.T) {
	testCases := []struct {
		point    Point2D
		expected bool
	}{
		{Point2D{0, 0}, false},
		{Point2D{0, 3}, false},
		{Point2D{3, 0}, false},
		{Point2D{3, 3}, false},

		{Point2D{100, 100}, false},
		{Point2D{-10, -10}, false},

		{Point2D{1.1, 1.1}, false},
		{Point2D{1.9, 1.9}, false},
		{Point2D{1.1, 1.9}, false},
		{Point2D{1.9, 1.1}, false},

		{Point2D{1, 1}, true},
		{Point2D{2, 1}, true},
		{Point2D{2, 2}, true},
		{Point2D{1, 2}, true},

		{Point2D{0.5, 0.5}, true},
		{Point2D{2.5, 0.5}, true},
		{Point2D{0.5, 2.5}, true},
		{Point2D{2.5, 2.5}, true},
	}

	polygon := NewPolygon(
		[]Point2D{{0, 0}, {0, 3}, {3, 3}, {3, 0}, {0, 0}},
		[]Point2D{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}},
	)

	for _, testCase := range testCases {
		if result := polygon.Contains(testCase.point); result != testCase.expected {
			t.Errorf("Point: %v Expected: %v Result: %v",
				testCase.point, testCase.expected, result)
		}
	}
}
