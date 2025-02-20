package load

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
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
