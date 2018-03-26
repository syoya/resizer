package main

import (
	"flag"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

var (
	flagOutputDir string

	raw = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	rawFx2 = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	rawFx3 = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
)

func init() {
	flag.StringVar(&flagOutputDir, "output", "", "set output directory path")
}

func buildImage() (*image.RGBA, *image.Gray) {
	colors := make([]int, 15*21)
	copy(colors, rawFx3)

	imgRGBA := image.NewRGBA(image.Rect(0, 0, 15, 21))
	imgGray := image.NewGray(image.Rect(0, 0, 15, 21))

	for y := 0; y < 21; y++ {
		for x := 0; x < 15; x++ {
			var er, eg, eb, ea uint8
			e := colors[(15*y)+x]

			if e == 1 {
				er = 0xff
				eg = 0xff
				eb = 0xff
			} else {
				er = 0x00
				eg = 0x00
				eb = 0x00
			}
			ea = 0x80

			rgbaColor := color.RGBA{er, eg, eb, ea}
			imgRGBA.Set(x, y, rgbaColor)

			grayColor := color.GrayModel.Convert(rgbaColor)
			imgGray.Set(x, y, grayColor)
		}
	}

	return imgRGBA, imgGray
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func saveImages(outputDir string) (err error) {
	_, imgGray := buildImage()

	var fPngAlpha *os.File
	fPngAlpha, err = os.OpenFile(filepath.Join(outputDir, "f.png"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer func() {
		err = fPngAlpha.Close()
	}()
	err = png.Encode(fPngAlpha, imgGray)
	if err != nil {
		return err
	}

	var fJpg *os.File
	fJpg, err = os.OpenFile(filepath.Join(outputDir, "f.jpg"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer func() {
		err = fJpg.Close()
	}()
	err = jpeg.Encode(fJpg, imgGray, nil)
	if err != nil {
		return err
	}

	var fGif *os.File
	fGif, err = os.OpenFile(filepath.Join(outputDir, "f.gif"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer func() {
		err = fGif.Close()
	}()

	err = gif.Encode(fGif, imgGray, nil)
	if err != nil {
		return err
	}

	return nil
}

func isExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	flag.Parse()

	if flagOutputDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !isExists(flagOutputDir) {
		os.MkdirAll(flagOutputDir, 0755)
	}

	checkErr(saveImages(flagOutputDir))

}
