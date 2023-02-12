package playlist

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Close(context.Context) error
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func Create(conn PgxIface, playlist *CreatePlaylistDTO) error {
	q := "INSERT INTO playlists (playlist_name) VALUES ($1)"

	if _, err := conn.Exec(context.Background(), q, playlist.Name); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Delete(conn *pgx.Conn, playlistId string) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})

	tracks, err := FindTracks(tx.Conn(), playlistId)

	if err != nil {
		tx.Rollback(context.Background())
		fmt.Println(err, "err 1")
		return err
	}

	deletePlaylistQuery := `DELETE FROM playlists WHERE playlist_id = $1`

	if _, err := tx.Exec(context.Background(), deletePlaylistQuery, playlistId); err != nil {
		tx.Rollback(context.Background())
		fmt.Println(err, "err 2")

		return err
	}

	for _, track := range tracks {
		var isPresentInOtherPlaylist = 0
		q := `SELECT 1 FROM playlist_to_track WHERE track_id = $1`

		conn.QueryRow(context.Background(), q, track.ID).Scan(&isPresentInOtherPlaylist)

		if isPresentInOtherPlaylist == 0 {
			os.Remove("./tracks/" + strconv.Itoa(track.ID) + ".mp3")

			deleteTrackQuery := `DELETE FROM tracks WHERE track_id = $1`

			if _, err := tx.Exec(context.Background(), deleteTrackQuery, track.ID); err != nil {
				tx.Rollback(context.Background())
				fmt.Println(err, "err 4")

				return err
			}
		}

	}

	tx.Commit(context.Background())

	return nil
}

func DeleteTrack(conn *pgx.Conn, playlistId string, trackId string) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})

	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	deleteConnectionQuery := `DELETE FROM playlist_to_track WHERE playlist_id = $1 and track_id = $2`

	if _, err := tx.Exec(context.Background(), deleteConnectionQuery, playlistId, trackId); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	tx.Commit(context.Background())
	return nil
}

func FindOne(conn *pgx.Conn, id string) (Playlist, error) {
	q := `
		SELECT playlist_id, playlist_name FROM playlists WHERE playlist_id = $1
	`

	var playlist Playlist
	err := conn.QueryRow(context.Background(), q, id).Scan(&playlist.ID, &playlist.Name)

	if err != nil {
		return Playlist{}, err
	}

	return playlist, nil
}

type Track struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FindTracks(conn *pgx.Conn, playlist_id string) ([]Track, error) {
	q := `SELECT track_id, track_name FROM 
			playlist_to_track JOIN tracks
			USING(track_id)
			WHERE playlist_id = $1`

	tracks := []Track{}

	rows, err := conn.Query(context.Background(), q, playlist_id)

	if err != nil {
		return tracks, err
	}

	for rows.Next() {
		var trackId int
		var trackName string
		err := rows.Scan(&trackId, &trackName)
		if err != nil {
			fmt.Println(err)
			return tracks, err
		}
		tracks = append(tracks, Track{trackId, trackName})
	}

	return tracks, nil
}

func CreateTrack(conn *pgx.Conn, track *CreateTrackDTO, playlistId string) (int, error) {
	createTrackQuery := `
		INSERT INTO tracks 
		    (track_name) 
		VALUES 
		       ($1) 
		RETURNING track_id
	`
	var createdTrackId int

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return createdTrackId, err
	}

	if err := tx.QueryRow(context.Background(), createTrackQuery, track.Name).Scan(&createdTrackId); err != nil {
		tx.Rollback(context.Background())
		return createdTrackId, err
	}

	bindTrackToPlaylistQuery := `
		INSERT INTO playlist_to_track 
		    (playlist_id, track_id) 
		VALUES 
		       ($1, $2) 
	`

	if _, err := tx.Exec(context.Background(), bindTrackToPlaylistQuery, playlistId, createdTrackId); err != nil {
		tx.Rollback(context.Background())
		return createdTrackId, err
	}

	tx.Commit(context.Background())

	return createdTrackId, nil
}
