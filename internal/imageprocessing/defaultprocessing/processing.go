package defaultprocessing

import (
	"context"
	"image"
	"strings"

	"github.com/Falokut/image_processing_service/internal/imageprocessing"
	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/transform"
)

type ImageProcessing struct {
}

func (p ImageProcessing) Resize(_ context.Context, img image.Image, width, height int, resampleMethod string) image.Image {
	return transform.Resize(img, width, height, resolveFiter(resampleMethod))
}

func (p ImageProcessing) Crop(_ context.Context, img image.Image, x0, y0, x1, y1 int) image.Image {
	return transform.Crop(img, image.Rectangle{
		Min: image.Point{X: x0, Y: y0},
		Max: image.Point{X: x1, Y: y1}})
}

func (p ImageProcessing) Desaturate(_ context.Context, img image.Image) image.Image {
	return effect.Grayscale(img)
}

func (p ImageProcessing) Hue(_ context.Context, img image.Image, hue int) image.Image {
	return adjust.Hue(img, hue%imageprocessing.HueMaxValue)
}

type BlurMethod string

const (
	Box      BlurMethod = "Box"
	Gaussian BlurMethod = "Gaussian"
)

func (p ImageProcessing) Blur(_ context.Context, img image.Image, radius float64, method string) image.Image {
	blurMethod := resolveBlurMethod(method)
	switch blurMethod {
	case Gaussian:
		return blur.Gaussian(img, radius)
	default:
		return blur.Box(img, radius)
	}
}

func resolveFiter(filter string) transform.ResampleFilter {
	switch filter {
	case "Lanczos":
		return transform.Lanczos
	case "CatmullRom":
		return transform.CatmullRom
	case "MitchellNetravali":
		return transform.MitchellNetravali
	case "Linear":
		return transform.Linear
	case "Box":
		return transform.Box
	case "NearestNeighbor":
		return transform.NearestNeighbor
	default:
		return transform.NearestNeighbor
	}
}

func resolveBlurMethod(method string) BlurMethod {
	method = strings.ToUpper(method)
	switch method {
	case "BOX":
		return Box
	case "GAUSSIAN":
		return Gaussian
	default:
		return Box
	}
}
