package imageproc

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/imgo"
)

func GetScreenBoundaries() (ScreenSection, error) {
	fmt.Println("Getting upper left boundaries in 2 seconds..")
	robotgo.Sleep(2)
	xh, yh := robotgo.Location()
	fmt.Println("Upper left caught: ", xh, yh)

	fmt.Println("Getting lower right boundaries in 2 seconds..")
	robotgo.Sleep(2)
	xl, yl := robotgo.Location()
	fmt.Println("Upper left caught: ", xl, yl)

	if xl <= xh || yl <= yh {
		return ScreenSection{}, &IncorrectBondaries{}
	}

	return ScreenSection{ Point{xh, yh}, Point{ xl, yl}}, nil
}

func CaptureScreenSection(section ScreenSection) {
	// getting the width and height
	w := section.LowerRight.x - section.UpperLeft.x
	h := section.LowerRight.y - section.UpperLeft.y	

	// capturing bitmap
	bit := robotgo.CaptureScreen(section.UpperLeft.x, section.UpperLeft.y, w, h)
	defer robotgo.FreeBitmap(bit)

	// converting to image
	img := robotgo.ToImage(bit)

	imgName := "imageproc/images/section.jpg"

	imgo.SaveToJpeg(imgName, img)
	

	fmt.Println("image saved as '", imgName, "'")
}

// func 