package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/repository"
	"github.com/indrawanagung/food-order-api/util"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	Database          *gorm.DB
	Validate          *validator.Validate
	ProductRepository repository.ProductRepositoryInterface
}

func NewProductService(database *gorm.DB, validate *validator.Validate, productRepository repository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductServiceImpl{
		Database:          database,
		Validate:          validate,
		ProductRepository: productRepository,
	}
}

func (s ProductServiceImpl) FindAll() []web.ProductResponse {

	products := s.ProductRepository.FindAll(s.Database)

	productsResponse := util.ToListProductResponse(products)
	return productsResponse
}

func (s ProductServiceImpl) FindById(id string) web.ProductResponse {
	product, err := s.ProductRepository.FindById(s.Database, id)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	return util.ToProductResponse(product)
}

func (s ProductServiceImpl) Save(request web.CreateOrUpdateProduct) {
	product := domain.Product{
		ID:        uuid.NewString(),
		Name:      request.Name,
		Price:     request.Price,
		ImageUrl:  request.ImageUrl,
		DeletedAt: gorm.DeletedAt{},
	}
	s.ProductRepository.Save(s.Database, product)
}
