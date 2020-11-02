package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience, smtpUser, smtpPass, smtpHost string, smtpPort int) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/message/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewMessage))).Methods(http.MethodGet)

	r.Handle("/message", mw.Handler(CreateMessage(smtpUser, smtpPass, smtpHost, smtpPort))).Methods(http.MethodPost)

	r.Handle("/message/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchMessages))).Methods(http.MethodGet)
	r.Handle("/message/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchMessages))).Methods(http.MethodGet)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "comms.messages.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
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
