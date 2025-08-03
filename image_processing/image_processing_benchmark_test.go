package imageprocessing

import (
	"image"
	"testing"
)

func BenchmarkGrayscale(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000)) // simulate large image

	for i := 0; i < b.N; i++ {
		_ = Grayscale(img)
	}
}
