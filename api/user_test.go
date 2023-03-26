package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/codeninjaug/authexample/db/mock"
	db "github.com/codeninjaug/authexample/db/sqlc"
	"github.com/codeninjaug/authexample/db/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetUserAPI(t *testing.T){
     user := randomUser()
	 ctrl := gomock.NewController(t)
	 defer ctrl.Finish()

	 mock_store := mockdb.NewMockStore(ctrl)
	 //build stubs
	 mock_store.EXPECT().
	 GetUser(gomock.Any(), gomock.Eq(user.ID)).
	 Times(1).
	 Return(user, nil)

	 //start test server and send request
	 server := NewServer(mock_store)
	 //records the response of the http response
	 recorder := httptest.NewRecorder()
	 //url for the endpoint
	 url := fmt.Sprintf("/users/%d", user.ID)
	 //making a variable request 
	 request, err := http.NewRequest(http.MethodGet, url, nil) 
	 require.NoError(t, err)
    //sends request with request and records the response in the recorder
	 server.router.ServeHTTP(recorder, request)
	 require.Equal(t, http.StatusOK, recorder.Code)
}  


func randomUser() db.User{
	return db.User{
		ID: int32(util.RandomInt(1,100)),
	    FirstName: util.RandomOwner(),
	    LastName:  util.RandomOwner(),
	    Phone:     util.RandomOwner(),
	    Email:     util.RandomOwner(),
	    Password:  util.RandomString(7),
	}
}