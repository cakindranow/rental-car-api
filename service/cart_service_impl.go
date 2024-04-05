package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/repository"
	"github.com/indrawanagung/food-order-api/util"
	"gorm.io/gorm"
)

type CartServiceImpl struct {
	CartRepository repository.CartRepositoryInterface
	Database       *gorm.DB
	Validate       *validator.Validate
}

func NewCartService(cartRepository repository.CartRepositoryInterface, database *gorm.DB, validate *validator.Validate) CartServiceInterface {
	return &CartServiceImpl{
		CartRepository: cartRepository,
		Database:       database,
		Validate:       validate,
	}
}

func (s CartServiceImpl) FindAll(userID string) []web.ListCartResponse {
	return util.ToCartProductResponses(s.CartRepository.FindAll(s.Database, userID))
}

func (s CartServiceImpl) FindByProductAndUserID(productID string, userID string) domain.Cart {
	cart, err := s.CartRepository.FindByProductAndUserID(s.Database, productID, userID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return cart

}

func (s CartServiceImpl) Save(request web.CartCreateRequest) {
	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	cart := domain.Cart{
		ProductID: request.ProductID,
		UserID:    request.UserID,
		Total:     request.Total,
	}
	s.CartRepository.Save(s.Database, cart)
}

func (s CartServiceImpl) Delete(productID string, userID string) {
	s.CartRepository.Delete(s.Database, productID, userID)
}
