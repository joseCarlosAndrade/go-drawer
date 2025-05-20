package app

import (
	"image"

	"github.com/joseCarlosAndrade/go-drawer/imageproc"
	"github.com/joseCarlosAndrade/go-drawer/screen"
	"go.uber.org/zap"
)

// this pkg will handle the logic, calling the service packages

type App struct {
	ColorPickerBoundaries screen.ScreenSection
	DrawingAreaBoundaries screen.ScreenSection

	CollorPickerImage image.Image

	CurrentTheme string
	CurrentColor string
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	zap.L().Info("------------- Running app -------------")
	zap.L().Info("Loading calibration data")
	// try to load calibration data
	data, err := imageproc.LoadCalibrationData("imageproc/files/calibration.json")
	if err != nil {
		zap.L().Error("No calibration data found. Please run on '--mode calibration' first", zap.Error(err))
		return
	}

	a.ColorPickerBoundaries = data.ColorPickerBoundaries
	a.DrawingAreaBoundaries = data.DrawingAreaBoundaries

	zap.L().Info("Calibration data loaded")
	zap.L().Info("Color picker boundaries", zap.Int("x", a.ColorPickerBoundaries.UpperLeft.X), zap.Int("y", a.ColorPickerBoundaries.UpperLeft.Y), zap.Int("width", a.ColorPickerBoundaries.Width), zap.Int("height", a.ColorPickerBoundaries.Height))
	zap.L().Info("Drawing area boundaries", zap.Int("x", a.DrawingAreaBoundaries.UpperLeft.X), zap.Int("y", a.DrawingAreaBoundaries.UpperLeft.Y), zap.Int("width", a.DrawingAreaBoundaries.Width), zap.Int("height", a.DrawingAreaBoundaries.Height))
}

func (a* App) Calibrate() error {
	zap.L().Info("Clibration Started")

	zap.L().Info("Getting color picker boundaries")

	// getting color picker boundaries
	boundaries, err := screen.GetScreenBoundaries()
	if err != nil {
		zap.L().Error("Error getting screen boundaries", zap.Error(err))
		return err
	}

	zap.L().Info("Color picker boundaries", zap.Int("x", boundaries.UpperLeft.X), zap.Int("y", boundaries.UpperLeft.Y), zap.Int("width", boundaries.Width), zap.Int("height", boundaries.Height))

	a.ColorPickerBoundaries = boundaries

	zap.L().Info("Capturing color picker image")
	img, err := screen.CaptureScreenSection(boundaries)
	if err != nil {
		zap.L().Error("Error capturing color picker image", zap.Error(err))
		return err
	}
	a.CollorPickerImage = img

	// saving color picker image 
	err = imageproc.SaveImageToJpeg(img, "imageproc/files/color_picker.jpg")
	if err != nil {
		zap.L().Error("Error saving color picker image", zap.Error(err))
		return err
	}
	
	// getting drawing area boundaries
	zap.L().Info("Getting drawing area boundaries")

	boundaries, err = screen.GetScreenBoundaries()
	if err != nil {
		zap.L().Error("Error getting screen boundaries", zap.Error(err))
		return err
	}

	zap.L().Info("Drawing area boundaries", zap.Int("x", boundaries.UpperLeft.X), zap.Int("y", boundaries.UpperLeft.Y), zap.Int("width", boundaries.Width), zap.Int("height", boundaries.Height))

	a.DrawingAreaBoundaries = boundaries

	// capturing drawing area image
	img, err = screen.CaptureScreenSection(boundaries)
	if err != nil {
		zap.L().Error("Error capturing drawing area image", zap.Error(err))
		return err
	}

	// saving drawing area image
	err = imageproc.SaveImageToJpeg(img, "imageproc/files/drawing_area.jpg")
	if err != nil {
		zap.L().Error("Error saving drawing area image", zap.Error(err))
		return err
	}

	zap.L().Info("Calibration successful")

	// saving calibration data
	err = imageproc.SaveCalibrationData(imageproc.CalibrationData{
		ColorPickerBoundaries: a.ColorPickerBoundaries,
		DrawingAreaBoundaries: a.DrawingAreaBoundaries,
	}, "imageproc/files/calibration.json")

	return nil
}