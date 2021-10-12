package db

import (
	"context"
	"database/sql"
	"github.com/gregvroberts/cart-buddy/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		UserFName:   util.RandomName(),
		UserLName:   util.RandomName(),
		UserEmail:   util.RandomEmail(),
		UserCity:    util.RandomString(15),
		UserState:   util.RandomString(15),
		UserPostal:  util.RandomPostal(),
		UserCountry: util.RandomString(20),
		UserAddr1:   util.RandomAddress(),
	}

	argEmailVerified := sql.NullBool{Bool: false, Valid: true}

	user, err := testQueries.CreateUser(context.Background(), arg)

	// Imported from testify. If requirements are not met, the test will automatically fail.
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserFName, user.UserFName)
	require.Equal(t, arg.UserLName, user.UserLName)
	require.Equal(t, arg.UserEmail, user.UserEmail)
	require.Equal(t, arg.UserCity, user.UserCity)
	require.Equal(t, arg.UserState, user.UserState)
	require.Equal(t, arg.UserPostal, user.UserPostal)
	require.Equal(t, arg.UserCountry, user.UserCountry)
	require.Equal(t, arg.UserAddr1, user.UserAddr1)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	require.Equal(t, user.UserEmailVerified, argEmailVerified)

}
