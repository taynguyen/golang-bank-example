package accounts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetUserTransactions(c *gin.Context) {
	// TODO: Get user ID from JWT token

	// Get user id from path parameter
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_USER_ID", "message": "Invalid user ID"})
		return
	}

	// Get transactions
	transactions, pagination, err := h.accountCtrl.GetTransactions(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_SERVER_ERROR", "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"pagination":   pagination,
	})
}
