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

	mock.ExpectExec(`SELECT (.+) FROM orders WHERE 1`).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectClose()

	_ = adp.Exec("SELECT id FROM orders WHERE 1")

	err = adp.Close()
	testutil.TestFailIfErr(t, err, "Can't not close db connection")
}

func TestDBAdapter_QueryRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	testutil.TestFailIfErr(t, err, "Can't not open sqlmock")

	adp := ConnectByDB(db)

	mock.ExpectQuery("SELECT (.+) FROM orders WHERE game_id = (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"test", "123"}))

	mock.ExpectClose()

	rows := adp.QueryRow("SELECT id FROM orders WHERE game_id = ?", 1234)

	// TODO How To check rows?
	t.Log(rows)

	err = adp.Close()
	testutil.TestFailIfErr(t, err, "Can't not close db connection")
}

func TestDBAdapter_PrepareQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	testutil.TestFailIfErr(t, err, "Can't not create sql mock")

	adp := ConnectByDB(db)

	mock.ExpectPrepare("SELECT (.+) FROM orders WHERE game_id = (.+)")

	mock.ExpectQuery("SELECT (.+) FROM orders WHERE game_id = (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"test", "123"}))

	mock.ExpectClose()

	rows := adp.PrepareQuery("SELECT id FROM orders WHERE game_id = ?", 1234)

	// TODO How To check rows?
	t.Log(rows)

	err = adp.Close()
	testutil.TestFailIfErr(t, err, "Can't not close db connection")
}