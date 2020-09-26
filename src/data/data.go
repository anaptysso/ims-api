package data

import (
	"context"
	config "imsapi/config"
	"reflect"

	mongobson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	gobson "gopkg.in/mgo.v2/bson"
)

//Data provides the way to interact with mongo database
type Data struct {
	client   *mongo.Client
	Database *mongo.Database
	Config   *config.Configuration
}

// Connect function connects with the mongo database
func (dt *Data) Connect() {
	var err error
	dt.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(dt.Config.Database.Path+":"+dt.Config.Database.Port))

	if err != nil {
		panic(err)
	}

	dt.Database = dt.client.Database(dt.Config.Database.Name)
}

// Disconnect function connects with the mongo database
func (dt *Data) Disconnect() {
	err := dt.client.Disconnect(context.TODO())

	if err != nil {
		panic(err)
	}
}

// CovertToMongobsonD function converts gobson.D to mongobson.D
func (dt *Data) CovertToMongobsonD(object gobson.D) mongobson.D {
	var convertedObject mongobson.D

	objectMap := object.Map()
	for key, value := range objectMap {
		if reflect.TypeOf(value) == reflect.TypeOf(object) {
			convertedObject = append(convertedObject, mongobson.E{key, dt.CovertToMongobsonD(value.(gobson.D))})
		} else {
			convertedObject = append(convertedObject, mongobson.E{key, value})
		}
	}

	return convertedObject
}
