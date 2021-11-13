package models

/*RespuestaLogin tiene el token que returna el login*/
type RespuestaLogin struct {
	Token string `json:"token, omitempty"`
}
