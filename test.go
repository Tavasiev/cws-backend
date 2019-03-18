package main

import (
    "./configs"
    "log"
)

func main() {
	config.InitConfigs("configs/config")
	log.Print(config.Cfg.Server.MainPort)
}