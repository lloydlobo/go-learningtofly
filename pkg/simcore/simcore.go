// Pacage simcore implements a simulation engine.
package simcore

import (
	"math"
	"math/rand"
)

type Simulation struct {
	World World // Our world is two-dimensional...
}

// Usage in JavaScript:
//
//	const world = simulation.world();
//	--------------------- ^---^
//	| Parsing already happens inside this automatically-generated
//	| function - we do not have to do anything more in here.
//	---
//
// returns a JSON:
//
//	{
//	  "animals": [
//	    { "x": 0.2, "y": 0.1 },
//	    { "x": 0.3, "y": 0.7 }
//	  ]
//	}
func (s *Simulation) Random(rng *rand.Rand) Simulation {
	world := s.World.Random(rng)

	// world=jsValueFromSerDe(world)
	// ^
	// | TODO
	// | What happened is that instead of letting wasm-pack serialize our models
	// | using its own serialization algorithm, we've forced them to be serialized
	// | into JSON — via JsValue.

	return Simulation{world}
}

// GetWorld implements a getter to access World from Simulation object's state.
func (s *Simulation) GetWorld(rng *rand.Rand) *World {
	return &s.World
}

// (if num < min: return max) or (if num > max: return min)
func wrap(num, min, max float32) float32 {
	if num < min {
		return max
	} else if num > min {
		return min
	}

	return num
	// | wrap() Faster inline usage:
	// |     x := s.World.Animals[i].Position.X
	// |     y := s.World.Animals[i].Position.Y
	// |     if x < min {
	// |     	s.World.Animals[i].Position.X = max
	// |     } else if x > max {
	// |     	s.World.Animals[i].Position.X = min
	// |     }
	// |     if y < min {
	// |     	s.World.Animals[i].Position.Y = max
	// |     } else if y > max {
	// |     	s.World.Animals[i].Position.Y = min
	// |     }
}

// Step performs a single step - a single second, of simulation.
//
// Reference: https://pwy.io/posts/learning-to-fly-pt4/#stepping-stones
func (s *Simulation) Step() {
	// Our map is bounded by <0.0,1.0>, anything beyond those coordinates
	// can exist, but rendered outside canvas.
	const (
		min = 0.0
		max = 1.0
	)

	for i, animal := range s.World.Animals {
		rotation, speed := float64(animal.Rotation), float64(animal.Speed)
		dx, dy := math.Cos(rotation)*speed, math.Sin(rotation)*speed

		s.World.Animals[i].Position.X += float32(dx)
		s.World.Animals[i].Position.Y += float32(dy)

		s.World.Animals[i].Position.X = wrap(s.World.Animals[i].Position.X, min, max)
		s.World.Animals[i].Position.Y = wrap(s.World.Animals[i].Position.Y, min, max)
	}
	// ^
	// | 	animal.Position += animal.Rotation * animal.Speed
	// | fix with nabgebra:    ^----------------------------^
	// |	animal.position += animal.rotation * na::Vector2::new(0.0, animal.speed);
	// |
	// | `::new(0.0, animal.speed)` says that: We're interested in rotating
	// | relative to the Y axis, that is: a bird with rotation of 0° will fly upwards.
	// | This decision neatly aligns with how we render triangles on <canvas>;
	// | we might've as well done e.g. ::new(-animal.speed, 0.0) and adjust our
	// | drawTriangle() to account for that.
}

// wasm-bindgen doesn't currently support exporting vectors of custom types.
//
// Even if it did, it's essential to maintain separation of concerns.
// lib-simulation (simcore) should focus on simulating evolution, not
// integrating with WebAssembly.
//
// By keeping lib-simulation(simcore) frontend-agnostic, it's easier to create
// different frontends such as lib-simulation-bevy or lib-simulation-cli
// sharing the same simulation code.
//
// Example:
//
//	#[wasm_bindgen]
//	#[derive(Debug)]
//	struct World {
//	    // Ok:
//	    animals: Vec<Animal>,
//
//	    // Error:
//	    pub foods: Vec<Food>,
//	}

type World struct {
	Animals []Animal
	Foods   []Food
}

func (w *World) Random(rng *rand.Rand) World {
	const (
		animalCount = 40
		foodCount   = 60
	)
	var (
		animal Animal
		food   Food
	)

	animals := make([]Animal, animalCount)
	for i := range animalCount {
		animals[i] = animal.Random(rng)
	}

	foods := make([]Food, foodCount)
	for i := range foodCount {
		foods[i] = food.Random(rng)
	}

	// ^ Our algorithm allows for animals and foods to overlap, so
	// | it's hardly ideal - but good enough for our purposes.
	// |
	// | A more complex solution could be based off of e.g.
	// | Poisson disk sampling:
	// |
	// | https://en.wikipedia.org/wiki/Supersampling
	// ---

	return World{animals, foods}
}

// GetAnimals implements a getter to access Animals from World object's state.
func (w *World) GetAnimals() *[]Animal {
	return &w.Animals
}

// GetFoods implements a getter to access Foods from World object's state.
func (w *World) GetFoods() *[]Food {
	return &w.Foods
}

type Animal struct {
	Position Point2
	Rotation Rotation2
	Speed    float32
}

func (a *Animal) Random(rng *rand.Rand) Animal {
	const speed = 0.002
	return Animal{
		Point2{}.Random(rng),
		NewRotation2(rng.Float32()),
		speed,
	}
}

// GetPosition implements a getter to access Position from Animal object's state.
func (a *Animal) GetPosition() Point2 {
	return a.Position
}

// GetRotation implements a getter to access Rotation from Animal object's state.
func (a *Animal) GetRotation() Rotation2 {
	return a.Rotation
}

type Food struct {
	Position Point2
}

func (f *Food) Random(rng *rand.Rand) Food {
	return Food{Point2{}.Random(rng)}
}

// GetPosition implements a getter to access Position from Food object's state.
func (f *Food) GetPosition() Point2 {
	return f.Position
}

type Rotation2 float32

func NewRotation2(angle float32) Rotation2 {
	// Normalize angle to be within [0, 2π)
	angle = float32(math.Mod(float64(angle), 2*math.Pi))
	if angle < 0.0 {
		angle += 2 * math.Pi
	}

	return Rotation2(angle)
}

func (r Rotation2) Angle() float32 {
	// Note: Is Rotation2 correctly implemented?
	// | Is Rotation2 a struct or just a type like enum?
	// |---v--------v
	return float32(r)
}

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
