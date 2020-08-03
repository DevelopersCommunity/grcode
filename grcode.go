package main

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
)

var helpFlag bool
var textFlag string
var outFlag string

func init() {
	const (
		helpUsage = "show this help message and exit"
		textUsage = "text to embed in QR Code"
		outUsage  = "QR Code output file"
	)

	flag.BoolVar(&helpFlag, "help", false, helpUsage)
	flag.BoolVar(&helpFlag, "h", false, helpUsage+" (shorthand)")

	flag.StringVar(&textFlag, "text", "", textUsage)
	flag.StringVar(&textFlag, "t", "", textUsage+" (shorthand)")

	flag.StringVar(&outFlag, "out", "", outUsage)
	flag.StringVar(&outFlag, "o", "", outUsage+" (shorthand)")
}

func usage() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Usage()
}

func createQrCode() error {
	err := qrcode.WriteFile(textFlag, qrcode.Medium, 256, outFlag)

	return errors.Wrap(err, "QR Code write file failed")
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if helpFlag || textFlag == "" || outFlag == "" {
		usage()
		return
	}

	err := createQrCode()

	if err != nil {
		glog.Errorf("%+v", err)
	}
}
