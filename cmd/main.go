package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/alarm_clock"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/cfg"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/mailer"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/olx_client/client"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/olx_client/http_olx_client"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/storage/olxstorage"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/webrpc"
	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	"net/http"
)

var log = logrus.New()

func main() {
	log.Info("Make mailer")
	mailer := mailer.NewMailer(cfg.Mail.Port, cfg.Mail.Host, cfg.Mail.Mail, cfg.Mail.Password)
	log.Info("Make mysql connection")
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%s?timeout=5s",
		cfg.Storage.Login,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.DBName)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("[MAIN] Cant create db connection %v", err)
	}
	stor := olxstorage.NewStorage(db)
	log.Info("Make http client")
	httpClient := http_olx_client.NewOLXHTTPClient(http.DefaultClient)

	log.Info("Make alarm clock")
	to := make(chan int, cfg.SizeChanForAlarmClock)
	alarmC := alarm_clock.NewAlarmClock(to, cfg.SizeChanForResend)
	addChan := alarmC.GetAddChan()

	quitChan := make(chan struct{})

	log.Info("Make OLXClient")
	olxC := client.OLXClient{
		Storage:         stor,
		Mailer:          mailer,
		HTTPClient:      httpClient,
		AddToAlarmClock: addChan,
		RequestElem:     to,
		ResendQuit:      quitChan,
		NumberWorkers:   cfg.CountOLXClientWorkers,
	}

	log.Info("Start alarm clock")
	alarmC.Start()
	log.Info("Start OLXClient")
	olxC.Start()

	cfgAPI := webrpc.CfgAPI{
		AccessPublicKey: cfg.PublicKey,
		OLXClient:       &olxC,
	}
	webrpc.Start(cfgAPI)
	log.Info("Listening on ", (cfg.Route + ":" + cfg.Port))
	log.Panic(http.ListenAndServe((cfg.Route + ":" + cfg.Port), nil))
}
