package accounts

import (
	accountsCtrlPkg "gin-boilerplate/internal/controllers/accounts"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	accountCtrl accountsCtrlPkg.IAccountController
}

func NewHandler(accountCtrl accountsCtrlPkg.IAccountController) Handler {
	return Handler{
		accountCtrl: accountCtrl,
	}
}

func (h Handler) getUserIDFromToken(c *gin.Context) (int64, error) {
	// TODO: Implement this
	return 0, nil
}
