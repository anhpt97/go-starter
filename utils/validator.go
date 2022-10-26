package utils

import (
	"encoding/json"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/render"
	"io"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

func ValidateRequestBody(w http.ResponseWriter, r *http.Request, dto any) (data map[string]any, err error) {
	b, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(b, &dto)
	if err != nil && strings.Contains(err.Error(), "parsing time") {
		errors.BadRequestException(w, r, enums.Error.InvalidTimeValue)
		return
	}

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err = validate.Struct(dto); err != nil {
		fieldErrors := err.(validator.ValidationErrors)
		validationErrors := make([]render.ValidationError, len(fieldErrors))

		for i, fieldError := range fieldErrors {
			validationErrors[i] = render.ValidationError{
				Field: fieldError.Field(),
				Message: func(fe validator.FieldError) string {
					switch fe.Tag() {
					case "required":
						return strings.Join([]string{fe.Field(), "is required"}, " ")
					case "max":
						return strings.Join([]string{fe.Field(), "cannot be longer than", fe.Param(), "character(s)"}, " ")
					case "gte":
						return strings.Join([]string{fe.Field(), "cannot be less than", fe.Param()}, " ")
					default:
						return strings.Split(fe.Error(), "Error:")[1]
					}
				}(fieldError),
			}
		}

		errors.BadRequestException(w, r, enums.Error.InvalidInputData, validationErrors...)
		return
	}

	bodyM := map[string]any{}
	dtoM := map[string]any{}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		json.Unmarshal(b, &bodyM)
	}()

	go func() {
		defer wg.Done()

		b, _ := json.Marshal(dto)
		json.Unmarshal(b, &dtoM)
	}()

	wg.Wait()

	return lo.PickByKeys(bodyM, maps.Keys(dtoM)), nil
}
