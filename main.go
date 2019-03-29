package main

import (
	"log"
	"os"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	_ "github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/comms/routers"

	"github.com/astaxie/beego"
)

func main() {
	mode := os.Getenv("RUNMODE")

	// Register with router
	name := beego.BConfig.AppName
	srv := mango.NewService(mode, name, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		showSMTPInfo()
		beego.Run()
	}
}

func showSMTPInfo() {
	smtpUser := beego.AppConfig.String("smtpUsername")
	smtpAddress := beego.AppConfig.String("smtpAddress")
	smtpPort := beego.AppConfig.String("smtpPort")

	log.Print(smtpUser, smtpAddress, smtpPort)
}
