package geometry

import (
	"math/rand"
	"testing"
)

func TestFindBoundingBox(t *testing.T) {
	points := []Point2D{{0, 0}, {0, 3}, {3, 3}, {3, 0}, {0, 0}}
	expected := BoundingBox{min: Point2D{0, 0}, max: Point2D{3, 3}}
	if found := FindBoundingBox(points); found != expected {
		t.Errorf("expected: %#v, found: %#v", expected, found)
	}
}

func BenchmarkFindBoundingBox1000Points(b *testing.B) {
	points := generateRandomPoints(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindBoundingBox(points)
	}
}

func generateRandomPoints(count int) []Point2D {
	points := make([]Point2D, count)
	r := rand.New(rand.NewSource(0))

	for i := 0; i < len(points); i++ {
		points[i] = Point2D{
			r.Float64()*360 - 180,
			r.Float64()*360 - 180,
		}
	}

	return points
}
