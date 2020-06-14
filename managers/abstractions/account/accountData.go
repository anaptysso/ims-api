package data

import (
	models "imsapi/models"

	"gopkg.in/mgo.v2/bson"
)

type AccountData interface {
	Connect()
	Disconnect()
	SelectAccounts(filter bson.D) []*models.Account
	InsertAccount(account *models.Account) bool
	UpdateAccounts(filter, update bson.D) (int64, int64)
}
