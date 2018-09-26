package kmeans

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"strconv"
)

/*MakeGif creates a gif from the generated images*/
func MakeGif(t int) {
	var fileStrs []string

	for i := 0; i < t; i++ {
		fileStr := "charts/" + strconv.Itoa(i) + ".png"
		fileStrs = append(fileStrs, fileStr)
	}

	// load static image and construct outGif
	outGif := &gif.GIF{}
	for _, name := range fileStrs {
		f, _ := os.Open(name)
		inPNG, _, _ := image.Decode(f)
		rect := inPNG.Bounds()
		palettedImage := image.NewPaletted(rect, palette.Plan9)
		draw.Draw(palettedImage, palettedImage.Rect, inPNG, rect.Min, draw.Over)
		f.Close()

		outGif.Image = append(outGif.Image, palettedImage)
		outGif.Delay = append(outGif.Delay, 10)
	}

	// save to out.gif
	f, _ := os.OpenFile("charts/animation.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
