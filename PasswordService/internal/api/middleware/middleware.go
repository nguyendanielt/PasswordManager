package middleware

import (
	"net/http"

	"passwordservice/pkg/service"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		authorized, userId, err := service.ValidateUser(token)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if authorized {
			// add user id to request header and pass the request to the next middleware (or final handler)
			r.Header.Set("userid", userId.String())
			next.ServeHTTP(w, r)
		} else {
			// write an error and stop the handler chain
			http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		}
	})
}
