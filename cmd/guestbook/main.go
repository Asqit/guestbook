package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/asqit/guestbook/internal/db"
	"github.com/asqit/guestbook/internal/handlers"
	"github.com/asqit/guestbook/internal/models"
	"github.com/asqit/guestbook/internal/tools"
)

func main() {
	host := tools.GetEnvVar("DB_HOST")
	port, err := strconv.Atoi(tools.GetEnvVar("DB_PORT"))
	tools.PanicIfErr(err)
	pwd := tools.GetEnvVar("DB_PWD")
	usr := tools.GetEnvVar("DB_USR")

	db.InitConnection(fmt.Sprintf("postgres://%s:%s@%s:%d/guestbook?sslmode=disable", usr, pwd, host, port))
	tools.PanicIfErr(models.CreateGuestTable())

	fmt.Println("starting the guestbook service\n\rlistening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.InitRouting()))
}
