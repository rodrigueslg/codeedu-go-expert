package main

import "github.com/rodrigueslg/codedu-goexpert/rest-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
