package app

// enum to handle the app state
type AppState int

const (
	AppStateStartup AppState = iota
	AppStateIdle
	AppStateCalibration
	AppStateGettingImage
	AppStateProcessingImage
	AppStateWaitingToStartDrawing
	AppStateDrawing
	AppStateFinished
)

func (s AppState) String() string {
	return []string{"Startup", "Idle", "Calibration", "GettingImage", "ProcessingImage", "Drawing"}[s]
}
