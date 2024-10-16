package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func SetSession(w http.ResponseWriter, r *http.Request, userID int) {
	session, _ := store.Get(r, "session")
	session.Values["user_id"] = userID
	session.Save(r, w)
}

func GetSessionUserID(r *http.Request) (int, bool) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["user_id"].(int)
	return userID, ok
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	// Get the session from the request
	session, _ := store.Get(r, "session")

	// Set the session's MaxAge to -1 to delete the session
	session.Options.MaxAge = -1

	// Save the session with the updated options
	session.Save(r, w)
}

func SetAdminSession(w http.ResponseWriter, r *http.Request, adminID int) {
	session, _ := store.Get(r, "admin_session")
	session.Values["admin_id"] = adminID
	session.Save(r, w)
}

func GetAdminSessionID(r *http.Request) (int, bool) {
	session, _ := store.Get(r, "admin_session")
	adminID, ok := session.Values["admin_id"].(int)
	return adminID, ok
}
func ClearAdminSession(w http.ResponseWriter, r *http.Request) {
	// Get the session from the request
	session, _ := store.Get(r, "admin_session")

	// Set the session's MaxAge to -1 to delete the session
	session.Options.MaxAge = -1

	// Save the session with the updated options
	session.Save(r, w)
}
