package environment

import (
	"fmt"

	"github.com/spf13/viper"
)

var Setting Config

type Config struct {
	Postgres Postgres `json:"postgres"`
	Gin      Gin      `json:"gin"`
	Logger	 Logger   `json:"logger"`
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
	IsTLS           bool   `json:"isTls"`
	CertificateFile string `json:"certificateFile"`
	KeyFile         string `json:"keyFile"`
}

type Logger struct {
	Level    string `json:"level"`
	Path     string `json:"path"`
	FileName string `json:"fileName"`
}

type Auth struct {
	JWTSecret      string `json:"jwtSecret"`
	PasswordPrefix string `json:"passwordPrefix"`
}

func SetConfig() {
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	Setting = c
	fmt.Printf("Setting: %+v\n", Setting)
}
