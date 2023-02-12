package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

const (
	GET_PLAYLISTS = "/user/playlists"
)

func Register(router *gin.Engine, conn *pgx.Conn) {
	provideConn := provider(conn)

	router.GET(GET_PLAYLISTS, provideConn(getPlaylists))
}

func provider(conn *pgx.Conn) func(cb func(c *gin.Context, conn *pgx.Conn)) func(c *gin.Context) {
	return func(cb func(c *gin.Context, conn *pgx.Conn)) func(c *gin.Context) {
		return func(c *gin.Context) {
			cb(c, conn)
		}
	}
}

func getPlaylists(c *gin.Context, conn *pgx.Conn) {
	userId := c.Param("id")

	playlists, err := FindAllPlaylists(conn, userId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Request failed",
		})
		return
	}

	c.JSON(http.StatusOK, &playlists)
}
