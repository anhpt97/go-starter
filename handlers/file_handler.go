package handlers

import (
	"go-starter/config"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/render"
	"net/http"

	"golang.org/x/exp/slices"
)

type FileHandler struct{}

type IFileHandler interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

func NewFileHandler() IFileHandler {
	return &FileHandler{}
}

// @Tags    file
// @Summary Upload a file
// @Param   file                formData file   false " "
// @Param   locale              query    string false " " enums(en,vi)
// @Success 201                 object   render.Response{data=boolean}
// @Router  /api/v1/file/upload [POST]
func (handler *FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	_, fileHeader, err := r.FormFile("file")
	if err != nil {
		switch err {
		case http.ErrMissingBoundary:
			errors.BadRequestException(w, r, enums.Error.FileNotFound)
		case http.ErrMissingFile:
			errors.BadRequestException(w, r, enums.Error.FileNotFound)
		default:
			errors.BadRequestException(w, r, err)
		}
		return
	}

	if !slices.Contains(
		[]string{
			config.File.ContentType.Jpeg,
			config.File.ContentType.Png,
		}, fileHeader.Header["Content-Type"][0]) {
		errors.BadRequestException(w, r, enums.Error.InvalidFileFormat)
		return
	}

	if fileHeader.Size > config.File.MaxSize {
		errors.PayloadTooLargeException(w, r)
		return
	}

	// data := make([]byte, fileHeader.Size)
	// file.Read(data)
	// f, _ := os.Create("./" + fileHeader.Filename)
	// f.Write(data)

	render.WriteJSON(w, r, render.Response{
		Data: true,
	})
}
