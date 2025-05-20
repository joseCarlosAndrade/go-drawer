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
	CurrentImage image.Image
	FinalDrawingImage image.Image

	CurrentTheme string
	CurrentColor string

	CurrentState AppState
}

func NewApp() *App {
	return &App{
		CurrentState: AppStateStartup,
	}
}

func (a *App) Run() {
	zap.L().Info("------------- Running App -------------")
	zap.L().Info("Loading calibration data")
	// try to load calibration data
	data, err := imageproc.LoadCalibrationData("imageproc/files/calibration.json")
	if err != nil {
		zap.L().Error("No calibration data found. Please run on '--mode calibration' first", zap.Error(err))
		return
	}

	a.ColorPickerBoundaries = data.ColorPickerBoundaries
	a.DrawingAreaBoundaries = data.DrawingAreaBoundaries

	// load color picker image
	img, err := imageproc.LoadImage("imageproc/files/color_picker.jpeg")
	if err != nil {
		zap.L().Error("Calibration incomplete. Please run on '--mode calibration' first", zap.Error(err))
		
		return
	}
	
	a.CollorPickerImage = img

	zap.L().Info("Calibration data loaded")
	// zap.L().Info("Color picker boundaries", zap.Int("x", a.ColorPickerBoundaries.UpperLeft.X), zap.Int("y", a.ColorPickerBoundaries.UpperLeft.Y), zap.Int("width", a.ColorPickerBoundaries.Width), zap.Int("height", a.ColorPickerBoundaries.Height))
	// zap.L().Info("Drawing area boundaries", zap.Int("x", a.DrawingAreaBoundaries.UpperLeft.X), zap.Int("y", a.DrawingAreaBoundaries.UpperLeft.Y), zap.Int("width", a.DrawingAreaBoundaries.Width), zap.Int("height", a.DrawingAreaBoundaries.Height))

	// start the app
	a.Game()
}

func (a* App) Calibrate() error {
	zap.L().Info("--------------------------------- Clibration Started ---------------------------------")

	zap.L().Info("Getting color picker boundaries")

	a.CurrentState = AppStateCalibration

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
	err = imageproc.SaveImageToJpeg(img, "imageproc/files/color_picker.jpeg")
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
	err = imageproc.SaveImageToJpeg(img, "imageproc/files/drawing_area.jpeg")
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

func (a *App) Game() error {
	zap.L().Info("--------------------------------- Game Started ---------------------------------")

	zap.L().Info("Press 'q' to exit")

	a.CurrentState = AppStateIdle
	
	// await for user input
	for {
		// wait for enter key
		switch a.CurrentState {
		case AppStateIdle:
			zap.L().Info("Press enter to get the next image")
			screen.WaitForKeyboardInput()
			a.CurrentState = AppStateGettingImage

		case AppStateGettingImage:
			img, err := a.GetNextImage()
			if err != nil {
				zap.L().Error("Error getting next image. Returning to idle state", zap.Error(err))
				a.CurrentState = AppStateIdle
				continue
			}

			a.CurrentImage = img
			a.CurrentState = AppStateProcessingImage

		case AppStateProcessingImage:
			err := a.ProcessImage()
			if err != nil {
				zap.L().Error("Error processing image. Returning to idle state", zap.Error(err))
				a.CurrentState = AppStateIdle
				continue
			}
			a.CurrentState = AppStateWaitingToStartDrawing

		case AppStateWaitingToStartDrawing:
			zap.L().Info("Everything is ready! Press any key to start drawing")
			screen.WaitForKeyboardInput()
			a.CurrentState = AppStateDrawing

		case AppStateDrawing:
			err := a.Draw()
			if err != nil {
				zap.L().Error("Error drawing image. Returning to idle state", zap.Error(err))
				a.CurrentState = AppStateIdle
				continue
			}
			a.CurrentState = AppStateIdle
		}

		if a.CurrentState == AppStateFinished {
			break
		}
	}

	return nil
}

func (a *App) GetNextImage() (image.Image, error) {
	boundaries, err := screen.GetScreenBoundaries()
	if err != nil {
		return nil, err
	}
	
	img, err := screen.CaptureScreenSection(boundaries)
	if err != nil {
		return nil, err
	}
	
	return img, nil
}

func (a *App) ProcessImage() error {
	
	zap.L().Info("Processing image")

	a.FinalDrawingImage = a.CurrentImage // TODO: process the image

	return nil
}

func (a *App) Draw() error {
	zap.L().Info("Drawing image")

	zap.L().Info("Finished drawing image!")
	return nil
}
