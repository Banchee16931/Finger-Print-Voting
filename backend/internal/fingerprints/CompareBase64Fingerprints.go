package fingerprints

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

func CompareBase64Fingerprints(printOne, printTwo string) (bool, error) {
	firstImage, err := ConvertBase64ToImage(printOne)
	if err != nil {
		return false, err
	}

	secondImage, err := ConvertBase64ToImage(printTwo)
	if err != nil {
		return false, err
	}

	return Compare(firstImage, secondImage, 0.9), nil
}

func ConvertBase64ToImage(base string) (image.Image, error) {
	if len(base) < 6 {
		return image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}}), fmt.Errorf("invalid image string")
	}

	coI := strings.Index(base, ",")
	if coI == -1 {
		return image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}}), fmt.Errorf("invalid image string")
	}

	rawImage := base[coI+1:]

	unbased, _ := base64.StdEncoding.DecodeString(rawImage)

	res := bytes.NewReader(unbased)

	switch strings.TrimSuffix(base[5:coI], ";base64") {
	case "image/png":
		return png.Decode(res)
		// ...
	case "image/jpeg":
		return jpeg.Decode(res)
	}

	return image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}}), fmt.Errorf("image is invalid type")
}
