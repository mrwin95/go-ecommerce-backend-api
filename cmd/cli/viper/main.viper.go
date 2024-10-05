package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read file config

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config file %w \n", err))
	}

	fmt.Println("Server port::", viper.GetInt("server.port"))
}
