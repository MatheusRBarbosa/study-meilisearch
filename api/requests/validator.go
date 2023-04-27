package requests

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/infra/db"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("exists", existsFunc)
}

func getValidator() *validator.Validate {
	return validate
}

type apiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func ValidateReqeust(ctx *fiber.Ctx, req interface{}) error {
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	if err := getValidator().Struct(req); err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		ctx.JSON(fiber.Map{"errors": parseError(err)})
		return err
	}

	return nil
}

func parseError(err error) []apiError {
	var validationErros validator.ValidationErrors
	var timeError *time.ParseError

	if errors.As(err, &validationErros) {
		out := make([]apiError, len(validationErros))
		for i, fe := range validationErros {
			out[i] = apiError{fe.Field(), parseErrorMessage(fe)}
		}
		return out
	} else if errors.As(err, &timeError) {
		out := make([]apiError, 1)
		out[0] = apiError{timeError.Value, fmt.Sprintf("Should have %s format", timeError.Layout)}

		return out
	}

	return nil
}

func parseErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "Este campo é obrigatório"
	case "email":
		return "E-mail inválido"
	case "gt":
		return fmt.Sprintf("Este campo deve ser maior que %s", fieldError.Param())
	case "gte":
		return fmt.Sprintf("Este campo deve ser maior ou igual a %s", fieldError.Param())
	case "lt":
		return fmt.Sprintf("Este campo deve ser menor que %s", fieldError.Param())
	case "lte":
		return fmt.Sprintf("Este campo deve ser menor ou igual a %s", fieldError.Param())
	case "exists":
		return fmt.Sprintf("%s não foi encontrado na base", fieldError.Value())
	}

	return fieldError.Error() // default error
}

// Custom validation functions
func existsFunc(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), "-")
	value := fl.Field().String()
	db := db.Context()

	var result int
	query := fmt.Sprintf("select count(id) from %s where %s = ?", params[0], params[1])
	db.Raw(query, value).Scan(&result)

	return result > 0
}
