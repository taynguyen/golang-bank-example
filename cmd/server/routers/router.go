package routers

import (
	"gin-boilerplate/cmd/server/routers/middleware"
	"gin-boilerplate/internal/handlers/accounts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// func SetupRoute(repo repository.Registry) *gin.Engine {
// 	environment := viper.GetBool("DEBUG")
// 	if environment {
// 		gin.SetMode(gin.DebugMode)
// 	} else {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	allowedHosts := viper.GetString("ALLOWED_HOSTS")
// 	router := gin.New()
// 	router.SetTrustedProxies([]string{allowedHosts})
// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())
// 	router.Use(middleware.CORSMiddleware())

// 	RegisterRoutes(router) //routes register

// 	return router
// }

type Router struct {

	// ... handlers
	accountHandler accounts.Handler
}

func NewRouter(accountHandler accounts.Handler) *Router {
	return &Router{
		accountHandler: accountHandler,
	}
}

func (router *Router) RegisterRoutes(g *gin.Engine) {
	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	// router := gin.New()
	g.SetTrustedProxies([]string{allowedHosts})
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.CORSMiddleware())

	router.authenticated(g)
}

func (router *Router) public(g *gin.Engine) {
	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	g.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
}

func (router *Router) authenticated(g *gin.Engine) {
	g.GET("/accounts/:id/transactions", middleware.LoggedIn(), router.accountHandler.GetUserTransactions)
}
