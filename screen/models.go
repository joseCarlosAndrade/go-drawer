package screen

type Point struct {
	X, Y int
}

type ScreenSection struct {
	UpperLeft, LowerRight Point
	Width, Height int
}

//  Errors

// image processing and screen handling
type IncorrectBondaries struct{}

type ImageIsNil struct {}

func (i * IncorrectBondaries) Error() string {
	return "Boundary must start from upper left corner and end in lower right corner."
}

func (e * ImageIsNil) Error() string {
	return "Image is nil"
}