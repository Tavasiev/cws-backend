package main

import (
    "./cwsConfig"
    "log"
)

func main() {
	log.Print(configs.GetConfig("Server"))
}