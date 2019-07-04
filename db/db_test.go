package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gitlab.3ag.xyz/backend/common/testutil"
	"testing"
)

func TestDBAdapter_Exec(t *testing.T) {
	db, mock, err := sqlmock.New()
	testutil.TestFailIfErr(t, err, "Can't not open sqlmock")

	adp := ConnectByDB(db)

	mock.ExpectExec(`SELECT (.+) FROM order WHERE 1`).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectClose()

	_ = adp.Exec("SELECT id FROM order WHERE 1")

	err = adp.Close()
	testutil.TestFailIfErr(t, err, "Can't not close db connection")
}

func TestDBAdapter_QueryRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	testutil.TestFailIfErr(t, err, "Can't not open sqlmock")

	adp := ConnectByDB(db)

	game_id := "5cf53a25-14d3-4a3a-87fe-cf473f74419b"
	mock.ExpectQuery("SELECT (.+) FROM order WHERE game_id = (.+)").
		WithArgs(game_id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1234))

	mock.ExpectClose()

	row := adp.QueryRow("SELECT id FROM order WHERE game_id = ?", game_id)

	var (
		id int
	)
	err = row.Scan(&id)
	testutil.TestFailIfErr(t, err, "occur fail in err")

	if id != 1234 {

		t.Errorf("Can't scan value id: %v", id)
	}

	err = adp.Close()
	testutil.TestFailIfErr(t, err, "Can't not close db connection")
}

func TestDBAdapter_PrepareQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	testutil.TestFailIfErr(t, err, "Can't not create sql mock")

	adp := ConnectByDB(db)

	mock.ExpectPrepare("SELECT (.+) FROM order WHERE game_id = (.+)")

	mock.ExpectQuery("SELECT (.+) FROM order WHERE game_id = (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"test", "123"}))

	mock.ExpectClose()

	rows := adp.PrepareQuery("SELECT id FROM order WHERE game_id = ?", 1234)

	// TODO How To check rows?
	t.Log(rows)

	err = adp.Close()
	testutil.TestFailIfErr(t, err, "Can't not close db connection")
}