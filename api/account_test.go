package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/petrostrak/simple-bank/db/mock"
	db "github.com/petrostrak/simple-bank/db/sqlc"
	"github.com/petrostrak/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountID     int64
		buildStub     func(*mockdb.MockStore)
		checkResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: account.ID,
			buildStub: func(ms *mockdb.MockStore) {
				// build stubs
				ms.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusOK, rr.Code)
				requireBodyMatchAccount(t, rr.Body, account)

			},
		},
		{
			name:      "NotFound",
			accountID: account.ID,
			buildStub: func(ms *mockdb.MockStore) {
				// build stubs
				ms.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusNotFound, rr.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: account.ID,
			buildStub: func(ms *mockdb.MockStore) {
				// build stubs
				ms.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusInternalServerError, rr.Code)
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			buildStub: func(ms *mockdb.MockStore) {
				// build stubs
				ms.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusBadRequest, rr.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
