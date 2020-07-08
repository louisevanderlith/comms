package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(secureUrl, scrt, smtpUser, smtpPass, smtpHost string, smtpPort int) http.Handler {
	r := mux.NewRouter()

	view := kong.ResourceMiddleware("comms.messages.view", scrt, secureUrl, ViewMessage)
	r.HandleFunc("/message/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("comms.messages.create", scrt, secureUrl, CreateMessage(smtpUser, smtpPass, smtpHost, smtpPort))
	r.HandleFunc("/message", create).Methods(http.MethodPost)
	
	search := kong.ResourceMiddleware("comms.messages.search", scrt, secureUrl, SearchMessages)
	r.HandleFunc("/message/{pagesize:[A-Z][0-9]+}", search).Methods(http.MethodGet)
	r.HandleFunc("/message/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", search).Methods(http.MethodGet)

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "comms.messages.view", scrt)

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
