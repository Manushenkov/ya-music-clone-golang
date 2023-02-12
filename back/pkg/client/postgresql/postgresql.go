package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CreateTables(conn *pgx.Conn) {
	conn.Exec(context.Background(),
		`CREATE TABLE tracks(
			track_id SERIAL primary key NOT NULL,
			track_name varchar(50)
		)`)

	conn.Exec(context.Background(),
		`CREATE TABLE playlists(
	  		playlist_id SERIAL primary key NOT NULL,
			playlist_name varchar(50)
		)`)

	conn.Exec(context.Background(),
		`CREATE TABLE playlist_to_track(
			playlist_id INT REFERENCES playlists ON DELETE CASCADE,
			track_id INT REFERENCES tracks ON DELETE CASCADE,
			PRIMARY KEY(playlist_id, track_id)
		  )`)
}
