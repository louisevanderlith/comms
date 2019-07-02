package main

import (
	"log"
	"os"
	"path"

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
		routers.Setup(srv, host)
		showSMTPInfo()
		beego.Run()
	}
}

func showSMTPInfo() {
	smtpUser := os.Getenv("SMTPUsername")
	smtpAddress := os.Getenv("SMTPAddress")
	smtpPort := os.Getenv("SMTPPort")

	log.Print(smtpUser, smtpAddress, smtpPort)
}
