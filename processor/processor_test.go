package processor_test

import (
	"bytes"
	"image"
	"math"
	"testing"

	"github.com/syoya/resizer/input"
	"github.com/syoya/resizer/processor"
	"github.com/syoya/resizer/storage"
)

const (
	u                        = 3
	png                      = "../fixtures/f-png24.png"
	maxEuclideanDistanceRGBA = 131070 // math.Sqrt(math.Pow(0xff, 2) * 4)
)

var (
	formats = []string{
		"../fixtures/f.jpg",
		"../fixtures/f-png8.png",
		"../fixtures/f-png24.png",
		"../fixtures/f.gif",
	}
	orientations = []string{
		"../fixtures/f-orientation-1.jpg",
		"../fixtures/f-orientation-2.jpg",
		"../fixtures/f-orientation-3.jpg",
		"../fixtures/f-orientation-4.jpg",
		"../fixtures/f-orientation-5.jpg",
		"../fixtures/f-orientation-6.jpg",
		"../fixtures/f-orientation-7.jpg",
		"../fixtures/f-orientation-8.jpg",
	}
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
)

func diff(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

func isNear(a, b uint32) bool {
	return diff(a, b) <= math.MaxUint8*4
}

// calcEuclideanDistance 2つのピクセルのRGBAからユークリッド距離を求める
func calcEuclideanDistance(r1, g1, b1, a1, r2, b2, g2, a2 uint32) float64 {
	r := math.Pow((float64(r1)/float64(0xffff))-(float64(r2)/float64(0xffff)), 2)
	g := math.Pow((float64(g1)/float64(0xffff))-(float64(g2)/float64(0xffff)), 2)
	b := math.Pow((float64(b1)/float64(0xffff))-(float64(b2)/float64(0xffff)), 2)
	a := math.Pow((float64(a1)/float64(0xffff))-(float64(a2)/float64(0xffff)), 2)
	return math.Sqrt(r + g + b + a)
}

func evalPixels(t *testing.T, path string, i image.Image, p image.Point, colors []int) {
	for y := 0; y < p.Y; y++ {
		for x := 0; x < p.X; x++ {
			var er, eg, eb, ea uint32
			e := colors[p.X*y+x]
			if e == 1 {
				er = 0xffff
				eg = 0xffff
				eb = 0xffff
			} else {
				er = 0x00
				eg = 0x00
				eb = 0x00
			}
			ea = 0xffff

			a := i.At(u/2>>0+u*x, u/2>>0+u*y)
			ar, ag, ab, aa := a.RGBA()

			// 距離が0.3を超える場合は近似色とはしない
			if distance := calcEuclideanDistance(er, eg, eb, ea, ar, ag, ab, aa); distance > 0.3 {
				t.Errorf(
					"wrong color at (%d, %d) expected {%d, %d, %d, %d}, but actual {%d, %d, %d, %d}, distance=%v, path=%s",
					x, y,
					er, eg, eb, ea,
					ar, ag, ab, aa,
					distance,
					path,
				)
			}
		}
	}
}

func eval(t *testing.T, path string, f storage.Image, size image.Point, colors []int) string {
	var b []byte
	w := bytes.NewBuffer(b)
	p := processor.New()
	f.ValidatedWidth *= u
	f.ValidatedHeight *= u
	pixels, err := p.Preprocess(path)
	if err != nil {
		t.Fatalf("cannot preprocess image: err=%v, path=%s", err, path)
	}

	f, err = f.Normalize(pixels.Bounds().Size())
	if err != nil {
		t.Fatalf("fail to normalize: err=%v, path=%s", err, path)
	}

	if _, err := p.Resize(pixels, w, f); err != nil {
		t.Fatalf("cannot process image: err=%v path=%s", err, path)
	}

	r := bytes.NewReader(w.Bytes())
	img, format, err := image.Decode(r)
	if err != nil {
		t.Fatalf("cannot decode image: err=%v, path=%s", err, path)
	}

	expectedSize := size.Mul(u)
	rect := img.Bounds()
	actualSize := rect.Size()
	if !actualSize.Eq(expectedSize) {
		t.Fatalf("wrong size expected %v, but actual %v, path=%s", expectedSize, actualSize, path)
	}

	evalPixels(t, path, img, size, colors)

	return format
}

func TestFormats(t *testing.T) {
	size := image.Point{5, 7}
	colors := []int{
		0, 0, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 1, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
		0, 0, 0, 0, 0,
	}
	for _, format := range []string{input.FormatPNG, input.FormatJPEG, input.FormatGIF} {
		f := storage.Image{
			ValidatedMethod:  input.MethodContain,
			ValidatedWidth:   5,
			ValidatedHeight:  7,
			ValidatedFormat:  format,
			ValidatedQuality: 100,
		}
		for _, path := range formats {
			eval(t, path, f, size, colors)
		}
	}
}

func TestOrientations(t *testing.T) {
	f := storage.Image{
		ValidatedMethod: input.MethodContain,
		ValidatedWidth:  5,
		ValidatedHeight: 7,
		ValidatedFormat: input.FormatPNG,
	}
	size := image.Point{5, 7}
	colors := []int{
		0, 0, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 1, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
		0, 0, 0, 0, 0,
	}
	for _, path := range orientations {
		eval(t, path, f, size, colors)
	}
}

func TestFormatNormal(t *testing.T) {
	size := image.Point{5, 7}
	colors := []int{
		0, 0, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 1, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
		0, 0, 0, 0, 0,
	}
	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodContain,
		ValidatedWidth:  5,
		ValidatedHeight: 100,
		ValidatedFormat: input.FormatPNG,
	}, size, colors)
	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodContain,
		ValidatedWidth:  100,
		ValidatedHeight: 7,
		ValidatedFormat: input.FormatPNG,
	}, size, colors)
}

func TestFormatThumbnail(t *testing.T) {
	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  3,
		ValidatedHeight: 7,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{3, 7}, []int{
		0, 0, 0,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		1, 0, 0,
		1, 0, 0,
		0, 0, 0,
	})

	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  5,
		ValidatedHeight: 3,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{5, 3}, []int{
		0, 1, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 1, 0, 0, 0,
	})

	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  100,
		ValidatedHeight: 100,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{10, 14}, []int{
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
	})

	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  6,
		ValidatedHeight: 100,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{6, 14}, []int{
		0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1,
		1, 1, 0, 0, 0, 0,
		1, 1, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1,
		1, 1, 0, 0, 0, 0,
		1, 1, 0, 0, 0, 0,
		1, 1, 0, 0, 0, 0,
		1, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0,
	})

	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  2,
		ValidatedHeight: 100,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{2, 14}, []int{
		0, 0,
		0, 0,
		1, 1,
		1, 1,
		0, 0,
		0, 0,
		1, 1,
		1, 1,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
	})

	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  100,
		ValidatedHeight: 10,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{10, 10}, []int{
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
	})

	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  100,
		ValidatedHeight: 6,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{10, 6}, []int{
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
	})

	eval(t, png, storage.Image{
		ValidatedMethod: input.MethodCover,
		ValidatedWidth:  100,
		ValidatedHeight: 2,
		ValidatedFormat: input.FormatPNG,
	}, image.Point{10, 2}, []int{
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 1, 1, 1, 1, 1, 1, 0, 0,
	})
}
