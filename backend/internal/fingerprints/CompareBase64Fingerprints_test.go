package fingerprints_test

import (
	"finger-print-voting-backend/internal/fingerprints"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Checks that ConvertBase64ToImage can successfully convert a base64 encoded image to a image.Image
func TestConvertBase64ToImage(t *testing.T) {
	_, err := fingerprints.ConvertBase64ToImage(ImageOne)
	assert.NoError(t, err, "ConvertBase64ToImage returned an error")
}

// Checks that ConvertBase64ToImage successfully reports errors
func TestConvertBase64ToImage_Error(t *testing.T) {
	_, err := fingerprints.ConvertBase64ToImage("")
	assert.Error(t, err, "ConvertBase64ToImage returned an error")
}

// Checks that CompareBase64Fingerprints can identify when two fingerprints are the same
func TestCompareBase64Fingerprints_Same(t *testing.T) {
	ret, err := fingerprints.CompareBase64Fingerprints(ImageOne, ImageOne)
	assert.NoError(t, err, "CompareBase64Fingerprints returned an error")
	assert.True(t, ret, "returned false when the fingerprints are the same")
}

// Checks that CompareBase64Fingerprints can identify when two fingerprints are the different
func TestCompareBase64Fingerprints_Different(t *testing.T) {
	ret, err := fingerprints.CompareBase64Fingerprints(ImageOne, ImageTwo)
	assert.NoError(t, err, "CompareBase64Fingerprints returned an error")
	assert.False(t, ret, "returned true when the fingerprints are different")
}

// Checks that CompareBase64Fingerprints successfully reports errors
func TestCompareBase64Fingerprints_Errors(t *testing.T) {
	cases := []struct {
		name     string
		imageOne string
		imageTwo string
	}{
		{
			name:     "image_one_error",
			imageOne: "not an image hahaha",
			imageTwo: ImageOne,
		},

		{
			name:     "image_two_error",
			imageOne: ImageOne,
			imageTwo: "not an image hahaha",
		},
		{
			name:     "invalid_jpeg",
			imageOne: "invalid image thing, after part",
			imageTwo: ImageOne,
		},
		{
			name:     "small_length_image",
			imageOne: "",
			imageTwo: ImageOne,
		},
	}

	for i := 0; i < len(cases); i++ {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			ret, err := fingerprints.CompareBase64Fingerprints("invalid image thing, after part", ImageOne)
			assert.Error(t, err, "ConvertBase64ToImage returned an error")
			assert.False(t, ret, "returned true when an error occured")
		})
	}

}
