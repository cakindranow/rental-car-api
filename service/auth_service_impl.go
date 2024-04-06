package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/repository"
	"github.com/indrawanagung/food-order-api/util"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepositoryInterface
	Database       *gorm.DB
	Validate       *validator.Validate
}

func NewAuthService(userRepository repository.UserRepositoryInterface, database *gorm.DB, validate *validator.Validate) AuthServiceInterface {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		Database:       database,
		Validate:       validate,
	}
}
func (s AuthServiceImpl) Login(request web.LoginRequest) web.LoginResponse {
	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}
	err, user := s.UserRepository.FindByEmail(s.Database, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//compare password has
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	return web.LoginResponse{
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		Token:   t,
	}
}
