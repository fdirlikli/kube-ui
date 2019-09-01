package service

import (
	"context"
	"github.com/fdirlikli/kube-ui/data"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/appengine/log"
)

func CreateApplication(app data.Application) {
	saved, err := data.ApplicationCollection().InsertOne(context.TODO(), app)

	if err != nil {
		panic(err.Error())
	}

	log.Infof(nil, "%s %s", "New Application Saved With Id:", saved.InsertedID)
}

func GetAllApplications() []data.Application {
	var response []data.Application
	cur, err := data.ApplicationCollection().Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		panic(err.Error())
	}

	for cur.Next(context.TODO()) {
		var elem data.Application
		err := cur.Decode(&elem)
		if err != nil {

		}
		response = append(response, elem)
	}

	return response
}
