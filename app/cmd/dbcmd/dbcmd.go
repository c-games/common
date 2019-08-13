package dbcmd

import (
	"database/sql"
	"errors"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

func IsDbCmd(cmd string) bool {
	for _, c := range []string{"up", "down", "status"} {
		if cmd == c {
			return true
		}
	}
	return false
}

func CheckCmdArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a color argument")
	}
	if IsDbCmd(args[0]) {
		return nil
	}
	return fmt.Errorf("invalid db cmd specified: %s", args[0])
}

func Init(dbUrl string, action string, migrationList []*migrate.Migration) {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: migrationList,
	}

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		// TODO Handle errors!
		panic(err.Error())
	}
	//

	var n int
	//var err error
	switch action {
	case "up":
		n, err = migrate.Exec(db, "mysql", migrations, migrate.Up)
	case "down":
		n, err = migrate.Exec(db, "mysql", migrations, migrate.Down)
	case "status":
		// TODO find status
		//   migrate.GetMigrationRecords(db, "")
		panic("status 尚未實作")
	default:
		panic("action not found")
	}
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
