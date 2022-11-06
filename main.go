package main

import (
	"fmt"
	"os"
	"os/exec"
	"particles-go/drawing"
	p "particles-go/particles"
	"path/filepath"
)

var frames = 60 * 10
var width = 1080.0
var height = 1080.0

func main() {
	outputDirectory := getOutputDirectory()

	renderFrames(outputDirectory)

	generateMovie(outputDirectory)

	deleteFrames(outputDirectory)
}

func generateMovie(outputDirectory string) {
	files := filepath.Join(outputDirectory, "Image%4d.png")
	output := filepath.Join(outputDirectory, "Movie.mp4")
	ffmpeg := exec.Command(
		"ffmpeg",
		"-framerate", "60",
		// Must come before the input files as it specifies the rate for them.
		"-r", "60",
		"-i", files,
		"-c:v", "h264",
		"-c:a", "aac",
		"-preset", "slow",
		"-pix_fmt", "yuv420p",
		"-crf", "23",
		"-b:v", "3500k",
		"-b:a", "256k",
		"-ar", "44100",
		"-f", "mp4",
		"-y",
		output)

	fmt.Printf("Running command %v", ffmpeg)

	err := ffmpeg.Run()
	if err != nil {
		panic(fmt.Sprintf("Error running command: %v", ffmpeg))
	}
}

func renderFrames(outputDirectory string) {
	particles := p.RandomParticles(width, height, 80)

	for f := 0; f < frames; f++ {
		filename := fmt.Sprintf("Image%04d.png", f)
		drawing.RenderParticles(filepath.Join(outputDirectory, filename), width, height, particles)

		p.UpdatePositions(particles)
	}
}

func deleteFrames(outputDirectory string) {
	for f := 0; f < frames; f++ {
		path := filepath.Join(outputDirectory, fmt.Sprintf("Image%04d.png", f))
		os.Remove(path)
	}
}

func getOutputDirectory() string {
	if len(os.Args) != 2 {
		panic("No output directory specified.")
	}

	outputDirectory := os.Args[1]

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		panic(fmt.Sprintf("The output directory %v does not exist.", outputDirectory))
	}

	return outputDirectory
}
