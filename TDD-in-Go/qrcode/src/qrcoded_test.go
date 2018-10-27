package main

import (
	"testing"
	"bytes"
	"image/png"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-2368")

	if result == nil {
		t.Errorf("Generated QRcode is null")
	}
	if len(result) == 0 {
		t.Errorf("Generated QRcode has no data")
	}
}

func TestGenerateQRCodeReturnsPNG(t *testing.T) {
	result := GenerateQRCode("555-2368")
	buffer := bytes.NewBuffer(result)
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}