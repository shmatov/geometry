package geometry

import "testing"

func TestSimplify(t *testing.T) {
	testCases := []struct {
		path      []Point2D
		threshold Distance
		result    []Point2D
	}{
		{
			[]Point2D{{0, 0}, {1, 1}, {2, 0}}, 0,
			[]Point2D{{0, 0}, {1, 1}, {2, 0}},
		},
		{
			[]Point2D{{0, 0}, {1, 1}, {2, 0}}, 1,
			[]Point2D{{0, 0}, {2, 0}},
		},
		{
			[]Point2D{{0, 0}, {1, 2}, {2, 0}}, 1,
			[]Point2D{{0, 0}, {1, 2}, {2, 0}},
		},
		{
			[]Point2D{{0, 0}, {1, 3}, {2, 6}, {3, 3}, {4, 0}}, 1,
			[]Point2D{{0, 0}, {2, 6}, {4, 0}},
		},
		{
			[]Point2D{{0, 0}}, 1,
			[]Point2D{{0, 0}},
		},
		{
			[]Point2D{{0, 0}, {0, 0}}, 1,
			[]Point2D{{0, 0}, {0, 0}},
		},
	}

	for i, tc := range testCases {
		simplified := Simplify(tc.path, tc.threshold)
		if !isEquals(tc.result, simplified) {
			t.Errorf("Test case #%v. Expected: %v Simplified: %v",
				i, tc.result, simplified)
		}
	}
}

func isEquals(a []Point2D, b []Point2D) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
