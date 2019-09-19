package db

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
)


// NOTE utils
// used for create fake rows
// ref: https://github.com/DATA-DOG/go-sqlmock/issues/42
func MockRowsToSqlRows(mockRows *sqlmock.Rows) *sql.Rows {
    db, mock, _ := sqlmock.New()
    mock.ExpectQuery("select").WillReturnRows(mockRows)
    rows, _ := db.Query("select")
    return rows
}

// End of utils

func TestNullTime(t *testing.T) {
	type fields struct {
		Time  NullTime
	}

	t.Run("Test db driver scan", func(t *testing.T) {
		rows := MockRowsToSqlRows(sqlmock.NewRows([]string{"fake_time"}).
			AddRow(time.Now()))
		rows.Next()
		f := fields{}
		wantErr := false

		if err := rows.Scan(&f.Time); (err != nil) != wantErr {
			t.Errorf("Scan() error = %v, wantErr %v", err, wantErr)
		}
	})

}
