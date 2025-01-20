package models

// Relacion es la relacion entre usuarios
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioid"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuariorelacionid"`
}
