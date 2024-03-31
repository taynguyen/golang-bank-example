package middleware

import "github.com/gin-gonic/gin"

func LoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// check if user is logged in
		// if not, return 401
		// if yes, continue
	}
}
