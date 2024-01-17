package exceptions

import "net/http"

type ErrorMessages struct {
	Message string           `json:"message" example:"error trying to process request"`
	Err     string           `json:"error" example:"internal_server_error"`
	Code    int              `json:"code" example:"500"`
	Causes  []CausesMessages `json:"causes"`
}

type CausesMessages struct {
	Field   string `json:"field" example:"name"`
	Message string `json:"message" example:"name is required"`
}

type ErrorUnmarshal struct {
	Message         string
	Err             string
	Code            int
	CausesUnmarshal CausesUnmarshal
}

type CausesUnmarshal struct {
	First_Name string
	Last_Name  string
	Email      string
	Password   string
}

/*
	Go possui uma interface predefinida `error` que contém apenas um método
	chamado `Error() string` e quando uma estrutura o implementa então ela
	se torna compátivel com a interface.

	Na estrutura abaixo podemos ver que ela é implementada por struct ErrorMessages{}
*/

func (r *ErrorMessages) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []CausesMessages) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewUnmarshalError(message string) *ErrorUnmarshal {
	var causes = CausesUnmarshal{
		First_Name: "Type string",
		Last_Name:  "Type string",
		Email:      "Type string",
		Password:   "Type int8 (age < 128)",
	}

	return &ErrorUnmarshal{
		Message:         message,
		Err:             "unmarshal_error",
		Code:            http.StatusBadRequest,
		CausesUnmarshal: causes,
	}
}
