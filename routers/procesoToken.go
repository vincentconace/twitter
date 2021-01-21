package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/vincentconace/twitter/bd"
	"github.com/vincentconace/twitter/models"
)

//Email valor es el Email usado en todos los EndPoints
var Email string

//IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints
var IDUsuario string

//ProcesoToken proceso token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formatos de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrdo, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrdo == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrdo, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token ivalido")
	}
	return claims, false, string(""), err
}
