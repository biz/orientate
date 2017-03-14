// Package orientate is used to calculate how an image should be rotated and flipped
// given the EXIF Orintation value of the image
package orientate

import "fmt"

type Flip int

// Constants to indicate if the image needs to be flipped
const (
	NoFlip Flip = iota
	FlipHorizontal
	FlipVertical
)

type Rotate int

// Constants to indicate how the image needs to be rotated
const (
	NoRotate  Rotate = 0
	Rotate90  Rotate = 90
	Rotate180 Rotate = 180
	Rotate270 Rotate = 270
)

// Op represents the actions needed to correct the orientation of an image
type Op struct {
	Rotate Rotate
	Flip   Flip
}

// Orientations store what Op needs to be performed for a given operation
var Orientations = [...]Op{
	{NoRotate, NoFlip},
	{NoRotate, FlipHorizontal},
	{Rotate180, NoFlip},
	{NoRotate, FlipVertical},
	{Rotate270, FlipHorizontal},
	{Rotate270, NoFlip},
	{Rotate90, FlipHorizontal},
	{Rotate90, NoFlip},
}

func Calc(orientation int) (rotate Rotate, flip Flip, err error) {
	if orientation < 1 || orientation > 8 {
		return 0, 0, fmt.Errorf("orientation should be between 1 and 8")
	}
	orientation--

	return Orientations[orientation].Rotate, Orientations[orientation].Flip, nil
}
