package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouting() *mux.Router {
	mux := mux.NewRouter()

	// create a new guest
	mux.HandleFunc("/api/guest", CreateNewGuest).Methods("POST")
	mux.HandleFunc("/api/guest", GetAllGuests).Methods("GET")
	mux.HandleFunc("/api/json/guest", CreateNewGuestJSON).Methods("POST")
	mux.HandleFunc("/api/json/guest", GetAllGuestsJSON).Methods("GET")

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static/"))))

	// Serve index page on all unhandled routes
	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/index.html")
	})

	return mux
}
