package imageproc

type Point struct {
	x, y int
}

type ScreenSection struct {
	UpperLeft, LowerRight Point
}

//  Errors

// image processing and screen handling
type IncorrectBondaries struct{}

func (i * IncorrectBondaries) Error() string {
	return "Boundary must start from upper left corner and end in lower right corner."
}