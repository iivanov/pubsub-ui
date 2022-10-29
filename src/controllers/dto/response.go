package dto

type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Success:      true,
		Data:         data,
		ErrorMessage: "",
	}
}

func NewErrorResponse(errorMessage string) Response {
	return Response{
		Success:      false,
		Data:         nil,
		ErrorMessage: errorMessage,
	}
}
