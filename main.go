package main

import (
	config "imsapi/config"
	data "imsapi/data"
	managers "imsapi/managers"
	routes "imsapi/routes"

	echo "github.com/labstack/echo"
)

func main() {
	configuration := config.GetConfiguration()
	e := echo.New()

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

	e.Logger.Fatal(e.Start(":8000"))
}
