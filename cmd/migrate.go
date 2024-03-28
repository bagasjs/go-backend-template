package cmd

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var migrationFileNames = []string {
    "./migrations/sqlite3/create_internal_users_table.sql",
    "./migrations/sqlite3/create_internal_tables_table.sql",
    "./migrations/sqlite3/create_internal_fields_table.sql",
}

func onMigrate(cmd *cobra.Command, args []string) {
    db, err := sql.Open("sqlite3", "./res/db.sqlite3")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    for _, fileName := range migrationFileNames {
        log.Println("Migrating ", fileName)
        data, err :=  os.ReadFile(fileName)
        if err != nil {
            log.Fatal(err)
        }
        _, err = db.Exec(string(data))
        if err != nil {
            log.Fatal(err)
        }
    }
}

var migrateCmd = &cobra.Command {
    Use: "migrate",
    Short: "Database migration",
    Long: "Running every migration step",
    Run: onMigrate,
}
