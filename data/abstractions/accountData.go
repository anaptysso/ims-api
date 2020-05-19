package data

import (
	models "imsapi/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountData interface {
	Connect()
	Disconnect()
	SelectAccounts(filter primitive.M) []*models.Account
	InsertAccount(account models.Account) bool
	UpdateAccounts(filter, update primitive.M) (int64, int64)
}
