package config

import "os"

type (
	DatabaseCfg struct {
		DBName string
		DBUser string
		DBPass string
		DBHost   string
		DBPort   string
	}
)

func NewDatabaseConfig() *DatabaseCfg {
	dbCfg := &DatabaseCfg{}

	dbCfg.DBName = os.Getenv("DB_NAME")
	dbCfg.DBUser = os.Getenv("DB_USER")
	dbCfg.DBPass = os.Getenv("DB_PASS")
	dbCfg.DBHost = os.Getenv("DB_HOST")
	dbCfg.DBPort = os.Getenv("DB_PORT")

	return dbCfg
}