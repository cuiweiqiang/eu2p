package eu2p

import (
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var EmojiMap map[string]image.Image

func init() {
	EmojiMap = map[string]image.Image{}
	basePath := "./img-apple-160"

	filepath.Walk(basePath, walkfunc)
}

func walkfunc(path string, info os.FileInfo, err error) error {
	ok := strings.HasSuffix(info.Name(), ".png")

	if ok {
		filename := strings.TrimSuffix(info.Name(), ".png")

		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		EmojiMap[filename] = img
	}

	return nil
}

func GetEmojiImg(code string, width, height uint) image.Image {
	img := EmojiMap[code]

	m := resize.Resize(width, height, img, resize.Lanczos3)

	return m

}
