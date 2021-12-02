package Middlewares

import (
	"Api.Calisma/src/Common/Models/ErrorModels"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"strings"
)

func ModelValidatorMiddleware(model interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Bind(model)
			validate := validator.New()
			err := validate.Struct(model)
			var ErrorText string
			if err != nil {
				english := en.New()
				uni := ut.New(english, english)
				trans, _ := uni.GetTranslator("en")
				_ = enTranslations.RegisterDefaultTranslations(validate, trans)
				errs := TranslateError(err, trans)
				ErrorText = ValidationErrors(errs).CombineAsString()
				panic(ErrorModels.InvalidModel.SetPublicDetail(ErrorText))
			}
			return next(c)
		}
	}
}

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

type ValidationErrors []error

func (errs ValidationErrors) CombineAsString() string {
	res := make([]string, len(errs))
	for i, e := range errs {
		res[i] = e.Error()
	}
	return strings.Join(res, ". ") + "."
}
