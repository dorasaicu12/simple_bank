package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/dorasaicu12/simplebank/db/mock"
	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamMatcher struct {
	arg db.CreateUserParams
	password string
}

func (e eqCreateUserParamMatcher) Matches(x interface{}) bool {
	arg,ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassWord(e.password,arg.HashedPassword)
	if err != nil {
		return false
	}
	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg,arg)
}
func (e eqCreateUserParamMatcher) String() string {
	return fmt.Sprintf("matches password")
}
func EqCreateUserParam(arg db.CreateUserParams, password string) eqCreateUserParamMatcher {
	return eqCreateUserParamMatcher{arg: arg,password: password}
}

func TestGetUserApi(t *testing.T) {
	User ,password:= randomUser(t)

	testCase := []struct {
		name          string
		body     gin.H
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "Ok",
			body: gin.H{
              "username":User.Username,
			  "password":password,
			  "fullname":User.FullName,
			  "email":User.Email,
			},
			buildStub: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username: User.Username,
					FullName: User.FullName,
					Email: User.Email,
					HashedPassword: User.HashedPassword,
				}
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParam(arg,password)).
					Times(1).
					Return(User, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				fmt.Println(recorder)
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, User)
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
			require.NoError(t,err)
			recorder := httptest.NewRecorder()
			//mashal body
			data,err := json.Marshal(tc.body)
			require.NoError(t,err)
			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			Server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t,recorder)
			//check response
		})
	}
}

func randomUser(t *testing.T) (user db.Users,password string) {
	passwordss := util.RandomString(6)
	hashedPasswod,err := util.HashePassword(passwordss)
	require.NoError(t,err)
	return db.Users{
		Username: util.RandomOwner(),
		HashedPassword: hashedPasswod,
	    FullName: util.RandomOwner(),
		Email: util.RandomEmail(),
	},passwordss
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, User db.Users) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.Users

	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	// require.Equal(t, User, gotUser)

}
