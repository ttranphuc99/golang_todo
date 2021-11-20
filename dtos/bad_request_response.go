package dtos

type BadRequestResponse struct {
	ErrorCode    string `json:errorCode`
	ErrorMessage string `json:errorMessage`
}
