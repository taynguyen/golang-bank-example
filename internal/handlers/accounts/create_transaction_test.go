package accounts

import (
	"bytes"
	"gin-boilerplate/internal/models"
	"gin-boilerplate/test/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	accountsCtrlPkg "gin-boilerplate/internal/controllers/accounts"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateTransaction(t *testing.T) {
	type mockCreateTransactionResponseStruct struct {
		tx  *models.Transaction
		err error
	}

	tx1 := &models.Transaction{
		UUID:      "uuid-1",
		AccountID: 1,
		Amount:    100,
		TypeID:    1,
		StatusID:  1,
	}

	type args struct {
		queryUrl                      string
		queryBody                     string
		mockCreateTransactionResponse mockCreateTransactionResponseStruct
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
				queryUrl:  "/users/1/transactions",
				queryBody: `{"account_id":1,"amount":100,"type_id":1}`,
				mockCreateTransactionResponse: mockCreateTransactionResponseStruct{
					tx: &models.Transaction{
						UUID:      "uuid-1",
						AccountID: 1,
						Amount:    100,
						TypeID:    1,
						StatusID:  1,
					},
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `{"uuid":"uuid-1","account_id":1,"amount":100,"status_id":1,"created_at":"0001-01-01T00:00:00Z"}`,
		},
		{
			name: "SHOULD return 400 and valid error msg body WHEN missing required field type_id",
			args: args{
				queryUrl:  "/users/1/transactions",
				queryBody: `{"account_id":1,"amount":100}`,
				mockCreateTransactionResponse: mockCreateTransactionResponseStruct{
					tx:  tx1,
					err: nil,
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"code":"invalid_body","message":"Key: 'CreateTransactionBody.TypeID' Error:Field validation for 'TypeID' failed on the 'required' tag"}`,
		},
		{
			name: "SHOULD return 400 and valid error msg body WHEN invalid type_id is out range",
			args: args{
				queryUrl:  "/users/1/transactions",
				queryBody: `{"account_id":1,"amount":100,"type_id":100}`,
				mockCreateTransactionResponse: mockCreateTransactionResponseStruct{
					tx:  tx1,
					err: nil,
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"code":"invalid_body","message":"Key: 'CreateTransactionBody.TypeID' Error:Field validation for 'TypeID' failed on the 'oneof' tag"}`,
		},
		{
			name: "SHOULD return 400 and valid error msg body WHEN amount is not number",
			args: args{
				queryUrl:  "/users/1/transactions",
				queryBody: `{"account_id":1,"amount":"not-number","type_id":1}`,
				mockCreateTransactionResponse: mockCreateTransactionResponseStruct{
					tx:  tx1,
					err: nil,
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"code":"invalid_body","message":"json: cannot unmarshal string into Go struct field CreateTransactionBody.amount of type float32"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountCtrl := new(accountsCtrlPkg.MockIAccountController)
			mockAccountCtrl.On("CreateTransaction", mock.Anything, mock.Anything, mock.Anything).Return(
				tt.args.mockCreateTransactionResponse.tx,
				tt.args.mockCreateTransactionResponse.err)

			h := Handler{accountCtrl: mockAccountCtrl}

			router := utils.CreateTestRouter("POST", "/users/:id/transactions", h.CreateTransaction)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tt.args.queryUrl, bytes.NewBufferString(tt.args.queryBody))
			router.ServeHTTP(w, req)

			require.JSONEq(t, tt.wantBody, w.Body.String(), "json body")
			require.Equal(t, tt.wantStatusCode, w.Code, "status code")

		})
	}
}
