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

type CarServiceImpl struct {
	Database       *gorm.DB
	Validate       *validator.Validate
	CarRepository  repository.CarRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func NewCarService(database *gorm.DB, validate *validator.Validate, carRepository repository.CarRepositoryInterface, userRepository repository.UserRepositoryInterface) CarServiceInterface {
	return &CarServiceImpl{
		Database:       database,
		Validate:       validate,
		CarRepository:  carRepository,
		UserRepository: userRepository,
	}
}

func (s CarServiceImpl) FindAll(startDate string, endDate string, name string) []web.CarResponse {
	var carsResponse []web.CarResponse
	if startDate != "" || endDate != "" {
		cars := s.CarRepository.CheckAvailableByOrderDate(s.Database, startDate, endDate, name)
		carsResponse = util.ToListCarResponse(cars)
	} else {
		cars := s.CarRepository.FindAll(s.Database, name)
		carsResponse = util.ToListCarResponse(cars)
	}
	return carsResponse
}

func (s CarServiceImpl) FindById(id string) web.CarResponse {
	product, err := s.CarRepository.FindById(s.Database, id)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	return util.ToCarResponse(product)
}

func (s CarServiceImpl) Save(request web.CreateOrUpdateCar) {
	product := domain.Car{
		ID:              uuid.NewString(),
		Brand:           request.Brand,
		Model:           request.Model,
		Plat:            request.Plat,
		DailyRentalRate: request.DailyRentalRate,
		ImageUrl:        request.ImageUrl,
		Desc:            request.Desc,
		DeletedAt:       gorm.DeletedAt{},
	}
	s.CarRepository.Save(s.Database, product)
}

func (s CarServiceImpl) CreateOrder(request web.CreateOrderRequest, userID string) {
	car, err := s.CarRepository.FindById(s.Database, request.CarsID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	isAvailableCar := s.CarRepository.CheckAvailableCar(s.Database, request.StartDate, request.EndDate, request.CarsID)
	if !isAvailableCar {
		panic(exception.NewNotFoundError("cars is not available"))
	}

	order := domain.Order{
		ID:          uuid.NewString(),
		StatusID:    util.StatusIDOnProcess(),
		CarsID:      request.CarsID,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		TotalDay:    request.TotalDay,
		OrderedAt:   util.GetUnixTimestamp(),
		RequestedBy: userID,
		Cost:        car.DailyRentalRate * int64(request.TotalDay),
	}
	s.CarRepository.CreateOrUpdateOrder(s.Database, order)
}

func (s CarServiceImpl) FindAllOrder(userID string) []web.OrderResponse {
	var orders []web.OrderResponse

	err, user := s.UserRepository.FindByID(s.Database, userID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}
	if user.IsAdmin {
		orders = s.CarRepository.FindAllOrderByAdmin(s.Database)
	} else {
		orders = s.CarRepository.FindAllOrderByUserID(s.Database, userID)
	}
	return orders
}

func (s CarServiceImpl) CanceledOrderByUserID(userID string, orderID string) {
	order, err := s.CarRepository.FindOrderByID(s.Database, orderID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	if order.RequestedBy != userID {
		log.Error(err)
		panic(exception.NewForbiddenError("forbidden access"))
	}

	order.StatusID = util.StatusIDCanceled()

	s.CarRepository.CreateOrUpdateOrder(s.Database, order)
}

func (s CarServiceImpl) ReturnedCarOrderByUserID(userID string, orderID string, plat string) {
	order, err := s.CarRepository.FindOrderByID(s.Database, orderID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	if order.RequestedBy != userID {
		log.Error(err)
		panic(exception.NewForbiddenError("forbidden access"))
	}

	car, err := s.CarRepository.FindById(s.Database, order.CarsID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	if car.Plat != plat {
		log.Error(err)
		panic(exception.NewBadRequestError("plat number is not valid"))
	}

	order.StatusID = util.StatusIDReturned()

	s.CarRepository.CreateOrUpdateOrder(s.Database, order)
}

func (s CarServiceImpl) ApproveCarOrderByAdmin(userID string, orderID string) {
	order, err := s.CarRepository.FindOrderByID(s.Database, orderID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	err, user := s.UserRepository.FindByID(s.Database, userID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	if !user.IsAdmin {
		log.Error(err)
		panic(exception.NewForbiddenError("user is not administrator"))
	}

	order.StatusID = util.StatusIDApproved()

	s.CarRepository.CreateOrUpdateOrder(s.Database, order)
}

func (s CarServiceImpl) RejectCarOrderByAdmin(userID string, orderID string, noteAdmin string) {
	order, err := s.CarRepository.FindOrderByID(s.Database, orderID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	err, user := s.UserRepository.FindByID(s.Database, userID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	if !user.IsAdmin {
		log.Error(err)
		panic(exception.NewForbiddenError("user is not administrator"))
	}

	order.StatusID = util.StatusIDRejected()
	order.NoteAdmin = noteAdmin

	s.CarRepository.CreateOrUpdateOrder(s.Database, order)
}
