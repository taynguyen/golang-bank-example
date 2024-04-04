package utils

import "github.com/gin-gonic/gin"

func CreateTestRouter(method string, path string, handlerFunc gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	if method == "POST" {
		r.POST(path, handlerFunc)
		return r
	}

	r.GET(path, handlerFunc)
	return r
}
