package main

import (
	"fmt"

	"github.com/aminasadiam/Shop/cmd"
	"github.com/aminasadiam/Shop/config"
)

func main() {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Println("Config loaded successfully")

	err = cmd.Router(config)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
