package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

const (
	OutputDir = "./book"
	ScanDir   = "."
)

func init() {
	//logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

func main() {
	myFigure := figure.NewColorFigure("hello", "", "green", true)
	myFigure.Print()

	RenderMd(ScanDir)
	RenderHTML(filepath.Join(ScanDir, OutputDir))
}
