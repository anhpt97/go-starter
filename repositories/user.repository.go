package repositories

import (
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/lib"
	"net/http"
)

type UserRepository struct {
	db lib.Db
}

type IUserRepository interface {
	FindOne(w http.ResponseWriter, r *http.Request, conditions entities.User) (user entities.User, ok bool)
}

func NewUserRepository(db lib.Db) IUserRepository {
	return &UserRepository{
		db,
	}
}

func (repository *UserRepository) FindOne(w http.ResponseWriter, r *http.Request, conditions entities.User) (user entities.User, ok bool) {
	err := repository.db.Model(user).
		Where(conditions).
		Take(&user).Error
	if err != nil {
		errors.SqlError(w, r, err, enums.Error.UserNotFound)
		return
	}

	switch user.Status {
	case enums.User.Status.NotActivated:
		errors.UnauthorizedException(w, r, enums.Error.NonActivatedAccount)
		return
	case enums.User.Status.IsDisabled:
		errors.ForbiddenException(w, r, enums.Error.DisabledAccount)
		return
	}
	return user, true
}
