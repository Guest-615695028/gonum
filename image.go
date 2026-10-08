package gonum

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

const (
	cameraRadius     = 0.0254 / 6
	pixelWidth       = 2e-6
	pixelLength      = 2e-6
	lensDistance     = 0.0325
	SignalNoiseRatio = 6760.83
)

func OpenImage(name string) (image.Image, error) {
	f, err := os.Open(name)
	if f == nil {
		return nil, err
	}
	defer f.Close()
	defer func() { recover() }()
	i := len(name)
	for i > 0 {
		if i--; name[i] == '.' {
			break
		}
	}
	switch name[i+1] {
	case 'P', 'p':
		return png.Decode(f)
	case 'G', 'g':
		return gif.Decode(f)
	case 'J', 'j':
		return jpeg.Decode(f)
	}
	m, _, err := image.Decode(f)
	return m, err
}

func OpenRGBA(name string) (*image.RGBA, error) {
	img, err := OpenImage(name)
	return img.(*image.RGBA), err
}

func WritePNG(name string, im image.Image) error {
	f, err := os.Create(name + ".png")
	if f == nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, im)
}
func WriteJPEG(name string, im image.Image, q int) error {
	f, err := os.Create(name + ".jpeg")
	if f == nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, im, &jpeg.Options{
		Quality: q,
	})
}
func WriteGIF(name string, im image.Image, n int) error {
	f, err := os.Create(name + ".gif")
	if f == nil {
		return err
	}
	defer f.Close()
	return gif.Encode(f, im, &gif.Options{
		NumColors: n,
	})
}
