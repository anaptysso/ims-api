package data

import (
	"context"

	models "imsapi/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountData struct {
	Data
}

func (this AccountData) InsertAccount(account models.Account) bool {
	collection := this.client.Database("imsdb").Collection("accounts")

	_, err := collection.InsertOne(context.TODO(), account)

	if err != nil {
		panic(err)
	}

	return true
}

func (this AccountData) SelectAccounts(filter primitive.M) []*models.Account {
	var results []*models.Account
	collection := this.client.Database("imsdb").Collection("accounts")

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		var account models.Account
		err := cur.Decode(&account)
		if err != nil {
			panic(err)
		}

		results = append(results, &account)
	}

	return results
}

func (this AccountData) UpdateAccounts(filter, update primitive.M) (int64, int64) {
	collection := this.client.Database("imsdb").Collection("accounts")
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return updateResult.MatchedCount, updateResult.ModifiedCount
}
