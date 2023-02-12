package main

import (
	"context"
	"fmt"
	"music-clone/web-service-gin/internal/playlist"
	"music-clone/web-service-gin/internal/track"
	"music-clone/web-service-gin/internal/user"
	"music-clone/web-service-gin/pkg/client/postgresql"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {
	var dbUrl string
	if os.Getenv("DATABASE_URL") != "" {
		dbUrl = os.Getenv("DATABASE_URL")
	} else {
		dbUrl = "postgres://postgres:postgres@127.0.0.1:5432/postgres"
	}

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	postgresql.CreateTables(conn)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	track.Register(router, conn)
	playlist.Register(router, conn)
	user.Register(router, conn)

	fmt.Println("SERVICE STARTED")
	if os.Getenv("PORT") != "" {
		router.Run("localhost:8080")
	} else {
		router.Run("localhost:8080")
	}
}
