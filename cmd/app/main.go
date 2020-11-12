package main

import (
	"fmt"
	"gotask/internal/app/adapter"
	"gotask/internal/app/migration"
	"log"

	"github.com/spf13/viper"
)

func main() {
	setupEnv()
	migration.Migrate()

	r := adapter.Router()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("The server failed to start: %v", err)
	}
}

func setupEnv() {
	viper.SetConfigName("database")
	viper.AddConfigPath("/go/src/gotask/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
