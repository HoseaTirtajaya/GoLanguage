package Router

import (
	"encoding/json"
	"net/http"
)

//Middleware
func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//Routing Handler
func handleBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseAPI{
		Status:  "Success",
		Message: "Hehehehehe",
	})
}
