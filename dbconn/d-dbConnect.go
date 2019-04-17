package dbconn

import (
	"github.com/Tavasiev/cws-backend/configs"
	pg "github.com/go-pg/pg"
)

var conn *pg.DB

// Connect create connection
func Connect() error {

	conn = pg.Connect(&pg.Options{
		Addr:     configs.Cfg.DataBase.Addr,
		User:     configs.Cfg.DataBase.User,
		Password: configs.Cfg.DataBase.Password,
		Database: configs.Cfg.DataBase.DB,
	})

	return nil
}

// GetConnect get the connection
func GetConnect() *pg.DB {
	return conn
}

// CloseDbConnection closing connection for defer in main
func CloseDbConnection(db *pg.DB) {
	db.Close()
}
