package imageprocessing

import (
	"image"
	"testing"
)

func TestGrayscale(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	result := Grayscale(img)
	if result.Bounds() != img.Bounds() {
		t.Errorf("Expected bounds %v but got %v", img.Bounds(), result.Bounds())
	}
}
