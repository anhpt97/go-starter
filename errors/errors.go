package errors

import (
	"errors"
	"go-starter/enums"
	"go-starter/i18n"
	"go-starter/response"
	"net/http"

	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func New(text string) error {
	return errors.New(text)
}

func getLanguageTag(r *http.Request) language.Tag {
	tag, err := language.Parse(r.URL.Query().Get(enums.Query.Locale))
	if err != nil {
		return language.English
	}
	return tag
}

func BadRequestException(w http.ResponseWriter, r *http.Request, err error, validationErrors ...response.ValidationError) {
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusBadRequest,
		Error: &response.Error{
			Code:    enums.ErrorCode(err.Error()),
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
			Details: validationErrors,
		},
	})
}

func UnauthorizedException(w http.ResponseWriter, r *http.Request, err error) {
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusUnauthorized,
		Error: &response.Error{
			Code:    enums.ErrorCode(err.Error()),
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func ForbiddenException(w http.ResponseWriter, r *http.Request, errs ...error) {
	err := enums.Error.PermissionDenied
	if len(errs) > 0 {
		err = enums.ErrorCode(errs[0].Error())
	}
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusForbidden,
		Error: &response.Error{
			Code:    err,
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func NotFoundException(w http.ResponseWriter, r *http.Request, errs ...error) {
	err := enums.Error.DataNotFound
	if len(errs) > 0 {
		err = enums.ErrorCode(errs[0].Error())
	}
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusNotFound,
		Error: &response.Error{
			Code:    err,
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func PayloadTooLargeException(w http.ResponseWriter, r *http.Request, errs ...error) {
	err := enums.Error.PayloadTooLarge
	if len(errs) > 0 {
		err = enums.ErrorCode(errs[0].Error())
	}
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusRequestEntityTooLarge,
		Error: &response.Error{
			Code:    err,
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func InternalServerErrorException(w http.ResponseWriter, r *http.Request, err error) {
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusInternalServerError,
		Error: &response.Error{
			Code:    enums.ErrorCode(err.Error()),
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func SqlError(w http.ResponseWriter, r *http.Request, err error, errs ...error) {
	switch err {
	case gorm.ErrRecordNotFound:
		if len(errs) > 0 {
			NotFoundException(w, r, errs[0])
		} else {
			NotFoundException(w, r)
		}
	default:
		InternalServerErrorException(w, r, err)
	}
}
