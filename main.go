package main

import (
	"log"
	"os"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/comms/routers"

	"github.com/astaxie/beego"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	name := beego.BConfig.AppName
	srv := mango.NewService(name, pubPath, enums.API)

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
