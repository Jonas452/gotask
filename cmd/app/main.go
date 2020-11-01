package main

import (
	"fmt"
	"gotask/internal/app/adapter"
	"gotask/internal/app/migration"

	"github.com/spf13/viper"
)

func main() {
	setupEnv()
	migration.Migrate()

	r := adapter.Router()
	r.Run(":8080")
}

func setupEnv() {
	viper.SetConfigName("database")
	viper.AddConfigPath("/go/src/gotask/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
