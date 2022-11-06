package drawing

import (
	"fmt"
	"github.com/fogleman/gg"
	"particles-go/particles"
)

func RenderParticles(filename string, width float64, height float64, particles []particles.Particle) {
	context := gg.NewContext(int(width), int(height))
	context.DrawRectangle(0, 0, width, height)
	context.SetRGB(0, 0, 0)
	context.Fill()

	DrawParticles(context, particles)

	err := context.SavePNG(filename)
	if err != nil {
		panic(fmt.Sprintf("Error saving image: %v", err))
	}
}

func DrawParticles(context *gg.Context, particles []particles.Particle) {
	for _, particle := range particles {
		DrawParticle(context, particle)
	}
}

func DrawParticle(context *gg.Context, particle particles.Particle) {
	alphaStep := 1.0 / (float32(particles.TrailSize) + 2.0)
	alpha := alphaStep
	r := float32(particle.Colour.R)
	g := float32(particle.Colour.G)
	b := float32(particle.Colour.B)

	particle.Trail.DoReverse(func(position *particles.Point64) {
		context.DrawCircle(position.X, position.Y, particle.Radius)
		context.SetRGBA255(int(r*alpha), int(g*alpha), int(b*alpha), 255)
		context.Fill()
		alpha += alphaStep
	})

	context.DrawCircle(particle.Position.X, particle.Position.Y, particle.Radius)
	context.SetColor(particle.Colour)
	context.Fill()
}
