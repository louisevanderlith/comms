package main

import (
	"os"
	"path"
	"strconv"

	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/comms/routers"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()

	r.GET("/article/:key", article.View)

	articles := r.Group("/article")
	articles.POST("", article.Create)
	articles.PUT("/:key", article.Update)
	articles.DELETE("/:key", article.Delete)

	r.GET("/articles", article.Get)
	r.GET("/articles/:pagesize/*hash", article.Search)

	r.POST("/submit/Submit Contact", msgCtrl.Create)

	err := r.Run(":8085")

	if err != nil {
		panic(err)
	}
}

// func main() {
// 	keyPath := os.Getenv("KEYPATH")
// 	pubName := os.Getenv("PUBLICKEY")
// 	host := os.Getenv("HOST")
// 	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
// 	appName := os.Getenv("APPNAME")
// 	pubPath := path.Join(keyPath, pubName)

// 	// Register with router
// 	srv := bodies.NewService(appName, "", pubPath, host, httpport, servicetype.API)

// 	routr, err := do.GetServiceURL("", "Router.API", false)

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = srv.Register(routr)

// 	if err != nil {
// 		panic(err)
// 	}

// 	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, "none"))
// 	routers.Setup(poxy)
// 	poxy.EnableCORS(host)

// 	core.CreateContext()
// 	defer core.Shutdown()

// 	err = droxolite.Boot(poxy)

// 	if err != nil {
// 		panic(err)
// 	}
// }
