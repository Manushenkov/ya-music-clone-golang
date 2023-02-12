package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Playlist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FindAllPlaylists(conn *pgx.Conn, user_id string) ([]Playlist, error) {
	q := `
		SELECT playlist_id, playlist_name FROM playlists
	`

	playlists := []Playlist{}

	rows, err := conn.Query(context.Background(), q)

	if err != nil {
		return playlists, err
	}

	for rows.Next() {
		var playlistId int
		var playlistName string
		err := rows.Scan(&playlistId, &playlistName)
		if err != nil {
			fmt.Println(err)
			return playlists, err
		}
		playlists = append(playlists, Playlist{Id: playlistId, Name: playlistName})
	}

	return playlists, nil
}
