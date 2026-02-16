package rasterize

import (
	"bytes"
	_ "embed"
	"image/color"
	"image/jpeg"
	"log/slog"
	"os"
	"testing"
)

const OUTPUT_TEST_JPG = false

//go:embed testdata/vpn-connection.svg
var vpnConnectionSVG []byte

func TestMain(m *testing.M) {
	SetLogger(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	m.Run()
}

func TestRastertize(t *testing.T) {
	img, err := Rasterize(vpnConnectionSVG, 64, 64)
	if err != nil {
		t.Fatalf("Rasterize failed: %v", err)
	}
	if img.Bounds().Dx() != 64 || img.Bounds().Dy() != 64 {
		t.Fatalf("Unexpected image dimensions: got %dx%d, want 64x64", img.Bounds().Dx(), img.Bounds().Dy())
	}
}

func TestRastertizeWithDims(t *testing.T) {
	green := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	yellow := color.RGBA{R: 255, G: 255, B: 0, A: 255}
	img, err := Rasterize(
		vpnConnectionSVG,
		160,
		160,
		WithBackgroundColor(yellow),
		WithUniformColor(green))
	if err != nil {
		t.Fatalf("Rasterize failed: %v", err)
	}
	if img.Bounds().Dx() != 160 || img.Bounds().Dy() != 160 {
		t.Fatalf("Unexpected image dimensions: got %dx%d, want 160x160", img.Bounds().Dx(), img.Bounds().Dy())
	}

	if img.At(0, 0) != yellow {
		// t.Fatalf("Unexpected background color")
	}
	if OUTPUT_TEST_JPG {
		outBuffer := &bytes.Buffer{}
		jpgOpts := &jpeg.Options{
			Quality: 100,
		}

		err = jpeg.Encode(outBuffer, img, jpgOpts)
		if err != nil {
			t.Fatal("Can't encode")
		}
		os.WriteFile("test.jpg", outBuffer.Bytes(), 0o644)
	}
}
