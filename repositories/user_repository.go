package repositories

import (
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/lib"
	"go-starter/utils"
	"net/http"
	"sync"

	"gorm.io/gorm"
)

type UserRepository struct {
	db lib.Db
}

type IUserRepository interface {
	FindAndCount(w http.ResponseWriter, r *http.Request) (users []*entities.User, total int64, err error)
	FindOne(w http.ResponseWriter, r *http.Request, conditions any) (user *entities.User, err error)
}

func NewUserRepository(db lib.Db) IUserRepository {
	return &UserRepository{
		db,
	}
}

func (repository *UserRepository) FindAndCount(w http.ResponseWriter, r *http.Request) (users []*entities.User, total int64, err error) {
	pagination := utils.Pagination(r)

	q := repository.db.Model(&users).
		Preload("Books", "TRUE ORDER BY book.id DESC")

	ch := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Limit(pagination.Limit).
			Offset(pagination.Offset).
			Order(pagination.Order).
			Find(&users).Error
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Count(&total).Error
		if err != nil {
			ch <- err
		}
	}()

	wg.Wait()
	close(ch)

	for err = range ch {
		if err != nil {
			errors.SqlError(w, r, err)
			return
		}
	}
	return
}

func (repository *UserRepository) FindOne(w http.ResponseWriter, r *http.Request, conditions any) (user *entities.User, err error) {
	err = repository.db.Model(&user).
		Where(conditions).
		Take(&user).Error
	if err != nil {
		errors.SqlError(w, r, err, enums.Error.UserNotFound)
		return
	}

	switch user.Status {
	case enums.User.Status.NotActivated:
		err = enums.Error.NonActivatedAccount
		errors.UnauthorizedException(w, r, err)
	case enums.User.Status.IsDisabled:
		err = enums.Error.DisabledAccount
		errors.ForbiddenException(w, r, err)
	}
	return
}
