package main

import (
	config "imsapi/config"
	data "imsapi/data"
	docs "imsapi/docs"
	managers "imsapi/managers"
	routes "imsapi/routes"

	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	configuration := config.GetConfiguration()
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	baseData := data.Data{
		Configuration: configuration,
	}

	routes.Account{
		Base: routes.Base{R: e.Group("/account")},
		SignUpManager: managers.AccountManager{
			Manager:     managers.Manager{},
			AccountData: &data.AccountData{Data: baseData},
		},
	}.New()

	docs.SwaggerInfo.Title = "IMS API Documentation"
	docs.SwaggerInfo.Description = "Simple API descriptions for ims API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + configuration.Port
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	e.Logger.Fatal(e.Start(":" + configuration.Port))
}
