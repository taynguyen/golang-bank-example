package accounts

import (
	"gin-boilerplate/internal/handlers"
	"gin-boilerplate/internal/repository/transactions"
	"gin-boilerplate/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	UUID      string     `json:"uuid"`
	AccountID uint       `json:"account_id"`
	Amount    float32    `json:"amount"`
	StatusID  uint       `json:"status_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type GetUserTransactionsQuery struct {
	AccountID *uint `form:"account_id"`
}

type GetUserTransactionsResponse struct {
	Transactions []Transaction       `json:"transactions"`
	Pagination   handlers.Pagination `json:"pagination"`
}

func (h Handler) GetUserTransactions(c *gin.Context) {
	// TODO: Get user ID from token

	// Path params
	userIDStr := c.Param("id")
	val, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_user_id", "message": "Invalid user ID"})
		return
	}
	userID := uint(val)
	filter := transactions.GetTransactionsFilter{
		UserID: utils.Ptr(userID),
	}

	p := GetUserTransactionsQuery{}
	if err := c.BindQuery(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_query", "message": "Invalid query params"})
		return
	}
	filter.AccountID = p.AccountID

	// Pagination
	defaultLimit := 10
	filter.Pagination, err = handlers.GetPaginationFromQuery(c, defaultLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_pagination", "message": err.Error()})
		return
	}

	// Get transactions
	transactions, pagination, err := h.accountCtrl.GetTransactions(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_SERVER_ERROR", "message": "Internal server error"})
		return
	}

	// Convert to model
	res := GetUserTransactionsResponse{
		Transactions: make([]Transaction, 0, len(transactions)),
		Pagination: handlers.Pagination{
			Total:  pagination.Total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}
	for _, tx := range transactions {
		res.Transactions = append(res.Transactions, Transaction{
			UUID:      tx.UUID,
			AccountID: tx.AccountID,
			Amount:    tx.Amount,
			StatusID:  tx.StatusID,
			CreatedAt: tx.CreatedAt,
			UpdatedAt: tx.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, res)
}
