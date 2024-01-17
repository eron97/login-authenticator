package request_validate

import (
	"errors"
	"net/http"

	"github.com/eron97/login-authenticator.git/src/config/exceptions"
	"github.com/eron97/login-authenticator.git/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateRequest(c *gin.Context, request *models.CreateUser) error {
	if err := c.ShouldBindJSON(request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessage := exceptions.NewBadRequestError("Erro de validação")

			for _, fieldError := range ve {
				cause := exceptions.CausesMessages{
					Field:   fieldError.Field(),
					Message: fieldError.Translate(transl),
				}
				errorMessage.Causes = append(errorMessage.Causes, cause)
			}

			c.JSON(http.StatusBadRequest, gin.H{"Erro de validação": errorMessage})
			return errors.New(errorMessage.Message)
		}

		errorMessage := exceptions.NewUnmarshalError("Campos preenchidos não correspondem aos tipos atribuídos")
		c.JSON(http.StatusBadRequest, errorMessage)
		return errors.New(errorMessage.Message)
	}

	return nil
}
