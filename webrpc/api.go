package webrpc

import "github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"

type EchoReq struct {
	Name string
}

type EchoRes struct {
	Answer string
}

type MakeOrderReq struct {
	Token     string
	URL       string
	PageLimit int
	Mail      string
	DateTo    int64
	Frequency int
	UserID    int
}

type MakeOrderRes struct {
}

type ShowAllOderReq struct {
	Token  string
	UserID int
}
type ShowAllOderResp struct {
	Orders []entities.Order
}

type GetAdvertisementByOrderReq struct {
	Token   string
	OrderID int
}

type GetAdvertisementByOrderResp struct {
	Advertisements []entities.Advertisement
}
