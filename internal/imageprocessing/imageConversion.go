package imageprocessing

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/Falokut/image_processing_service/internal/models"
	"github.com/disintegration/imaging"
	"github.com/gabriel-vasile/mimetype"
)

func EncodeImage(img image.Image, extension string) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error

	format, err := imaging.FormatFromExtension(extension)
	if err != nil {
		return []byte{}, models.Error(models.Internal, err.Error())
	}
	err = imaging.Encode(buf, img, format, imaging.PNGCompressionLevel(png.BestSpeed), imaging.JPEGQuality(jpeg.DefaultQuality))

	if err != nil {
		return []byte{}, models.Error(models.Internal, err.Error())
	}

	return buf.Bytes(), nil
}

func DecodeImage(img []byte) (decoded image.Image, mimeType *mimetype.MIME, err error) {
	mimeType = mimetype.Detect(img)
	buf := bytes.NewBuffer(img)
	decoded, err = imaging.Decode(buf)
	if err != nil {
		err = models.Error(models.Internal, err.Error())
		return
	}

	return
}

func ConvertToBase64(img []byte) []byte {
	base64Encoded := make([]byte, base64.StdEncoding.EncodedLen(len(img)))
	base64.StdEncoding.Encode(base64Encoded, img)
	return base64Encoded
}

func GetMimeTypeExt(img []byte) string {
	return mimetype.Detect(img).Extension()
}
