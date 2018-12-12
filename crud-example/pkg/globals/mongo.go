package globals

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func LoadClient() *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	if err != nil {
		//fail early
		panic(err)
	}

	return c
}

func CloseConnection() {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err := Client.Disconnect(ctx)

	if err != nil {
		panic(err)
	}
}
