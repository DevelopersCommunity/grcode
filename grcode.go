package main

import (
	"flag"
	"os"

	"github.com/DevelopersCommunity/grcode/qr"
	"github.com/golang/glog"
)

var helpFlag bool
var textFlag string
var outFlag string
var logoFlag string

func init() {
	const (
		helpUsage = "show this help message and exit"
		textUsage = "text to embed in QR Code"
		outUsage  = "QR Code output file (PNG)"
		logoUsage = "logo to be included in QR Code (PNG)"
	)

	flag.BoolVar(&helpFlag, "help", false, helpUsage)
	flag.BoolVar(&helpFlag, "h", false, helpUsage+" (shorthand)")

	flag.StringVar(&textFlag, "text", "", textUsage)
	flag.StringVar(&textFlag, "t", "", textUsage+" (shorthand)")

	flag.StringVar(&outFlag, "out", "", outUsage)
	flag.StringVar(&outFlag, "o", "", outUsage+" (shorthand)")

	flag.StringVar(&logoFlag, "logo", "", logoUsage)
	flag.StringVar(&logoFlag, "l", "", logoUsage+" (shorthand)")
}

func usage() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Usage()
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if helpFlag || textFlag == "" || outFlag == "" {
		usage()
		return
	}

	err := qr.CreateQRCode(textFlag, outFlag, logoFlag)

	if err != nil {
		glog.Errorf("%+v", err)
	}
}
