package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/dorasaicu12/simplebank/db/mock"
	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/token"
	"github.com/dorasaicu12/simplebank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountApi(t *testing.T) {
	user,_ := randomUser(t)
	account := randomAccount(user.Username)

	

	testCase := []struct {
		name          string
		accountId     int64
		setUpAuth     func(t *testing.T, request *http.Request, tokenMkaer token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "Ok",
			accountId: account.ID,
			setUpAuth:     func(t *testing.T, request *http.Request, tokenMkaer token.Maker){
                addAuthorization(t,request,tokenMkaer,authorizationTypeBearer,user.Username,time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
	}
    //TODO:add more cases

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name,func(t *testing.T){
			clrt := gomock.NewController(t)
			defer clrt.Finish()
			store := mockdb.NewMockStore(clrt)
			tc.buildStub(store)
			//start new server
			config,err := util.LoadConfig("..")
			require.NoError(t, err)
			Server,err := NewServer(config,store)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/account/%d", account.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			tc.setUpAuth(t,request,Server.tokenMaker)
			Server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t,recorder)
			//check response
		})
	}
}

func randomAccount(owner string) db.Accounts {
	return db.Accounts{
		ID:       util.RandomInt(1, 1000),
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Accounts) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Accounts

	err = json.Unmarshal(data, &gotAccount)

	require.NoError(t, err)
	require.Equal(t, account, gotAccount)

}
