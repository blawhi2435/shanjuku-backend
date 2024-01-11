package environment

import (
	"github.com/spf13/viper"
)

var Setting Config

type Config struct {
	Postgres Postgres `json:"postgres"`
	Gin      Gin      `json:"gin"`
	Auth     Auth     `json:"auth"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	TimeZone string `json:"timezone"`
}

type Gin struct {
	Address         string `json:"address"`
	Port            string `json:"port"`
	IsTLS           bool   `json:"is_tls"`
	CertificateFile string `json:"certificate_file"`
	KeyFile         string `json:"key_file"`
}

type Auth struct {
	JWTSecret      string `json:"jwt_secret"`
	PasswordPrefix string `json:"password_prefix"`
}

func SetConfig() {
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	Setting = c
}
