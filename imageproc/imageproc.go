package imageproc

import (
	"encoding/json"
	"fmt"
	"image"
	"os"

	"github.com/joseCarlosAndrade/go-drawer/screen"
	"github.com/vcaesar/imgo"
	"go.uber.org/zap"

	"reflect"
)

// saves the img on the path as jpg. Returns an error if there's one
func SaveImageToJpeg(img image.Image, path string) error {
	err := imgo.SaveToJpeg(path, img)
	
	if err != nil {
		return err
	}

	fmt.Println("image saved as '", path, "'")
	return nil
}

func ResizeImage(img image.Image, height, width int) (image.Image, error) {
	if img == nil {
		return nil, &ImageIsNil{}
	}

	// todo
	return img, nil
}

// func 
func ImageInfo(img image.Image) {
	// fmt.Println(img)
	zap.L().Info("img.Bounds(): ", zap.Any("bounds", img.Bounds()))
	bounds := img.Bounds()
	width, height := bounds.Size().X, bounds.Size().Y

	zap.L().Info("w h: ", zap.Int("width", width), zap.Int("height", height))

	// fmt.Println(img)
	fmt.Println(reflect.TypeOf(img))

	switch v := img.(type) {
	case *image.RGBA:

		zap.L().Info("length: ", zap.Int("length", len(v.Pix)), zap.Int("mult", width*height))
			
	}
}

type CalibrationData struct {
	ColorPickerBoundaries screen.ScreenSection
	DrawingAreaBoundaries screen.ScreenSection
}

func LoadCalibrationData(path string) (CalibrationData, error) {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		return CalibrationData{}, err
	}

	var data CalibrationData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return CalibrationData{}, err
	}

	return data, nil
}

func SaveCalibrationData(data CalibrationData, path string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
