package database

import (
	"strconv"

	"github.com/blawhi2435/shanjuku-backend/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_postgresDriver "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func Migrate(postgreService *service.PostgresService, version string) {

	db, err := postgreService.DB.DB()
	if err != nil {
		logrus.Panic(err.Error())
	}

	driver, err := _postgresDriver.WithInstance(db, &_postgresDriver.Config{})
	if err != nil {
		logrus.Panic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migration",
		"postgres", driver)

	if version == "latest" {
		err = m.Up()
		if err != nil {
			logrus.Panic(err.Error())
		}
	} else {
		v, err := strconv.ParseUint(version, 10, 64)
		if err != nil {
			logrus.Panic(err.Error())
		}
		err = m.Migrate(uint(v))
		if err != nil {
			logrus.Panic(err.Error())
		}
	}

	currentVersion, _, err := m.Version()
	logrus.Infof("change version to %d\n: ", currentVersion)
}
