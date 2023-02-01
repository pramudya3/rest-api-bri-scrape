package helper

import (
	"log"

	"github.com/otiai10/gosseract/v2"
)

func Captcha2Text(captcha []byte) string {
	client := gosseract.NewClient()
	client.SetImageFromBytes(captcha)
	text, err := client.Text()
	if err != nil {
		log.Fatal("error parse image to text", err)
	}
	return text
}
