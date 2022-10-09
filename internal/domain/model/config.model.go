package model

// Config es la configuracion del servidor
type Config struct {
	ApplicationPort string
	Hostdb          string
	Portdb          int
	Database        string
	Debug           bool
}
