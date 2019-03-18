package config

import (
    "github.com/spf13/viper"
    "log"
)

type MainConfigs struct {
	Server struct{
		MainPort string
	}

	DataBase struct{
		Addr string
		User string
		Password string
		DB string
	}
}

var Cfg MainConfigs
	
//InitConfigs Initializes the main programm settings
func InitConfigs(name string) {
	initDefaults()
	viper.SetConfigName(name)
	viper.AddConfigPath("/configs/")
	viper.AddConfigPath(".")
	
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Read configuration error. %s",err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Printf("Unmarshaling configuration error. %s",err)
	}
}

//initDefaults Initializes default values
func initDefaults() {
	viper.SetDefault("Server.MainPort",":1323")
	viper.SetDefault("DataBase.Addr","localhost:5432")
	viper.SetDefault("DataBase.User","testu1")
	viper.SetDefault("DataBase.Password","testpass1")
	viper.SetDefault("DataBase.DB","vscale_db")	
}