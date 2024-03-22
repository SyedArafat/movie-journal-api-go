package dto

type SuccessResponseDto struct {
	Status     string `json:"status"`
	StatusCode int    `json:"code"`
	Data       any    `json:"data"`
	//Meta       any    `json:"meta"`
}

type ErrorResponseDto struct {
	Status       string `json:"status"`
	StatusCode   int    `json:"code"`
	ErrorMessage any    `json:"message"`
	ErrorDetails any    `json:"errors"`
}
