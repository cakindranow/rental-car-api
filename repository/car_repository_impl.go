package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
	"gorm.io/gorm"
)

type CarRepositoryImpl struct {
}

func NewCarRepository() CarRepositoryInterface {
	return &CarRepositoryImpl{}
}

func (p CarRepositoryImpl) Save(db *gorm.DB, car domain.Car) {
	err := db.Save(&car).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func (p CarRepositoryImpl) FindAll(db *gorm.DB, name string) []domain.Car {
	var cars []domain.Car

	if name != "" {
		err := db.Find(&cars, "lower(brand) like lower(?) or lower(model) like lower(?)", fmt.Sprint("%", name, "%"), "%"+name+"%").Error
		if err != nil {
			log.Error(err)
			panic(err)
		}
	} else {
		err := db.Find(&cars).Error
		if err != nil {
			log.Error(err)
			panic(err)
		}
	}

	return cars
}

func (p CarRepositoryImpl) FindById(db *gorm.DB, carId string) (domain.Car, error) {
	var car domain.Car

	err := db.Take(&car, "id = ?", carId).Error
	return car, err
}

func (p CarRepositoryImpl) CreateOrUpdateOrder(db *gorm.DB, order domain.Order) {
	err := db.Save(&order).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func (p CarRepositoryImpl) FindAllOrderByUserID(db *gorm.DB, userID string) []web.OrderResponse {
	var orders []web.OrderResponse

	err := db.Raw("SELECT o.id, o.status_id, o.cars_id, u.\"name\" as requested_name, s.status ,c.brand, c.plat ,c.model ,o.start_date, o.end_date, o.total_day, o.ordered_at, o.requested_by, o.note_admin, o.\"cost\"\nFROM public.orders o\ninner join users u on u.id = o.requested_by \ninner join status s on s.id = o.status_id \ninner join cars c on c.id  = o.cars_id where u.id = ? order by o.serial desc", userID).Scan(&orders).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return orders
}

func (p CarRepositoryImpl) FindAllOrderByAdmin(db *gorm.DB) []web.OrderResponse {
	var orders []web.OrderResponse

	err := db.Raw("SELECT o.id, o.status_id, o.cars_id, u.\"name\" as requested_name, s.status ,c.brand, c.plat ,c.model ,o.start_date, o.end_date, o.total_day, o.ordered_at, o.requested_by, o.note_admin, o.\"cost\"\nFROM public.orders o\ninner join users u on u.id = o.requested_by \ninner join status s on s.id = o.status_id \ninner join cars c on c.id  = o.cars_id order by o.serial desc\n").Scan(&orders).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return orders
}

func (p CarRepositoryImpl) FindOrderByID(db *gorm.DB, orderID string) (domain.Order, error) {
	var order domain.Order

	err := db.Take(&order, "id = ?", orderID).Error
	return order, err
}

func (p CarRepositoryImpl) CheckAvailableByOrderDate(db *gorm.DB, startDate string, endDate string, name string) []domain.Car {
	var carsNotAvailable []web.OrderResponse

	err := db.Raw("select * from public.orders o where o.status_id in ('1','2') and (start_date  <= ? and end_date  >= ?)", endDate, startDate).Scan(&carsNotAvailable).Error
	if err != nil {
		log.Error(err)
		fmt.Println(123)
		panic(err)
	}

	var carIDNotAvailable []string

	for _, order := range carsNotAvailable {
		carIDNotAvailable = append(carIDNotAvailable, order.CarsID)
	}

	if len(carIDNotAvailable) == 0 {
		carIDNotAvailable = append(carIDNotAvailable, "0")
	}
	var carsAvailable []domain.Car

	if name != "" {
		err = db.Raw("select * from cars where id not in (?) and (lower(brand) like lower(?) or lower(model) like lower(?))", carIDNotAvailable, fmt.Sprint("%", name, "%"), "%"+name+"%").Scan(&carsAvailable).Error
		if err != nil {
			log.Error()
			panic(err)
		}
	} else {
		err = db.Raw("select * from cars where id not in (?)", carIDNotAvailable).Scan(&carsAvailable).Error
		if err != nil {
			log.Error()
			panic(err)
		}
	}

	if err != nil {
		log.Error(err)
		panic(err)
	}

	return carsAvailable
}

func (p CarRepositoryImpl) CheckAvailableCar(db *gorm.DB, startDate string, endDate string, carID string) bool {
	var order web.OrderResponse

	err := db.Raw("select * from public.orders o where o.status_id in ('1','2') and (start_date  <= ? and end_date  >= ? and o.cars_id = ?)", endDate, startDate, carID).Scan(&order).Error
	if err != nil {
		log.Error(err)
	}

	if order.CarsID == carID {
		return false
	}
	return true
}
