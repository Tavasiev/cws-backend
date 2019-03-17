package main

import (
    "github.com/BurntSushi/toml"
    "log"
)

type Server struct{
	Port string
}

type Configuration struct {
	Server Server
}

//GetConfig возвращает заданное поле конфигурации
func GetConfig()interface{} {
    var config Configuration
    _, err := toml.DecodeFile("config.toml", &config); 
    if err != nil {
        log.Printf("Error decode config %s",err)
    }
    return config.Server.Port
}
