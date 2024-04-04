package accounts

import (
	"gin-boilerplate/internal/handlers"
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository/transactions"
	"gin-boilerplate/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	UUID      string                     `json:"uuid"`
	AccountID uint                       `json:"account_id"`
	Amount    float32                    `json:"amount"`
	StatusID  models.TransactionStatusID `json:"status_id"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt *time.Time                 `json:"updated_at,omitempty"`
}

type GetUserTransactionsUri struct {
	UserID uint `uri:"id" binding:"required"`
}

type GetUserTransactionsQuery struct {
	// UserID    uint  `uri:"user_id" binding:"required"`
	AccountID *uint `form:"account_id"`
}

type GetUserTransactionsResponse struct {
	Transactions []Transaction       `json:"transactions"`
	Pagination   handlers.Pagination `json:"pagination"`
}

func (h Handler) GetUserTransactions(c *gin.Context) {
	// TODO: Get user ID from token

	// Path params
	uriParam := GetUserTransactionsUri{}
	if err := c.ShouldBindUri(&uriParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_uri", "message": err.Error()})
		return
	}
	filter := transactions.GetTransactionsFilter{
		UserID: utils.Ptr(uriParam.UserID),
	}

	p := GetUserTransactionsQuery{}
	if err := c.BindQuery(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_query", "message": "Invalid query params"})
		return
	}
	filter.AccountID = p.AccountID

	// Pagination
	defaultLimit := 10
	maxTxLimit := 100
	var err error
	filter.Pagination, err = handlers.GetPaginationFromQuery(c, defaultLimit, maxTxLimit)
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
		res.Transactions = append(res.Transactions, txResponseFromModel(tx))
	}

	c.JSON(http.StatusOK, res)
}

func txResponseFromModel(tx models.Transaction) Transaction {
	return Transaction{
		UUID:      tx.UUID,
		AccountID: tx.AccountID,
		Amount:    tx.Amount,
		StatusID:  tx.StatusID,
		CreatedAt: tx.CreatedAt,
		UpdatedAt: tx.UpdatedAt,
	}
}
