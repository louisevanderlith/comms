package main

import (
	"flag"
	"github.com/louisevanderlith/comms/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/comms/core"
)

func main() {
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	audience := flag.String("audience", "comms", "Token target 'aud'")
	smtpUser := flag.String("smtpUsername", "frikkie@mango.avo", "User used to authenticate SMTP calls")
	smtpPass := flag.String("smtpPassword", "not_real_password", "Password used to authenticate SMTP calls")
	smtpHost := flag.String("smtpHost", "41.0.0.0", "Host used for SMTP connections")
	smtpPort := flag.Int("smtpPort", 587, "Port used to connect to SMTP")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8085",
		Handler:      handles.SetupRoutes(*issuer, *audience, *smtpUser, *smtpPass, *smtpHost, *smtpPort),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
