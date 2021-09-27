package infra

import (
	"database/sql"
	"fmt"
	"github.com/dedensmkn4/movie-backend/internal/config"
	"github.com/dedensmkn4/movie-backend/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type Databases struct {
	MySQL *sql.DB `name:"mysql"`
}

func NewDatabase(cfg *config.DatabaseCfg) Databases{
	return Databases{
		MySQL: openMySQL(cfg),
	}
}

func openMySQL(p *config.DatabaseCfg) *sql.DB {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&parseTime=true",
		p.DBUser, p.DBPass, p.DBHost, p.DBPort, p.DBName))
	if err != nil {
		logger.Log.Fatal("mysql", zap.String("reason", err.Error()))
	}

	if err = db.Ping(); err != nil {
		logger.Log.Fatal("mysql", zap.String("reason", err.Error()))
	}
	return db
}