package track

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

const (
	LISTEN_URL = "/track/:id"
)

func Register(router *gin.Engine, conn *pgx.Conn) {
	provideConn := provider(conn)

	router.GET(LISTEN_URL, provideConn(listenTrack))
}

func provider(conn *pgx.Conn) func(cb func(c *gin.Context, conn *pgx.Conn)) func(c *gin.Context) {
	return func(cb func(c *gin.Context, conn *pgx.Conn)) func(c *gin.Context) {
		return func(c *gin.Context) {
			cb(c, conn)
		}
	}
}

func listenTrack(c *gin.Context, conn *pgx.Conn) {
	trackId := c.Param("id")

	c.FileAttachment("./tracks/"+trackId+".mp3", trackId+".mp3")
}
