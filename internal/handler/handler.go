package handler

import (
	"context"
	"errors"

	"github.com/Falokut/image_processing_service/internal/models"
	"github.com/Falokut/image_processing_service/internal/service"
	image_service "github.com/Falokut/image_processing_service/pkg/image_processing_service/v1/protos"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageProcessingServiceHandler struct {
	image_service.UnimplementedImageProcessingServiceV1Server
	s service.ImageProcessingService
}

func NewImageProcessingServiceHandler(s service.ImageProcessingService) *ImageProcessingServiceHandler {
	return &ImageProcessingServiceHandler{s: s}
}

func (h *ImageProcessingServiceHandler) Crop(ctx context.Context,
	in *image_service.CropRequest) (res *httpbody.HttpBody, err error) {

	if in.Image == nil {
		return nil, status.Error(codes.InvalidArgument, "image mustn't be empty")
	}
	defer h.handleError(&err)

	processed, err := h.s.Crop(ctx,
		in.Image.Image, in.StartX, in.StartY, in.EndX, in.EndY)
	if err != nil {
		return
	}
	res = &httpbody.HttpBody{
		ContentType: processed.ContentType,
		Data:        processed.Img,
	}
	return
}

func (h *ImageProcessingServiceHandler) Resize(ctx context.Context,
	in *image_service.ResizeRequest) (res *httpbody.HttpBody, err error) {
	if in.Image == nil {
		return nil, status.Error(codes.InvalidArgument, "image mustn't be empty")
	}
	defer h.handleError(&err)

	processed, err := h.s.Resize(ctx,
		in.Image.Image, in.Width, in.Height, in.ResampleFilter.String())
	if err != nil {
		return
	}
	res = &httpbody.HttpBody{
		ContentType: processed.ContentType,
		Data:        processed.Img,
	}
	return
}

func (h *ImageProcessingServiceHandler) Validate(ctx context.Context,
	in *image_service.ValidateRequest) (res *image_service.ValidateResponse, err error) {
	if in.Image == nil {
		return nil, status.Error(codes.InvalidArgument, "image mustn't be empty")
	}
	defer h.handleError(&err)

	validation, err := h.s.Validate(ctx, &models.ValidateRequest{
		Img:            in.Image.Image,
		SupportedTypes: in.SupportedTypes,
		MaxWidth:       in.MaxWidth,
		MinWidth:       in.MinWidth,
		MaxHeight:      in.MaxHeight,
		MinHeight:      in.MinHeight,
	})
	if err != nil {
		return
	}
	res = &image_service.ValidateResponse{
		ImageValid: validation.IsImageValid,
		Details:    &validation.Details,
	}
	return
}

func (h *ImageProcessingServiceHandler) Blur(ctx context.Context,
	in *image_service.BlurRequest) (res *httpbody.HttpBody, err error) {
	if in.Image == nil {
		return nil, status.Error(codes.InvalidArgument, "image mustn't be empty")
	}
	defer h.handleError(&err)
	processed, err := h.s.Blur(ctx, in.Image.Image, in.BlurRadius, in.Method.String())
	if err != nil {
		return
	}

	res = &httpbody.HttpBody{
		ContentType: processed.ContentType,
		Data:        processed.Img,
	}
	return
}

func (h *ImageProcessingServiceHandler) Hue(ctx context.Context,
	in *image_service.HueRequest) (res *httpbody.HttpBody, err error) {
	if in.Image == nil {
		return nil, status.Error(codes.InvalidArgument, "image mustn't be empty")
	}
	defer h.handleError(&err)
	processed, err := h.s.Hue(ctx, in.Image.Image, in.Hue)
	if err != nil {
		return
	}

	res = &httpbody.HttpBody{
		ContentType: processed.ContentType,
		Data:        processed.Img,
	}
	return
}

func (h *ImageProcessingServiceHandler) Desaturate(ctx context.Context,
	in *image_service.Image) (res *httpbody.HttpBody, err error) {
	if in.Image == nil {
		return nil, status.Error(codes.InvalidArgument, "image mustn't be empty")
	}
	defer h.handleError(&err)
	processed, err := h.s.Desaturate(ctx, in.Image)
	if err != nil {
		return
	}

	res = &httpbody.HttpBody{
		ContentType: processed.ContentType,
		Data:        processed.Img,
	}
	return
}

func (h *ImageProcessingServiceHandler) handleError(err *error) {
	if err == nil || *err == nil {
		return
	}

	serviceErr := &models.ServiceError{}
	if errors.As(*err, &serviceErr) {
		*err = status.Error(convertServiceErrCodeToGrpc(serviceErr.Code), serviceErr.Msg)
	} else if _, ok := status.FromError(*err); !ok {
		e := *err
		*err = status.Error(codes.Unknown, e.Error())
	}
}

func convertServiceErrCodeToGrpc(code models.ErrorCode) codes.Code {
	switch code {
	case models.Internal:
		return codes.Internal
	case models.InvalidArgument:
		return codes.InvalidArgument
	case models.Unauthenticated:
		return codes.Unauthenticated
	case models.Conflict:
		return codes.AlreadyExists
	case models.NotFound:
		return codes.NotFound
	case models.Canceled:
		return codes.Canceled
	case models.DeadlineExceeded:
		return codes.DeadlineExceeded
	case models.PermissionDenied:
		return codes.PermissionDenied
	default:
		return codes.Unknown
	}
}
