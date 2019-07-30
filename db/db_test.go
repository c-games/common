package db

import (
	"strings"
	"testing"

	//"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	//"gitlab.3ag.xyz/backend/common/testutil"
)

// TODO fix import cycle
/*func TestDBAdapter_Exec(t *testing.T) {
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
}*/



func TestGenDropTable(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				struct{
					name string
				}{},

			},
			want: "DROP TABLE cg_;",
		},
		{
			args: args {
				func() interface{} {
					type TestStruct struct {}
					return TestStruct{}
				}(),
			},
			want: "DROP TABLE cg_test_struct;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenDropTable(tt.args.s); got != tt.want {
				t.Errorf("GenDropTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenCreateTable(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				struct{
				}{},

			},
			want: `CREATE TABLE cg_(
) ENGINE=INNODB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`,
		},
		{
			args: args {
				func() interface{} {
					type TestStruct struct {}
					return TestStruct{}
				}(),
			},
						want: `CREATE TABLE cg_test_struct (
) ENGINE=INNODB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`,
		},
		{
			args: args {
				func() interface{} {
					type TestStruct struct {
						Name string `sql:"varchar(64)"`
						Id   string `sql:"bigint"`
					}
					return TestStruct{}
				}(),
			},
			want:
			`CREATE TABLE cg_test_struct (
name varchar(64),
id bigint) ENGINE=INNODB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`,
		},
		{
			args: args {
				func() interface{} {
					type TestStruct struct {
						Name string `sql:"varchar(64)"`
						Id   string `sql:"bigint" pk:""`
					}
					return TestStruct{}
				}(),
			},
			want:
			`CREATE TABLE cg_test_struct (
name varchar(64),
id bigint,
PRIMARY KEY (id)) ENGINE=INNODB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenCreateTable(tt.args.s);  strings.TrimRight(got, "\n") !=  strings.TrimRight(tt.want, "\n") {
				t.Errorf("GenCreateTable():\n%v\nwant:\n%v", got, tt.want)
			}
		})
	}
}

func TestGenIndexTable(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				struct{
				}{},

			},
			want: ``,
		},
		{
			args: args {
				func() interface{} {
					type TestStruct struct {}
					return TestStruct{}
				}(),
			},
			want: ``,
		},
		{
			args: args {
				func() interface{} {
					type TestStruct struct {
						Id string `index:"a"`
						Name string `index:"a"`
						Age int `index:"b"`
					}
					return TestStruct{}
				}(),
			},
			want: `CREATE INDEX a ON cg_test_struct (id,name);
CREATE INDEX b ON cg_test_struct (age);`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenCreateIndex(tt.args.s);  strings.TrimRight(got, "\n") !=  strings.TrimRight(tt.want, "\n") {
				t.Errorf("GenCreateTable():\n%v\nwant:\n%v", got, tt.want)
			}
		})
	}
}