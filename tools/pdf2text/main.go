package main

import (
	"log"
	"os"
	"strconv"

	"github.com/otiai10/gosseract/v2"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please provide a file path and an action")
	}
	file_path := os.Args[2]
	action := os.Args[1]
	switch action {
	case "pdf2img":
		pdf_to_image(file_path)
	case "img2txt":
		text := image_to_text(file_path)
		log.Printf("Text: %s", text)
	default:
		log.Fatal("Please provide a valid action")
	}
}

func image_to_text(file_path string) string {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(file_path)
	text, _ := client.Text()
	return text
}
func pdf_to_image(file_path string) error {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	defer mw.Destroy()
	mw.SetResolution(300, 300)
	err := mw.ReadImage(file_path)
	if err != nil {
		return err
	}
	// -background white -alpha remove
	pxw := imagick.NewPixelWand()
	defer pxw.Destroy()
	pxw.SetColor("white")
	mw.SetImageBackgroundColor(pxw)
	// Loop through pages
	for i := 0; i < int(mw.GetNumberImages()); i++ {
		mw.SetIteratorIndex(i) // This being the page offset

		mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE)

		mw.SetImageFormat("jpg")
		// Remove .pdf extension and add .jpg
		file_path = file_path[:len(file_path)-4] + "-" + strconv.Itoa(i) + ".jpg"
		log.Printf("Writing %s", file_path)
		err = mw.WriteImage(file_path)
		if err != nil {
			return err
		}
	}
	return nil
}
