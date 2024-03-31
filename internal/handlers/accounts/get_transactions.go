package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetUserTransactions(c *gin.Context) {
	// Get user ID from JWT token
	userID, err := h.getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Get transactions
	transactions, pagination, err := h.accountCtrl.GetTransactions(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"pagination":   pagination,
	})
}
