package errors

import (
	"errors"
	"go-starter/enums"
	"go-starter/i18n"
	"go-starter/render"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/language"
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

func BadRequestException(w http.ResponseWriter, r *http.Request, err error, errors ...render.ValidationError) {
	render.WriteJSON(w, r, render.Response{
		StatusCode: http.StatusBadRequest,
		Error: &render.Error{
			Code:    enums.ErrorCode(err.Error()),
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
			Details: errors,
		},
	})
}

func UnauthorizedException(w http.ResponseWriter, r *http.Request, err error) {
	render.WriteJSON(w, r, render.Response{
		StatusCode: http.StatusUnauthorized,
		Error: &render.Error{
			Code:    enums.ErrorCode(err.Error()),
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func ForbiddenException(w http.ResponseWriter, r *http.Request, errors ...error) {
	err := enums.Error.PermissionDenied
	if len(errors) > 0 {
		err = enums.ErrorCode(errors[0].Error())
	}
	render.WriteJSON(w, r, render.Response{
		StatusCode: http.StatusForbidden,
		Error: &render.Error{
			Code:    err,
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func NotFoundException(w http.ResponseWriter, r *http.Request, errors ...error) {
	err := enums.Error.DataNotFound
	if len(errors) > 0 {
		err = enums.ErrorCode(errors[0].Error())
	}
	render.WriteJSON(w, r, render.Response{
		StatusCode: http.StatusNotFound,
		Error: &render.Error{
			Code:    err,
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func PayloadTooLargeException(w http.ResponseWriter, r *http.Request, errors ...error) {
	err := enums.Error.PayloadTooLarge
	if len(errors) > 0 {
		err = enums.ErrorCode(errors[0].Error())
	}
	render.WriteJSON(w, r, render.Response{
		StatusCode: http.StatusRequestEntityTooLarge,
		Error: &render.Error{
			Code:    err,
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func InternalServerErrorException(w http.ResponseWriter, r *http.Request, err error) {
	render.WriteJSON(w, r, render.Response{
		StatusCode: http.StatusInternalServerError,
		Error: &render.Error{
			Code:    enums.ErrorCode(err.Error()),
			Message: i18n.Translate(err.Error(), getLanguageTag(r)),
		},
	})
}

func SqlError(w http.ResponseWriter, r *http.Request, err error, errors ...error) {
	switch err {
	case mongo.ErrNoDocuments:
		if len(errors) > 0 {
			NotFoundException(w, r, errors[0])
		} else {
			NotFoundException(w, r)
		}
	default:
		InternalServerErrorException(w, r, err)
	}
}
