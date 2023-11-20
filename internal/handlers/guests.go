package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/asqit/guestbook/internal/models"
	"github.com/asqit/guestbook/internal/tools"
)

func CreateNewGuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	comment := r.PostFormValue("comment")

	var guest = models.Guest{
		Name:    name,
		Email:   email,
		Comment: comment,
		Time:    time.Now(),
		ID:      nil,
	}

	if err := models.InsertNewGuest(&guest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	htmlStr := fmt.Sprintf(`
	<li>
		<h3>NAME:<span>%s</span></h3>
		<h4>%s</h4>
		<a href="emailto:%s"><span>%s</span></a>
		<p>
			%s
		</p>
	</li>
	`, name, guest.Time.Format("2006-01-02 15:04:05"), email, email, comment)

	tmpl, err := template.New("t").Parse(htmlStr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func GetAllGuests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	guests, err := models.GetAllGuests()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	htmlStrs := make([]string, len(guests))

	for index, guest := range guests {
		htmlStrs[index] = fmt.Sprintf(`
		<li>
			<h3>NAME:<span>%s</span></h3>
			<h4>%s</h4>
			<a href="emailto:%s"><span>%s</span></a>
			<p>
				%s
			</p>
		</li>
		`, guest.Name, guest.Time.Format("2006-01-02 15:04:05"), guest.Email, guest.Email, guest.Comment)

		tmpl, err := template.New("t").Parse(htmlStrs[index])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}

}

// -------------------------------------- JSON API

func GetAllGuestsJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	guests, err := models.GetAllGuests()
	if err != nil {
		tools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	err = tools.WriteJSON(w, &guests)
	if err != nil {
		tools.WriteError(w, err, http.StatusInternalServerError)
		return
	}
}

func CreateNewGuestJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	guest, err := tools.ReadJSON[models.Guest](r.Body)
	if err != nil {
		tools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	err = models.InsertNewGuest(&guest)
	if err != nil {
		tools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	err = tools.WriteJSON(w, &guest)
	if err != nil {
		tools.WriteError(w, err, http.StatusInternalServerError)
		return
	}
}
