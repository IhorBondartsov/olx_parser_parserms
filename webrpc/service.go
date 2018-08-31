package webrpc

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/IhorBondartsov/OLX_Parser/lib/jwtLib"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"github.com/Sirupsen/logrus"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/olx_client/client"
	"github.com/go-errors/errors"
)

var log = logrus.New()

func Start(cfg CfgAPI) {
	// Server export an object of type ExampleSvc.
	if err := rpc.Register(NewAPI(cfg)); err != nil {
		log.Panic(err)
	}

	// Server provide a HTTP transport on /rpc endpoint.
	http.Handle("/rpc", jsonrpc2.HTTPHandler(nil))

}

func NewAPI(cfg CfgAPI) *API {
	atp, err := jwtLib.NewJWTParser(cfg.AccessPublicKey)
	if err != nil {
		log.Errorf("Cant create AccessTokenParser. Err %v", err)
		return nil
	}
	return &API{
		AccessTokenParser: atp,
		OLXClient:cfg.OLXClient,
	}
}

type CfgAPI struct {
	AccessPublicKey []byte
	OLXClient *client.OLXClient
}

type API struct {
	AccessTokenParser jwtLib.JWTParser
	OLXClient *client.OLXClient
}

// Echo method for checking service
func (a *API) Echo(req EchoReq, res *EchoRes) error {
	fmt.Println("I called")
	res.Answer = fmt.Sprintf("Hello %s!!!", req.Name)
	return nil
}

// MakeOrder - make row to db
func (a *API) MakeOrder(req MakeOrderReq, res *MakeOrderRes) error {
	_, err := a.AccessTokenParser.Parse(req.Token)
	if err != nil {
		return err
	}

	order := entities.Order{
		Mail: req.Mail,
		ExpirationTime: req.DateTo,
		Frequency:      req.Frequency,
		PageLimit:      req.PageLimit,
		URL:            req.URL,
		UserID:         req.UserID,
	}

	return a.OLXClient.AddNewOrder(order)
}

func (a *API) ShowAllOder(req ShowAllOderReq, res *ShowAllOderResp) error {
	_, err := a.AccessTokenParser.Parse(req.Token)
	if err != nil {
		return err
	}
	res.Orders, err = a.OLXClient.Storage.GetOrdersByUserID(req.UserID)
	return err
}

func (a *API) GetAdvertisementByOrder(req GetAdvertisementByOrderReq, res *GetAdvertisementByOrderResp) error {
	c, err := a.AccessTokenParser.Parse(req.Token)
	if err != nil {
		return err
	}
	order, err := a.OLXClient.Storage.GetOrderByID(req.OrderID)
	if err != nil{ return err}
	if c.ID != fmt.Sprintf("%d", order.UserID){
		return errors.New("Forbidden user id not your")
	}
	res.Advertisements, err = a.OLXClient.Storage.GetAdvertisementByOrderID(req.OrderID)
	return err
}