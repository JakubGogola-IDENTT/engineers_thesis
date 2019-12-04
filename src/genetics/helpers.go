package genetics

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

// GetRandomColor returns random color in RGBA (without transparency)
func GetRandomColor(isGrayScale bool) color.Color {
	if isGrayScale {
		c := uint8(rand.Intn(256))

		return color.RGBA{R: c, G: c, B: c, A: 0}
	}

	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	return color.RGBA{R: r, G: g, B: b, A: 0}
}

// CompareColors compares color of two pixels
func CompareColors(c1, c2 color.Color, isGrayScale bool) (diff float64) {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	if isGrayScale {
		return math.Abs(float64(r1 - r2))
	}

	diff += math.Abs(float64(r1 - r2))
	diff += math.Abs(float64(g1 - g2))
	diff += math.Abs(float64(b1 - b2))
	diff += math.Abs(float64(a1 - a2))

	return diff
}

// MixColors returns average of two colors
func MixColors(c1, c2 color.Color) (c color.RGBA) {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	// Take average of every channel
	c.R = uint8(r1+r2) >> 1
	c.G = uint8(g1+g2) >> 1
	c.B = uint8(b1+b2) >> 1

	// c.A = uint8(a1+a2) >> 1
	c.A = 255

	return c
}

// GetRandomRectBounds returns random bounds of rectangle
func GetRandomRectBounds(rect image.Rectangle) image.Rectangle {
	// x and y of left upper corner of mutated rectangle
	minPt := image.Pt(rand.Intn(rect.Max.X), rand.Intn(rect.Max.Y))

	// width and height of mutated rectangle
	maxPt := image.Pt(minPt.X+rand.Intn(rect.Max.X-minPt.X), minPt.Y+rand.Intn(rect.Max.Y-minPt.Y))

	return image.Rectangle{minPt, maxPt}
}

// CopySpecimen - makes deep copy of specimen
func CopySpecimen(spec *Specimen, specToCopy Specimen) {
	img := *image.NewRGBA(spec.Spec.Bounds())

	copy(img.Pix, specToCopy.Spec.Pix)

	spec.Spec = img
}

// CrossSpecimens - crosses two specimens
func CrossSpecimens(spec1, spec2 *Specimen) {
	rect := GetRandomRectBounds(spec1.Spec.Bounds())

	minPt := rect.Bounds().Min
	maxPt := rect.Bounds().Max

	for x := minPt.X; x < maxPt.X; x++ {
		for y := minPt.Y; y < maxPt.Y; y++ {
			tmp := spec1.Spec.At(x, y)
			spec1.Spec.Set(x, y, spec2.Spec.At(x, y))
			spec2.Spec.Set(x, y, tmp)
		}
	}
}
