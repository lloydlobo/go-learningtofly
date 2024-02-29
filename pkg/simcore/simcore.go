// Pacage simcore implements a simulation engine.
//
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
package simcore

import (
	"math"
	"math/rand"

	"simcore/internal/nalgebra"
)

type Simulation struct {
	World World // Our world is two-dimensional...
}

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

// Step performs a single step - a single second, of simulation.
//
// Reference: https://pwy.io/posts/learning-to-fly-pt4/#stepping-stones
func (s *Simulation) Step(rng *rand.Rand) {
	processCollisions(s, rng)
	processMovements(s)
}

// hit-testing:
//
// process of checking whether two polygons collide:
//   - Birds are triangles, Foods are circles
//   - but triangle circle collision hit-testing is complex.
//
// circle-circle hit-testing:
//
// We can keep drawing birds as tringles, but assume them as circles.
// So circle-circle hit testing relies on checking if distance between two
// circles is shorter or equal than the sum of their radii.
//
//	distance(A,B) >  radius(A) + radius(B) -> no collison
//	distance(A,B) <= radius(A) + radius(B) -> no collison
//
// A distance of 0.5 -> animal and food are half a map apart from each other,
// while distance of 0.0 -> animal and food are at exact same coordinates.
//
// Reference: https://pwy.io/posts/learning-to-fly-pt4/#ur-somebody-else
func processCollisions(s *Simulation, rng *rand.Rand) {
	const foodRadius float32 = 0.01

	for _, animal := range s.World.Animals {
		for j, food := range s.World.Foods {
			distance := (nalgebra.Point2{}).Distance(&animal.Position, &food.Position)

			if animalCollidesWithFood := distance <= foodRadius; animalCollidesWithFood {
				s.World.Foods[j].Position = (nalgebra.Point2{}).Random(rng)
			}
		}
	}
}

func processMovements(s *Simulation) {
	const (
		min = 0.0
		max = 1.0
		// Our map is bounded by <0.0,1.0>, anything beyond those coordinates
		// can exist, but rendered outside canvas.
	)

	for i, animal := range s.World.Animals {
		rotation, speed := float64(animal.Rotation), float64(animal.Speed)
		dx, dy := math.Cos(rotation)*speed, math.Sin(rotation)*speed

		s.World.Animals[i].Position.X += float32(dx)
		s.World.Animals[i].Position.Y += float32(dy)

		s.World.Animals[i].Position.X = wrapFloat(s.World.Animals[i].Position.X, min, max)
		s.World.Animals[i].Position.Y = wrapFloat(s.World.Animals[i].Position.Y, min, max)
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

// (if num < min: return max) or (if num > max: return min)
func wrapFloat[T float32 | float64](num, min, max T) T {
	if num < min {
		return max
	} else if num > max {
		return min
	}
	return num
}

// ^
// | Faster inline implementation:
// |
// |	x := s.World.Animals[i].Position.X
// |	y := s.World.Animals[i].Position.Y
// |	if x < min {
// |		s.World.Animals[i].Position.X = max
// |	} else if x > max {
// |		s.World.Animals[i].Position.X = min
// |	}
// |	if y < min {
// |		s.World.Animals[i].Position.Y = max
// |	} else if y > max {
// |		s.World.Animals[i].Position.Y = min
// |	}
// |	func wrap(num, min, max float32) float32 {
// |		if num < min {
// |			return max
// |		} else if num > min {
// |			return min
// |		}
// |		return num
// |	}
// |
// | References:
// |   - generics: https://stackoverflow.com/a/70370013
// |   - constraints: https://go.dev/ref/spec#Type_constraints
// |
// | 	func wrap[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64]( num, min, max T,) T {
// ---
