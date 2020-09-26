package main

import (
	config "imsapi/config"
	docs "imsapi/docs"
	auxiliaries "imsapi/src/auxiliaries"
	data "imsapi/src/data"
	accountlogic "imsapi/src/logics/account"
	routes "imsapi/src/routes"

	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()
	configuration := config.GetConfiguration("./config/config.json")

	connectRoutes(e, configuration)
	setupSwagger(e, configuration)

	e.Logger.Fatal(e.Start(":" + configuration.Port))
}

func connectRoutes(e *echo.Echo, configuration *config.Configuration) {
	baseData := data.Data{
		Config: configuration,
	}
	routes.Account{
		Base: &routes.Base{R: e.Group("/account")},
		SignupLogic: &accountlogic.EmailSignupLogic{
			StringEssentials: &auxiliaries.StringEssentials{},
			SignupData: &data.AccountData{
				Data: &baseData,
			},
		},
	}.ConnectRoutes()
}

func setupSwagger(e *echo.Echo, configuration *config.Configuration) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	docs.SwaggerInfo.Title = "IMS API Documentation"
	docs.SwaggerInfo.Description = "Simple API descriptions for ims API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + configuration.Port
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
