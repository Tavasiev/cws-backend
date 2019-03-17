package configs

import (
    "github.com/BurntSushi/toml"
    "log"
)

//GetConfig возвращает заданное поле конфигурации
func GetConfig(Name string)map[string]interface{} {
    var config map[string]interface{}
    _, err := toml.DecodeFile("config.toml", &config); 
    if err != nil {
        log.Printf("Error decode config %s",err)
    }
    data, isOk := config[Name].(map[string]interface {})
    if isOk == false {log.Print("Error check .toml file")}
    return data
}