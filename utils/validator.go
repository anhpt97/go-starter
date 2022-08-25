package utils

import (
	"encoding/json"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/response"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/maps"
)

func ValidateRequestBody(w http.ResponseWriter, r *http.Request, bodyStruct any) ([]string, error) {
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, bodyStruct)

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.Struct(bodyStruct); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		err := make([]response.ValidationError, len(validationErrors))

		for i, fieldError := range validationErrors {
			err[i] = response.ValidationError{
				Field: strings.SplitN(fieldError.Namespace(), ".", 2)[1],
				Message: func(fe validator.FieldError) string {
					switch fe.Tag() {
					case "required":
						return "This field is required"
					case "max":
						return "Max length: " + fe.Param()
					default:
						return strings.Split(fe.Error(), "Error:")[1]
					}
				}(fieldError),
			}
		}

		errors.BadRequestException(w, r, enums.Error.InvalidInputData, err...)
		return nil, validationErrors
	}

	bodyMap := map[string]any{}
	json.Unmarshal(body, &bodyMap)
	return maps.Keys(bodyMap), nil
}
