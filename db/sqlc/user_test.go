package db

import (
	"context"
	"database/sql"
	"github.com/gregvroberts/cart-buddy/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/*createRandomUser Generates a random user in the database
@param t *testing.T The test object type
@return User The randomly generated user
*/
func createRandomUser(t *testing.T) User {
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

	return user // return the new user generated
}

/*TestQueries_CreateUser Test for the CreateUser function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

/*TestQueries_GetUser Test for the GetUser function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_GetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.UserID)

	require.NoError(t, err)   // Verify no error
	require.NotZero(t, user2) // Verify we got an object back

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.UserFName, user2.UserFName)
	require.Equal(t, user1.UserLName, user2.UserLName)
	require.Equal(t, user1.UserEmail, user2.UserEmail)
	require.Equal(t, user1.UserEmailVerified, user2.UserEmailVerified)
	require.Equal(t, user1.UserCity, user2.UserCity)
	require.Equal(t, user1.UserState, user2.UserState)
	require.Equal(t, user1.UserPostal, user2.UserPostal)
	require.Equal(t, user1.UserCountry, user2.UserCountry)
	require.Equal(t, user1.UserAddr1, user2.UserAddr1)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

/*TestQueries_UpdateUser Tests the UpdateUser function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_UpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		UserID:            user1.UserID,
		UserFName:         util.RandomName(),
		UserLName:         util.RandomName(),
		UserEmail:         util.RandomEmail(),
		UserEmailVerified: sql.NullBool{Bool: true, Valid: true},
		UserCity:          util.RandomString(15),
		UserState:         util.RandomString(15),
		UserPostal:        util.RandomPostal(),
		UserCountry:       util.RandomString(20),
		UserAddr1:         util.RandomAddress(),
		UserAddr2:         sql.NullString{String: util.RandomString(10), Valid: true},
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.NotEqual(t, user1.UserFName, user2.UserFName)
	require.NotEqual(t, user1.UserLName, user2.UserLName)
	require.NotEqual(t, user1.UserEmail, user2.UserEmail)
	require.NotEqual(t, user1.UserEmailVerified, user2.UserEmailVerified)
	require.NotEqual(t, user1.UserCity, user2.UserCity)
	require.NotEqual(t, user1.UserState, user2.UserState)
	require.NotEqual(t, user1.UserPostal, user2.UserPostal)
	require.NotEqual(t, user1.UserCountry, user2.UserCountry)
	require.NotEqual(t, user1.UserAddr1, user2.UserAddr1)
	require.NotEqual(t, user1.UserAddr2, user2.UserAddr2)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.NotEqual(t, user1.UpdatedAt, user2.UpdatedAt)
}

/*TestQueries_DeleteUser Test the DeleteUser function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_DeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.UserID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.UserID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)

}
