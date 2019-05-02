package dbconn

import (
	"github.com/Tavasiev/cws-backend/configs"
	pg "github.com/go-pg/pg"
)

var Conn *pg.DB

// Connect create connection
func Connect() error {

	Conn = pg.Connect(&pg.Options{
		Addr:     configs.Cfg.DataBase.Addr,
		User:     configs.Cfg.DataBase.User,
		Password: configs.Cfg.DataBase.Password,
		Database: configs.Cfg.DataBase.DB,
	})

	return nil
}


// CloseDbConnection closing connection for defer in main
func CloseDbConnection(db *pg.DB) {
	db.Close()
}
