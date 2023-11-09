package main

import "github.com/danielhessell/api-go/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
