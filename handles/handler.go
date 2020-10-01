package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong/middle"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(securityUrl, managerUrl, scrt, smtpUser, smtpPass, smtpHost string, smtpPort int) http.Handler {
	r := mux.NewRouter()
	ins := middle.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
	view := ins.Middleware("comms.messages.view", scrt, ViewMessage)
	r.HandleFunc("/message/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	create := ins.Middleware("comms.messages.create", scrt, CreateMessage(smtpUser, smtpPass, smtpHost, smtpPort))
	r.HandleFunc("/message", create).Methods(http.MethodPost)

	search := ins.Middleware("comms.messages.search", scrt, SearchMessages)
	r.HandleFunc("/message/{pagesize:[A-Z][0-9]+}", search).Methods(http.MethodGet)
	r.HandleFunc("/message/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", search).Methods(http.MethodGet)

	lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "comms.messages.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
