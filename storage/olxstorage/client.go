package olxstorage

import (
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/jmoiron/sqlx"
)

func NewStorage(db *sqlx.DB) *parserStorage {
	return &parserStorage{db: db}
}

type parserStorage struct {
	db *sqlx.DB
}

const (
	createOrderStmt = `
				INSERT INTO
					orderOLX
				SET
					user_id = :user_id,
					url = :url,
					page_limit = :page_limit,
					mail = :mail,
					expiration_time = :expiration_time,
					frequency = :frequency;
`
	deleteAdvertisementsStmtByID = `
				DELETE
				FROM
					advertisements
				WHERE
					order_id = ?;
`
	deleteOrderStmtByID = `
				DELETE
				FROM
					user
				WHERE
					id = ?;
`
	getOrderByUserIDAndURLStmt = `
				SELECT *
				FROM orderOLX WHERE
					user_id = ?
				AND	url = ?;
`
	getOrderByIDStmt = `
				SELECT *
				FROM orderOLX WHERE
					id = ?;
`
	getOrderByUserIDStmt = `
				SELECT *
				FROM orderOLX WHERE
					user_id = ?;
`
	// table advertisements
	createAdvertisementsStmt = `
				INSERT INTO
					advertisements
				SET
					order_id = :order_id,
					title = :title,
					url = :url,
					created_at = :created_at;
`

	getAdvertismentByOrderIDStmt = `
				SELECT *
				FROM advertisements WHERE
					order_id = ?;
`

)

func (c *parserStorage) CreateOrder(order entities.Order) (int, error) {
	res, err := c.db.NamedExec(createOrderStmt, order)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (c *parserStorage) CreateAdvertisement(a entities.Advertisement) error {
	_, err := c.db.NamedExec(createAdvertisementsStmt, a)
	return err
}
func (c *parserStorage) GetAdvertisementByOrderID(oid int) ([]entities.Advertisement, error) {
	var advs []entities.Advertisement
	err := c.db.Get(&advs, getAdvertismentByOrderIDStmt, oid)
	return advs, err
}

func (c *parserStorage) GetOrderByUserIDAndURL(uid int, url string) (entities.Order, error) {
	var order entities.Order
	err := c.db.Get(&order, getOrderByUserIDAndURLStmt, uid, url)
	return order, err
}

func (c *parserStorage) GetOrderByID(oid int) (entities.Order, error){
	var order entities.Order
	err := c.db.Get(&order, getOrderByIDStmt, oid)
	return order, err
}

func (c *parserStorage) GetOrdersByUserID(uid int) ([]entities.Order, error){
	var order []entities.Order
	err := c.db.Get(&order, getOrderByUserIDStmt, uid)
	return order, err
}

func (c *parserStorage)DeleteAllInformationForOrder(oid int) error {
	_, err := c.db.Query(deleteAdvertisementsStmtByID, oid)
	if err != nil{
		return err
	}
	_, err = c.db.Query(deleteOrderStmtByID, oid)
	return  err
}