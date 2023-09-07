package fingerprints

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

// Compares fingerprints are returns if the are the same
func Compare(databaseImage image.Image, queryImage image.Image, confidenceThreshold float64) bool {
	mse := calculateMSE(databaseImage, queryImage)
	return (1 - mse) > confidenceThreshold
}

// Calculates the confidence value of a comparison between two fingerprints
func calculateMSE(img1, img2 image.Image) float64 {
	bounds := img1.Bounds()
	mse := 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c1 := color.GrayModel.Convert(img1.At(x, y)).(color.Gray)
			c2 := color.GrayModel.Convert(img2.At(x, y)).(color.Gray)
			delta := float64(c1.Y) - float64(c2.Y)

			mse += math.Abs(delta)
		}
	}

	mse /= 255
	mse /= float64(bounds.Dx() * bounds.Dy())
	return mse
}
