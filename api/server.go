package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/gregvroberts/cart-buddy/db/sqlc"
)

type Server struct {
	store  *db.Store   // allows interaction with database while processing API requests from clients
	router *gin.Engine // router to help send each API request to the correct handler
}

/*NewServer creates a new Server object and returns it
@param store *db.Store Database store for interacting with db
@return *Server Newly generated Server object with a Gin router and db store
*/
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default() // create a new router

	router.POST("/users", server.createUser)

	server.router = router // define the router with our newly created Gin router

	return server
}

type CreateUserRequest struct {
	UserFName   string         `json:"user_f_name" binding:"required"`
	UserLName   string         `json:"user_l_name" binding:"required"`
	UserEmail   string         `json:"user_email" binding:"required,email"`
	UserCity    string         `json:"user_city" binding:"required"`
	UserState   string         `json:"user_state" binding:"required"`
	UserPostal  string         `json:"user_postal" binding:"required"`
	UserCountry string         `json:"user_country" binding:"required"`
	UserAddr1   string         `json:"user_addr_1" binding:"required"`
	UserAddr2   sql.NullString `json:"user_addr_2"`
}

func (server *Server) createUser(ctx *gin.Context) {
	//TODO finish this (where you left off
}
