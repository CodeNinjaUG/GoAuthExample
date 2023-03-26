package db

import (
	"context"
	"testing"
	"time"
    "github.com/codeninjaug/authexample/db/util"
	"github.com/stretchr/testify/require"
)


func CreateRandomUser(t *testing.T) User{
     arg := CreateUserParams{
		FirstName : "Andrew",
	    LastName:"Segayi",
	    Phone: "+256704277679",    
	    Email :"segayiandrew@gmail.com",
	    Password : "8bdc7axyzex!@",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.FirstName,user.FirstName)
	require.Equal(t, arg.LastName,user.LastName)
	require.Equal(t, arg.Phone,user.Phone)
	require.Equal(t, arg.Email,user.Email)
	require.Equal(t, arg.Password,user.Password)
	require.NotZero(t,user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T){
	 CreateRandomUser(t)
}

func TestGetUser(t *testing.T){
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(),user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)	
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T){
	user1 := CreateRandomUser(t)
	arg   := UpdateUserParams{
	         ID: user1.ID,
	         FirstName: util.RandomOwner(),
	         LastName:  util.RandomOwner(),
	         Phone:     util.RandomOwner(),
	         Email:     util.RandomOwner(),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.FirstName, user2.FirstName)	
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, arg.Phone, user2.Phone)
	require.Equal(t, arg.Email, user2.Email)
}