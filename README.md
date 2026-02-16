# Go Rasterize
Simple interface to render an SVG to a go image.
Uses a modified version of [`github.com/srwiley/oksvg`](https://github.com/srwiley/oksvg) to do the heavy lifing.
If your use case is rendering to PNG, you might want to use that instead.


# Example


```golang

func renderToSquare(svgData []byte, sideLength int) (image.Image, error) {
    return rasterize.Rasterize(svgData, sideLength, sideLendth)
}


func renderFullSizeWithYellowOnGreenBackground(svgData[]) (image.Image, error) {
    return rasterize.Rasterize(svgdata,
                               0,
                               0, 
                               rasterize.WithBackgroundColor(color.RGBA{R:0, G: 255, B:0, A: 255}),
                               rasterize.WithUniformColor(color.RGBA{R: 255, G: 255, B: 0, A: 255}))
}
```