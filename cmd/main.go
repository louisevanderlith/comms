package main

import (
	"flag"
	"github.com/louisevanderlith/comms/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/comms/core"
)

func main() {
	scurty := flag.String("security", "http://localhost:8086", "Security Provider's URL")
	manager := flag.String("manager", "http://localhost:8097", "Manager Provider's URL")
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")
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
		Handler:      handles.SetupRoutes(*scurty, *manager, *srcSecrt, *smtpUser, *smtpPass, *smtpHost, *smtpPort),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
