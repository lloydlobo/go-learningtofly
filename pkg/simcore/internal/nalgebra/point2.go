package nalgebra

import (
	"math"
	"math/rand"
)

type Point2 struct {
	X float32
	Y float32
}

func NewPoint2(x, y float32) Point2               { return Point2{x, y} }
func (Point2) Random(rng *rand.Rand) Point2       { return Point2{rng.Float32(), rng.Float32()} }
func (p *Point2) Add(other *Point2) Point2        { return Point2{p.X + other.X, p.Y + other.Y} }
func (p *Point2) Sub(other *Point2) Point2        { return Point2{p.X - other.X, p.Y - other.Y} }
func (p *Point2) Dot(other *Point2) Point2        { return Point2{p.X * other.X, p.Y * other.Y} }
func (p *Point2) MulScalar(scalar float32) Point2 { return Point2{p.X * scalar, p.Y * scalar} }
func (p *Point2) Length() float32                 { return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y))) }

func (Point2) Distance(p1, p2 *Point2) float32 {
	dx := math.Abs(float64(p1.X) - float64(p2.X))
	dy := math.Abs(float64(p1.Y) - float64(p2.Y))
	return float32(math.Sqrt(dx + dy))
}
