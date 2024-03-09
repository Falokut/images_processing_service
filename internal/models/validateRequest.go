package models

type ValidateRequest struct {
	Img            []byte
	SupportedTypes []string

	MaxWidth  *int32
	MaxHeight *int32
	MinWidth  *int32
	MinHeight *int32
}
