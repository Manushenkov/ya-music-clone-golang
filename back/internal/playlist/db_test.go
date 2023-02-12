package playlist

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock"
)

func TestDb_Create(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(context.Background())

	q := "INSERT INTO playlists (playlist_name) VALUES ($1)"

	mock.ExpectExec(q).WithArgs("example name")

	if err = Create(mock, &CreatePlaylistDTO{Name: "example name"}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
