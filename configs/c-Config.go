package configs

import (
	"github.com/BurntSushi/toml"
)

//TomlConfig main структура конфига.
//Содержит в себе другие структуры, типа Database и тд.
type TomlConfig struct {
	MainPort string   `toml:"MainPort"`
	DB       Database `toml:"database"`
}

//Database Структура DB, отвечает за коннект к бд.
type Database struct {
	Addr     string
	User     string
	Password string
	Database string
}

//MakeConfig Возвращает заполненую структуру TomlConfig со вло-
//женными в нее структурами Database и тд.
func MakeConfig() TomlConfig {
	var conf TomlConfig
	if _, err := toml.DecodeFile("configs/config.toml", &conf); err != nil {
		panic(err)
	}

	return conf
}
