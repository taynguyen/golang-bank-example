package accounts

import (
	"gin-boilerplate/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateTransactionUri struct {
	UserID uint `uri:"id" binding:"required"`
}

type CreateTransactionBody struct {
	AccountID uint                     `json:"account_id" binding:"required"`
	Amount    float32                  `json:"amount" binding:"required,gt=0"`
	TypeID    models.TransactionTypeID `json:"type_id" binding:"required,oneof=1 2"`
}

func (h Handler) CreateTransaction(c *gin.Context) {
	var uri CreateTransactionUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"code": "invalid_uri", "message": err.Error()})
		return
	}

	var body CreateTransactionBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"code": "invalid_body", "message": err.Error()})
		return
	}

	// Get account
	tx, err := h.accountCtrl.CreateTransaction(c.Request.Context(), uri.UserID, models.Transaction{
		AccountID: body.AccountID,
		Amount:    body.Amount,
		TypeID:    uint(body.TypeID),
	})
	if err != nil {
		c.JSON(500, gin.H{"code": "internal_error", "message": err.Error()})
		return
	}

	// Convert to response
	// TODO: convert to response
	c.JSON(200, tx)
}
