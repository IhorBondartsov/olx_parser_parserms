package client

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/alarm_clock"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/mailer"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/olx_client/http_olx_client"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/storage"
)

var log = logrus.New()

type OLXClient struct {
	HTTPClient *http_olx_client.OlxHttpClient
	Storage    storage.Storage
	Mailer     mailer.MailSender

	ResendQuit      chan struct{}
	AddToAlarmClock chan alarm_clock.Item
	RequestElem     chan int
	NumberWorkers   int
}

// Start - started worker
func (c *OLXClient) Start() {
	log.Info("OLXClient started")
	for i := 0; i < c.NumberWorkers; i++ {
		go c.ResendRequestWorker()
	}
}

func (c *OLXClient) ResendRequestWorker() {
	for {
		select {
		case id := <-c.RequestElem:
			order, err := c.Storage.GetOrderByID(id)
			if err != nil {
				log.Errorf("[OLXClient][ResendRequestWorker] cant get order from db %v", err)
				go c.sendItemToAlarmClock(time.Now().Add(60*time.Second).Unix(), id)
				continue
			}
			err = c.GetAndSendAdvertisement(order)
			if err != nil {
				log.Errorf("[OLXClient][ResendRequestWorker] cant send Advertisement to user %v", err)
				go c.sendItemToAlarmClock(time.Now().Add(60*time.Second).Unix(), id)
				continue
			}

			if order.ExpirationTime < time.Now().Unix() {
				log.Infof("Order id %v is expired and deleted", id)
				c.Storage.DeleteAllInformationForOrder(id)
			} else {
				c.sendItemToAlarmClock(time.Now().Add(time.Duration(order.Frequency)*time.Second).Unix(), order.ID)
			}

		case <-c.ResendQuit:
			return
		}
	}
}

// AddNewOrder - add new order to storage, and make first request to OLX, collect all
// advertisements from OLX and saved them to database, also added to AlarmClock this order, and send to mail
func (c *OLXClient) AddNewOrder(order entities.Order) error {
	var err error
	if !c.hasOrder(order) {
		log.Error("[OLXClient][AddNewOrder]Duplicate order")
		return errors.New("Duplicate order")
	}

	order.ID, err = c.Storage.CreateOrder(order)
	if err != nil {
		log.Error("[OLXClient][AddNewOrder]Cant create order %v", err)
		return err
	}

	if err = c.GetAndSendAdvertisement(order); err != nil {
		log.Errorf("[OLXClient][GetAndSendAdvertisement] %v", err)
		return err
	}
	return nil
}

// GetAndSendAdvertisement - get all Advertisement from OLX and send them to user
func (c *OLXClient) GetAndSendAdvertisement(order entities.Order) error {
	advrtsmnts := c.HTTPClient.GetHTMLPages(order.URL, order.PageLimit)
	for _, v := range advrtsmnts {
		v.OrderID = order.ID
		err := c.Storage.CreateAdvertisement(v)
		if err != nil {
			log.Errorf("[OLXClient][AddNewOrder] Database error %v", err)
		}
	}
	go c.sendItemToAlarmClock(time.Now().Add(time.Duration(order.Frequency)*time.Second).Unix(), order.ID)
	return c.Mailer.SendMail(advrtsmnts, order.Mail)
}

func (c *OLXClient) sendItemToAlarmClock(time int64, id int) {
	item := alarm_clock.Item{
		Time: time,
		Id:   id,
	}
	c.AddToAlarmClock <- item
}

func (c *OLXClient) hasOrder(order entities.Order) bool {
	_, err := c.Storage.GetOrderByUserIDAndURL(order.UserID, order.URL)
	log.Infof("[OLXClient][hasOrder] err == sql.ErrNoRows %v", err == sql.ErrNoRows)
	if err == sql.ErrNoRows {
		return true
	}
	return false
}

func (c *OLXClient) FiltredAdvertisment(orderID int, advrtsmnts []entities.Advertisement) ([]entities.Advertisement, error) {
	savedAdvrtsmnts, err := c.Storage.GetAdvertisementByOrderID(orderID)
	if err == sql.ErrNoRows {
		log.Errorf("[OLXClient][FiltredAdvertisment] Database error %v", err)
		return advrtsmnts, nil
	}
	if err != nil {
		log.Errorf("[OLXClient][FiltredAdvertisment] Database error %v", err)
		return nil, err
	}

	var newAdvrtsmnts []entities.Advertisement
	for _, v := range advrtsmnts {
		has := false
		for _, sv := range savedAdvrtsmnts {
			if sv.URL == v.URL {
				has = true
				break
			}
		}
		if !has {
			newAdvrtsmnts = append(newAdvrtsmnts, v)
		}
	}
	return newAdvrtsmnts, nil
}
