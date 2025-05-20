package main

import (
	// "fmt"

	// "github.com/go-vgo/robotgo"
	// "github.com/vcaesar/imgo"
	// "strconv"
	"flag"
	"strings"

	"github.com/joseCarlosAndrade/go-drawer/app"
	// "github.com/joseCarlosAndrade/go-drawer/imageproc"
	// "github.com/joseCarlosAndrade/go-drawer/screen"

	// "github.com/joseCarlosAndrade/go-drawer/mousec"
	"go.uber.org/zap"
)

var logger *zap.Logger
var development bool = true

func init() {
  if development {
    logger, _ = zap.NewDevelopment()
  } else {
    logger, _ = zap.NewProduction()
  }
  logger.Info("Logger initialized")
}

func main() {

  undo := zap.ReplaceGlobals(logger)
  defer undo()
  defer logger.Sync()

  var app *app.App = app.NewApp()

  flagMode := flag.String("mode", "app", "start in app or calibration mode")
  flag.Parse()
  if s := strings.ToLower(*flagMode); s == "app" {
      app.Run()
  } else if s == "calibration" {
      app.Calibrate()
  } else {
    panic("Invalid mode")
  }
}