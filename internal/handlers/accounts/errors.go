package accounts

import (
	"gin-boilerplate/infra/logger"
	"gin-boilerplate/internal/controllers/accounts"

	"github.com/gin-gonic/gin"
)

func convertAndResponseError(c *gin.Context, err error) {
	switch err {
	case accounts.ErrAccountNotFound:
		c.JSON(404, gin.H{"code": "account_not_found", "message": "account not found"})
	case accounts.ErrInsufficientBalance:
		c.JSON(400, gin.H{"code": "insufficient_funds", "message": "insufficient funds"})
	default:
		logger.GetLogger().Errorf("internal error: %v", err)
		c.JSON(500, gin.H{"code": "internal_error", "message": "internal error"})
	}
}
