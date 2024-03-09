package models

type ValidateResponse struct {
	IsImageValid bool
	// If IsImageValid is false will contains message about failed check
	Details string
}
