package middlew

import (
	"net/http"

	"github.com/vincentconace/twittor/bd"
)

//ChequeoBD es el middlew que me permite saber el estado de la base de datos
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConection() == 0 {
			http.Error(w, "Conecion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
