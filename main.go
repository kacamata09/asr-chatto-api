package main

import (
	httpRoutes "asrChat/http"
	"fmt"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("env.yaml")
	if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

	fmt.Println("ans")
	e := echo.New()
	httpRoutes.StartHttp(e)

	port := viper.GetInt("server.port")
    e.Start(fmt.Sprintf(":%d", port))
	
}