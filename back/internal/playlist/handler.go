package playlist

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

const (
	GET_TRACKS       = "/playlist/:id/tracks"
	CREATE_PLAYLIST  = "/playlist/create/:name"
	DELETE_PLAYLIST  = "/playlist/delete/:id"
	UPLOAD_TRACK_URL = "/playlist/:id/tracks/upload"
	DELETE_TRACK     = "/playlist/:playlistId/tracks/delete/:trackId"
)

func Register(router *gin.Engine, conn *pgx.Conn) {
	provideConn := provider(conn)

	router.GET(GET_TRACKS, provideConn(getTracks))
	router.POST(CREATE_PLAYLIST, provideConn(createPlaylist))
	router.DELETE(DELETE_PLAYLIST, provideConn(deletePlaylist))
	router.POST(UPLOAD_TRACK_URL, provideConn(uploadTrack))
	router.DELETE(DELETE_TRACK, provideConn(deleteTrack))
}

func provider(conn *pgx.Conn) func(cb func(c *gin.Context, conn *pgx.Conn)) func(c *gin.Context) {
	return func(cb func(c *gin.Context, conn *pgx.Conn)) func(c *gin.Context) {
		return func(c *gin.Context) {
			cb(c, conn)
		}
	}
}

func getTracks(c *gin.Context, conn *pgx.Conn) {
	playlistId := c.Param("id")

	playlist, err := FindOne(conn, playlistId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Request failed",
		})
		return
	}

	tracks, err := FindTracks(conn, playlistId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Request failed",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"tracks": tracks, "name": playlist.Name})
}

func createPlaylist(c *gin.Context, conn *pgx.Conn) {
	playlistName := c.Param("name")

	Create(conn, &CreatePlaylistDTO{Name: playlistName})
}

func deletePlaylist(c *gin.Context, conn *pgx.Conn) {
	playlistId := c.Param("id")

	if err := Delete(conn, playlistId); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Request failed",
		})
		return
	}

	c.String(http.StatusOK, "success")
}

func uploadTrack(c *gin.Context, conn *pgx.Conn) {
	file, err := c.FormFile("file")
	playlistId := c.Param("id")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	trackId, err := CreateTrack(conn, &CreateTrackDTO{Name: file.Filename}, playlistId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to save file",
		})
		return
	}

	var dst = "/Users/golubets/music-clone/back/tracks/" + trackId + ".mp3"
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func deleteTrack(c *gin.Context, conn *pgx.Conn) {
	playlistId := c.Param("playlistId")
	trackId := c.Param("trackId")

	if err := DeleteTrack(conn, playlistId, trackId); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete file",
		})

	}

	os.Remove("./tracks/" + trackId + ".mp3")
}
