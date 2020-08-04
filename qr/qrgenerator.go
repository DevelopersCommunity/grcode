package qr

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
)

// CreateQRCode creates a QR Code.
func CreateQRCode(text string, out string, logo string) error {
	q, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return errors.Wrap(err, "QR Code creation failed")
	}
	qrImage := q.Image(256)
	qrLogo := image.NewRGBA(image.Rect(0, 0, 256, 256))
	draw.Draw(qrLogo, qrLogo.Bounds(), qrImage, image.Point{0, 0}, draw.Src)

	if logo != "" {
		reader, err := os.OpenFile(logo, os.O_RDONLY, 0)
		if err != nil {
			return errors.Wrap(err, "Open file failed")
		}
		defer reader.Close()

		logoImage, _, err := image.Decode(reader)
		if err != nil {
			return errors.Wrap(err, "Image decode failed")
		}
		logoImage = resize.Resize(64, 64, logoImage, resize.MitchellNetravali)
		draw.Draw(qrLogo, qrLogo.Bounds(), logoImage, image.Point{-96, -96}, draw.Src)
	}

	writer, err := os.Create(out)
	if err != nil {
		return errors.Wrap(err, "Create file failed")
	}
	defer writer.Close()
	err = png.Encode(writer, qrLogo)
	if err != nil {
		return errors.Wrap(err, "PNG encode failed")
	}

	return nil
}
