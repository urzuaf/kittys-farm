package load

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFromImage(path string, fs embed.FS) *ebiten.Image {
	file, err := fs.ReadFile(path)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)

}

func LoadFont(fs embed.FS, path string) font.Face {
	fontData, err := fs.ReadFile(path)
	if err != nil {
		panic(err)
	}
	tt, err := opentype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    64,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
	return face
}
