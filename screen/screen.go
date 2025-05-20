package screen

// this pkg will handle screen interaction (mouse, screenshots, etc)

import (
	"image"

	"github.com/go-vgo/robotgo"
	"go.uber.org/zap"
)



func GetScreenBoundaries() (ScreenSection, error) {
	zap.L().Info("Getting upper left boundaries in 2 seconds..")
	robotgo.Sleep(2)
	xh, yh := robotgo.Location()

	zap.L().Info("Upper left caught: ", zap.Int("x", xh), zap.Int("y", yh))

	zap.L().Info("Getting lower right boundaries in 2 seconds..")
	robotgo.Sleep(2)
	xl, yl := robotgo.Location()
	zap.L().Info("Upper left caught: ", zap.Int("x", xl), zap.Int("y", yl))

	if xl <= xh || yl <= yh {
		return ScreenSection{}, &IncorrectBondaries{}
	}

	return ScreenSection{ Point{xh, yh}, Point{ xl, yl}, xl - xh, yl - yh}, nil
}

// captures a section from the screen and returns it as image.Image
func CaptureScreenSection(section ScreenSection) (image.Image, error) {
	// getting the width and height
	w := section.LowerRight.X - section.UpperLeft.X
	h := section.LowerRight.Y - section.UpperLeft.Y	

	if w <= 0 || h <= 0 { // if w and h are invalid
		return nil, &IncorrectBondaries{}
	}

	// capturing bitmap
	bit := robotgo.CaptureScreen(section.UpperLeft.X, section.UpperLeft.Y, w, h)
	defer robotgo.FreeBitmap(bit)

	// converting to image
	img := robotgo.ToImage(bit)

	return img, nil
}

