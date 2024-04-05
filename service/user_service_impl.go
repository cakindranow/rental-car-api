package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/repository"
	"github.com/indrawanagung/food-order-api/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Database       *gorm.DB
	UserRepository repository.UserRepositoryInterface
	Validate       *validator.Validate
}

func NewUserService(database *gorm.DB, userRepository repository.UserRepositoryInterface, validate *validator.Validate) UserServiceInterface {
	return &UserServiceImpl{
		Database:       database,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s UserServiceImpl) Save(request web.UserCreateOrUpdateRequest) string {
	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	id := util.GenerateUUID()

	err, _ = s.UserRepository.FindByEmail(s.Database, request.Email)

	if err == nil {
		panic(exception.NewBadRequestError("email address has been already exist"))
	}
	err = s.UserRepository.SaveOrUpdate(s.Database, domain.User{
		ID:       id,
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passwordHash),
		Address:  request.Address,
		Phone:    request.Phone,
		SIM:      request.SIM,
	})

	if err != nil {
		log.Error(err)
		panic(err)
	}
	return id
}

func (s UserServiceImpl) Update(ID string, request web.UserCreateOrUpdateRequest) {

	err, user := s.UserRepository.FindByID(s.Database, ID)
	if err != nil {
		panic(exception.NewNotFoundError(fmt.Sprintf("user id %s is not found", ID)))
	}

	var passwordHash string
	if request.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err)
		}
		passwordHash = string(hash)
	} else {
		passwordHash = user.Password
		request.Password = util.GetUnixTimestamp() // random string, just for pass validate password request
	}

	err = s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	err, _ = s.UserRepository.FindByEmail(s.Database, request.Email)

	if err == nil && request.Email != user.Email {
		panic(exception.NewBadRequestError("email address has been already exist"))
	}

	err = s.UserRepository.SaveOrUpdate(s.Database, domain.User{
		ID:       ID,
		Name:     request.Name,
		Email:    request.Email,
		Password: passwordHash,
		Address:  request.Address,
		Phone:    request.Phone,
		SIM:      request.SIM,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (s UserServiceImpl) FindByID(ID string) web.UserResponse {
	err, user := s.UserRepository.FindByID(s.Database, ID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(fmt.Sprintf("user id %s is not found", ID)))
	}

	return web.ToUserResponse(user)
}

func (s UserServiceImpl) FindAll() []domain.User {
	users := s.UserRepository.FindAll(s.Database)
	return users
}
