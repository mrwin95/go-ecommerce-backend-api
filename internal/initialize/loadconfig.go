package initialize

import (
	"fmt"

	"github.com/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadingConfig() {
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

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
