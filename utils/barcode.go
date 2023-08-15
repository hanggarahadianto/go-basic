package utils

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func Barcode(){

	qrCode, _ := qr.Encode("https://letmegooglethat.com/?q=docker", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 300, 300)

	// create the output file
	file, _ := os.Create("uploads/qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

}
