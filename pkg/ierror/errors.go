package ierror

import "net/http"

type (
	ICustomType         string
	ICustomErrorOptions struct {
		Message    string
		StatusCode int
	}
)

type ResponseError struct {
	Type       string `json:"type"`
	Msg        string `json:"msg"`
	StatusCode int    `json:"-"`
}

const (
	E_BAD_REQUEST      ICustomType = "BAD_REQUEST"
	E_VALIDATION_ERROR ICustomType = "VALIDATION_ERROR"
	E_NOT_FOUND        ICustomType = "NOT_FOUND"
	E_INTERNAL_SERVER  ICustomType = "INTERNAL_SERVER"
)

var (
	ErrorMessages = map[ICustomType]ICustomErrorOptions{
		E_BAD_REQUEST:      {Message: "Неправильный запрос", StatusCode: http.StatusBadRequest},
		E_INTERNAL_SERVER:  {Message: "Внутренняя ошибка сервера", StatusCode: http.StatusInternalServerError},
		E_VALIDATION_ERROR: {Message: "Ошибка валидации", StatusCode: http.StatusBadRequest},
		E_NOT_FOUND:        {Message: "Не найдено", StatusCode: http.StatusNotFound},
	}
)

func New(ierrorType ICustomType) *ResponseError {
	return &ResponseError{
		Type:       string(ierrorType),
		Msg:        ErrorMessages[ierrorType].Message,
		StatusCode: ErrorMessages[ierrorType].StatusCode,
	}
}
