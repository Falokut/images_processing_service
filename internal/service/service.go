package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Falokut/image_processing_service/internal/imageprocessing"
	"github.com/Falokut/image_processing_service/internal/models"
)

type ImageProcessingService interface {
	Crop(ctx context.Context, imageBody []byte, x0, y0, x1, y1 uint32) (res *models.ProcessedImageResponse, err error)
	Resize(ctx context.Context,
		imageBody []byte, width, height int32, resampleFilter string) (res *models.ProcessedImageResponse, err error)
	Validate(ctx context.Context,
		in *models.ValidateRequest) (res models.ValidateResponse, err error)
	Hue(ctx context.Context,
		imageBody []byte, hue int32) (res *models.ProcessedImageResponse, err error)
	Desaturate(ctx context.Context, imageBody []byte) (res *models.ProcessedImageResponse, err error)
	Blur(ctx context.Context,
		imageBody []byte, radius float64, method string) (res *models.ProcessedImageResponse, err error)
}

type imageProcessingService struct {
	imagesProcessing imageprocessing.ImagesProcessing
}

func NewImagesProcessingService(
	imagesProcessing imageprocessing.ImagesProcessing) *imageProcessingService {
	return &imageProcessingService{
		imagesProcessing: imagesProcessing,
	}
}

func (s *imageProcessingService) Crop(ctx context.Context,
	imageBody []byte, x0, y0, x1, y1 uint32) (res *models.ProcessedImageResponse, err error) {
	img, Type, err := imageprocessing.DecodeImage(imageBody)
	if err != nil {
		return
	}

	encoded, err := imageprocessing.EncodeImage(s.imagesProcessing.Crop(ctx, img,
		int(x0), int(y0), int(x1), int(y1)), Type.String())
	if err != nil {
		return
	}
	res = &models.ProcessedImageResponse{
		ContentType: Type.String(),
		Img:         encoded,
	}
	return
}

func (s *imageProcessingService) Resize(ctx context.Context,
	imageBody []byte, width, height int32, resampleFilter string) (res *models.ProcessedImageResponse, err error) {
	img, Type, err := imageprocessing.DecodeImage(imageBody)
	if err != nil {
		return
	}

	encoded, err := imageprocessing.EncodeImage(s.imagesProcessing.Resize(ctx, img,
		int(width), int(height), resampleFilter), Type.String())
	if err != nil {
		return
	}

	res = &models.ProcessedImageResponse{
		ContentType: Type.String(),
		Img:         encoded,
	}
	return
}

func (s *imageProcessingService) Validate(_ context.Context,
	in *models.ValidateRequest) (res models.ValidateResponse, err error) {
	checked := false

	img, format, err := imageprocessing.DecodeImage(in.Img)
	if err != nil {
		return
	}

	if len(in.SupportedTypes) != 0 {
		supported := false
		checked = true
		for _, ftype := range in.SupportedTypes {
			if strings.EqualFold(ftype, format.String()) {
				supported = true
				break
			}
		}

		if !supported {
			res.Details = fmt.Sprintf("image has unsupported type, supported types: %s, image has: %s",
				strings.Join(in.SupportedTypes, ","), format)
			return
		}
	}

	if in.MaxHeight != nil {
		checked = true
		if img.Bounds().Dy() > int(*in.MaxHeight) {
			res.Details = fmt.Sprintf("image has height bigger than max height, image height: %d max height: %d",
				img.Bounds().Dy(), *in.MaxHeight)
			return
		}
	}

	if in.MaxWidth != nil {
		checked = true
		if img.Bounds().Dx() > int(*in.MaxWidth) {
			res.Details = fmt.Sprintf("image has width bigger than max width, image width: %d max width: %d",
				img.Bounds().Dx(), *in.MaxWidth)
			return
		}
	}

	if in.MinHeight != nil {
		checked = true
		if img.Bounds().Dy() < int(*in.MinHeight) {
			res.Details = fmt.Sprintf("image has height less than min height, image height: %d min height: %d",
				img.Bounds().Dy(), *in.MinHeight)
			return
		}
	}

	if in.MinWidth != nil {
		checked = true
		if img.Bounds().Dx() < int(*in.MinWidth) {
			res.Details = fmt.Sprintf("image has width less than min width, image width: %d min width: %d",
				img.Bounds().Dx(), *in.MinWidth)
			return
		}
	}

	if !checked {
		res.Details = "no received instructions for image validation"
		return
	}

	return
}

func (s *imageProcessingService) Hue(ctx context.Context,
	imageBody []byte, hue int32) (res *models.ProcessedImageResponse, err error) {
	img, Type, err := imageprocessing.DecodeImage(imageBody)
	if err != nil {
		return
	}

	encoded, err := imageprocessing.EncodeImage(s.imagesProcessing.Hue(ctx, img, int(hue)), Type.String())
	if err != nil {
		return
	}

	res = &models.ProcessedImageResponse{
		ContentType: Type.String(),
		Img:         encoded,
	}
	return
}

func (s *imageProcessingService) Desaturate(ctx context.Context,
	imageBody []byte) (res *models.ProcessedImageResponse, err error) {
	img, Type, err := imageprocessing.DecodeImage(imageBody)
	if err != nil {
		return
	}

	encoded, err := imageprocessing.EncodeImage(s.imagesProcessing.Desaturate(ctx, img), Type.String())
	if err != nil {
		return nil, err
	}

	res = &models.ProcessedImageResponse{
		ContentType: Type.String(),
		Img:         encoded,
	}
	return
}

func (s *imageProcessingService) Blur(ctx context.Context,
	imageBody []byte, radius float64, method string) (res *models.ProcessedImageResponse, err error) {
	img, Type, err := imageprocessing.DecodeImage(imageBody)
	if err != nil {
		return
	}

	encoded, err := imageprocessing.EncodeImage(s.imagesProcessing.Blur(ctx, img, radius, method), Type.String())
	if err != nil {
		return
	}

	res = &models.ProcessedImageResponse{
		ContentType: Type.String(),
		Img:         encoded,
	}
	return
}
