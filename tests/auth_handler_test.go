package tests

import (
	"bytes"
	"encoding/json"
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/handlers"
	"go-starter/render"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockUserRepository struct{}

func (repository *MockUserRepository) FindAndCount(w http.ResponseWriter, r *http.Request) (books []*entities.User, total int64, err error) {
	return
}

func (repository *MockUserRepository) FindOne(w http.ResponseWriter, r *http.Request, conditions any) (user *entities.User, err error) {
	m := map[string][]map[string]any{}
	b, _ := json.Marshal(conditions)
	json.Unmarshal(b, &m)

	if m["$or"][0]["username"] != "superadmin" {
		err = mongo.ErrNoDocuments
		errors.SqlError(w, r, err, enums.Error.UserNotFound)
		return
	}
	return &entities.User{
		ID:             primitive.NewObjectID(),
		Username:       "superadmin",
		HashedPassword: "$2a$10$XajjQvNhvvRt5GSeFk1xFeyqRrsxkhBkUiQeg0dt.wU1qD4aFDcga",
		Role:           enums.User.Role.Admin,
	}, nil
}

func NewMockAuthHandler() handlers.IAuthHandler {
	return handlers.NewAuthHandler(&MockUserRepository{}, env)
}

func TestLogin(t *testing.T) {
	mockAuthHandler := NewMockAuthHandler()

	testCases := []struct {
		input         dto.LoginDto
		expectedError render.Error
	}{
		{
			input: dto.LoginDto{
				Username: "username",
				Password: "",
			},
			expectedError: render.Error{Code: enums.Error.InvalidInputData},
		},
		{
			input: dto.LoginDto{
				Username: "username",
				Password: "123456",
			},
			expectedError: render.Error{Code: enums.Error.UserNotFound},
		},
		{
			input: dto.LoginDto{
				Username: "superadmin",
				Password: "123456",
			},
			expectedError: render.Error{Code: enums.Error.InvalidPassword},
		},
		{
			input: dto.LoginDto{
				Username: "superadmin",
				Password: "allmine",
			},
		},
	}

	for _, tc := range testCases {
		b, _ := json.Marshal(tc.input)
		r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(b))

		t.Run("", func(t *testing.T) {
			mockAuthHandler.Login(w, r)

			b, _ := io.ReadAll(w.Body)
			response := render.Response{}
			json.Unmarshal(b, &response)

			if response.Error != nil && response.Error.Code != tc.expectedError.Code {
				t.Errorf("Error code mismatch: got %v, want %v", tc.expectedError.Code, response.Error.Code)
			}
		})
	}
}
