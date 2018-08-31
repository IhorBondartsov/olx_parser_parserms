package storage

import "github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"

type Storage interface {
	CreateOrder(order entities.Order) (int, error)
	GetOrderByUserIDAndURL(uid int, url string) (entities.Order, error)
	GetOrdersByUserID(uid int) ([]entities.Order, error)
	GetOrderByID(oid int) (entities.Order, error)
	DeleteAllInformationForOrder(oid int) error

	CreateAdvertisement(a entities.Advertisement) error
	GetAdvertisementByOrderID(oid int) ([]entities.Advertisement, error)
}
