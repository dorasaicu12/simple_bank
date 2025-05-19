package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/dorasaicu12/simplebank/db/mock"
	"github.com/dorasaicu12/simplebank/token"
	"github.com/dorasaicu12/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func addAuthorization(t *testing.T,request *http.Request,tokenMake token.Maker,authoriztionType string,username string,duration time.Duration) {
	token,err := tokenMake.CreateToken(username,duration)
	if err != nil {
        require.NoError(t,err)
	}
	authorizationHeader := fmt.Sprintf("%s %s",authoriztionType,token)
	request.Header.Set(authorizationHeaderKey,authorizationHeader)
}
func TestAuthMiddleWare(t *testing.T) {
	testCase := []struct {
		name          string
		setUpAuth     func(t *testing.T, request *http.Request, tokenMkaer token.Maker)
		checkResponse func(t *testing.T, recoreder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setUpAuth:     func(t *testing.T, request *http.Request, tokenMkaer token.Maker){
                addAuthorization(t,request,tokenMkaer,authorizationTypeBearer,util.RandomOwner(),time.Minute)
			},
			checkResponse: func(t *testing.T, recoreder *httptest.ResponseRecorder){
               require.Equal(t,http.StatusOK,recoreder.Code)
			},
		},

		{
			name: "No Authorization",
			setUpAuth:     func(t *testing.T, request *http.Request, tokenMkaer token.Maker){

			},
			checkResponse: func(t *testing.T, recoreder *httptest.ResponseRecorder){
               require.Equal(t,http.StatusUnauthorized,recoreder.Code)
			},
		},
		{
			name: "Unsupported Authorization",
			setUpAuth:     func(t *testing.T, request *http.Request, tokenMkaer token.Maker){
				addAuthorization(t,request,tokenMkaer,"asdsadsad",util.RandomOwner(),time.Minute)
			},
			checkResponse: func(t *testing.T, recoreder *httptest.ResponseRecorder){
               require.Equal(t,http.StatusUnauthorized,recoreder.Code)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			clrt := gomock.NewController(t)
			defer clrt.Finish()
			store := mockdb.NewMockStore(clrt)

			config, err := util.LoadConfig("..")
			require.NoError(t, err)
			server, err := NewServer(config, store)
			require.NoError(t, err)
			authPath := "/auth"
			server.router.GET(
				authPath,
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setUpAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		},
		)
	}
}
