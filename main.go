package main

import (
  // "fmt"

  // "github.com/go-vgo/robotgo"
  // "github.com/vcaesar/imgo"
  // "strconv"
  "github.com/joseCarlosAndrade/go-drawer/imageproc"
  // "github.com/joseCarlosAndrade/go-drawer/mousec"
)

func main() {
  boundaries, err := imageproc.GetScreenBoundaries()

  if err != nil {
    panic(err)
  }

  imageproc.CaptureScreenSection(boundaries)
}