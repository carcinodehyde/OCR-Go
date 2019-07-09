package helpers

import (
	"errors"
	"image"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
)

func EnhanceKTP(dst string, img image.Image) error {

	rg := adjust.Brightness(img, 0.5)
	rg = adjust.Contrast(rg, 0.5)
	res := effect.Grayscale(rg)

	if err := imgio.Save(dst+"_enhanced.jpg", res, imgio.JPEGEncoder(95)); err != nil {
		return errors.New("save_failed")
	}

	return nil
}

func EnhanceGelap(dst string, img image.Image) error {

	rg := adjust.Brightness(img, -0.5)
	rg = adjust.Contrast(rg, 0.1)

	if err := imgio.Save(dst+"_enhanced.jpg", rg, imgio.JPEGEncoder(95)); err != nil {
		return errors.New("save_failed")
	}

	return nil
}

func EnhanceTerang(dst string, img image.Image) error {

	rg := adjust.Brightness(img, 0.5)
	rg = adjust.Contrast(rg, -0.1)

	if err := imgio.Save(dst+"_enhanced.jpg", rg, imgio.JPEGEncoder(95)); err != nil {
		return errors.New("save_failed")
	}

	return nil
}
