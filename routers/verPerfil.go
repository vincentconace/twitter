package routers

import (
	"encoding/json"
	"net/http"

	"github.com/vincentconace/twitter/bd"
)

//VerPerfil - permite extraer los valores del perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debes enviar el parametro ID", http.StatusBadRequest)
		return
	}
	// si todo anduvo bien:
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un erro al intertar buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "encouding/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
