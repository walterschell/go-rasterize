package rasterize

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log/slog"

	"github.com/walterschell/oksvg"
	"github.com/srwiley/rasterx"
)

var log *slog.Logger = slog.New(slog.DiscardHandler)

// Sets the logger for this library.
// Not threadsafe
func SetLogger(logger *slog.Logger) {
	log = logger.With(slog.Attr{Key: "package", Value: slog.StringValue("rasterize")})
}

// Rasterize converts SVG data to a raster image with the specified width and height.
// If width and height are zero, it uses the SVG's intrinsic dimensions.
func Rasterize(svgData []byte, width, height int, options ...Option) (image.Image, error) {
	opts, err := applyOptions(options)
	if err != nil {
		return nil, err
	}

	icon, err := oksvg.ReadIconStream(bytes.NewBuffer(svgData))
	if err != nil {
		return nil, err
	}

	if opts.UniformDrawColor != nil {
		for i, path := range icon.SVGPaths {

			if path.GegFillColorRaw() == nil && path.GetLineColorRaw() == nil {
				log.Info("not setting line/fill color")
			} else {
				log.Info("setting line/fill color", " color", opts.UniformDrawColor)
				icon.SVGPaths[i].SetLineColor(opts.UniformDrawColor)
				icon.SVGPaths[i].SetFillColor(opts.UniformDrawColor)
			}
		}
	}
	if width == 0 || height == 0 {
		if width != 0 || height != 0 {
			return nil, fmt.Errorf("both width and height must be zero if one is zero (h=%d, w=%d)", height, width)
		}
		height = int(icon.ViewBox.H)
		width = int(icon.ViewBox.W)
	}

	icon.SetTarget(0, 0, float64(width), float64(height))
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	// Fill background
	drawColor := opts.BackgroundColor
	if drawColor == nil {
		drawColor = color.Black
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgba.Set(x, y, drawColor)
		}
	}
	drawer := rasterx.NewDasher(width, height, rasterx.NewScannerGV(width, height, rgba, rgba.Bounds()))
	icon.Draw(drawer, 1.0)
	return rgba, nil
}
