package particles

import (
	"image/color"
	"math/rand"
)

var colours = []color.Color{
	color.RGBA{R: 255, G: 255, B: 255, A: 255},
	color.RGBA{R: 255, G: 0, B: 0, A: 255},
	color.RGBA{R: 0, G: 255, B: 0, A: 255},
	color.RGBA{R: 0, G: 0, B: 255, A: 255},
	color.RGBA{R: 255, G: 255, B: 0, A: 255},
	color.RGBA{R: 255, G: 0, B: 255, A: 255},
	color.RGBA{R: 0, G: 255, B: 255, A: 255},
}

type Particle struct {
	Position Point64
	Velocity Vector64
	Mass     float64
	Radius   float64
	Colour   color.Color
}

func RandomParticles(width float64, height float64, count int) []Particle {
	particles := make([]Particle, count)

	for f := 0; f < count; f++ {
		particles[f] = RandomParticle(width, height)
	}

	return particles
}

func RandomParticle(width float64, height float64) Particle {
	return Particle{
		Position: Point64{
			X: rand.Float64() * width,
			Y: rand.Float64() * height,
		},
		Velocity: Vector64{
			X: 0,
			Y: 0,
		},
		Mass:   rand.Float64() * 5,
		Radius: 5 + rand.Float64()*5,
		Colour: colours[rand.Intn(len(colours))],
	}
}

func UpdatePositions(particles []Particle) {
	for f := 0; f < len(particles); f++ {
		for g := f + 1; g < len(particles); g++ {
			applyGravity(&particles[f], &particles[g])
		}
		particles[f].Position.Move(particles[f].Velocity)
	}
}

func applyGravity(x *Particle, y *Particle) {
	// Acceleration due to gravity is inversely proportional to the square of the distance between the particles.
	gOverDSquared := x.Mass * y.Mass / SeparationSquared(x.Position, y.Position)

	xChange := (y.Position.X - x.Position.X) * gOverDSquared

	x.Velocity.X += xChange
	y.Velocity.X -= xChange

	yChange := (y.Position.Y - x.Position.Y) * gOverDSquared
	x.Velocity.Y += yChange
	y.Velocity.Y -= yChange
}
