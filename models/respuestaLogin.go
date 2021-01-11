package models

//RespuestaLogin tiene el token que se devuelve con el ligin
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
