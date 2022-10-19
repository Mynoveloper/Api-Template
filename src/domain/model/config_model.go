package model

// Config is the basic parametrization to configurate the program
type Config struct {
	ApplicationPort string
	Hostdb          string
	Portdb          int
	Database        string
	Debug           bool
}
