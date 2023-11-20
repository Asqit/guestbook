package db

import (
	"database/sql"
	"fmt"

	"github.com/asqit/guestbook/internal/tools"
	_ "github.com/lib/pq"
)

var DB *sql.DB = nil

func InitConnection(connectionStr string) {
	var err error = nil
	DB, err = sql.Open("postgres", connectionStr)
	tools.PanicIfErr(err)

	// check connection
	tools.PanicIfErr(DB.Ping())

	fmt.Println("successfully connected to the database")
}

func CloseConnection() {
	tools.PanicIfErr(DB.Close())
}
