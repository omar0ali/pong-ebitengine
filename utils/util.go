package utils

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"os"

	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadImage(path string) *ebiten.Image {
	file, err := os.ReadFile("assets/" + path)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

var scoreFont font.Face

func LoadFont(path string, size float64) font.Face {
	b, err := os.ReadFile("assets/" + path)
	if err != nil {
		log.Fatal(err)
	}

	fnt, err := opentype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return face
}

func LoadImagesFromFolder(folder string) []*ebiten.Image {
	images := []*ebiten.Image{}

	assetFolder := filepath.Join("assets", folder)
	files, err := os.ReadDir(assetFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(assetFolder, file.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}

		images = append(images, ebiten.NewImageFromImage(img))
	}

	return images
}
