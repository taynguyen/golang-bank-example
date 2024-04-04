package accounts

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	accountsCtrlPkg "gin-boilerplate/internal/controllers/accounts"
	"gin-boilerplate/internal/models"
)

func TestHandler_GetUserTransactions(t *testing.T) {
	type mockGetTransactionResponseStruct struct {
		Transactions []models.Transaction
		Pagination   models.Pagination
		err          error
	}

	type args struct {
		queryUrl                     string
		mockGetTransactionsResponses mockGetTransactionResponseStruct
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "SHOULD return 200 and valid json body WHEN everything is ok",
			args: args{
				queryUrl: "/users/1/transactions",
				mockGetTransactionsResponses: mockGetTransactionResponseStruct{
					Transactions: []models.Transaction{
						{
							UUID:      "uuid-1",
							AccountID: 1,
							Amount:    100,
							TypeID:    1,
							StatusID:  1,
						},
					},
					Pagination: models.Pagination{Limit: 10, Offset: 0, Total: 1},
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `{"transactions":[{"uuid":"uuid-1","account_id":1,"amount":100,"status_id":1,"created_at":"0001-01-01T00:00:00Z"}],"pagination":{"limit":10,"offset":0,"total":1}}`,
		},
		{
			name: "SHOULD return 500 and valid json body WHEN error occurred",
			args: args{
				queryUrl: "/users/1/transactions",
				mockGetTransactionsResponses: mockGetTransactionResponseStruct{
					Transactions: nil,
					Pagination:   models.Pagination{},
					err:          errors.New("some thing bad"),
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       `{"code":"INTERNAL_SERVER_ERROR", "message":"Internal server error"}`,
		},
		{
			name: "SHOULD return 400 and valid json body WHEN invalid user id",
			args: args{
				queryUrl: "/users/invalid/transactions",
				mockGetTransactionsResponses: mockGetTransactionResponseStruct{
					Transactions: nil,
					Pagination:   models.Pagination{},
					err:          nil,
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"code":"invalid_uri_param", "message":"Invalid URI params"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountCtrl := new(accountsCtrlPkg.MockIAccountController)
			mockAccountCtrl.On("GetTransactions", mock.Anything, mock.Anything).Return(
				tt.args.mockGetTransactionsResponses.Transactions,
				tt.args.mockGetTransactionsResponses.Pagination,
				tt.args.mockGetTransactionsResponses.err,
			)

			h := Handler{
				accountCtrl: mockAccountCtrl,
			}

			router := createTestRouter("/users/:id/transactions", h.GetUserTransactions)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.args.queryUrl, nil)
			router.ServeHTTP(w, req)

			require.Equal(t, tt.wantStatusCode, w.Code, "status code")
			require.JSONEq(t, tt.wantBody, w.Body.String(), "json body")
		})
	}
}

func createTestRouter(path string, handlerFunc gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.GET(path, handlerFunc)
	return r
}
