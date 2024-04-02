package routers

import (
	"gin-boilerplate/cmd/server/routers/middleware"
	"gin-boilerplate/internal/handlers/accounts"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

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
	g.SetTrustedProxies([]string{allowedHosts})
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.CORSMiddleware())

	router.public(g)
	router.authenticated(g)

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
}

func (router *Router) public(g *gin.Engine) {
	g.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
}

func (router *Router) authenticated(g *gin.Engine) {
	v1 := g.Group("/api/v1")
	v1.GET("/users/:id/transactions", middleware.LoggedIn(), router.accountHandler.GetUserTransactions)
}
