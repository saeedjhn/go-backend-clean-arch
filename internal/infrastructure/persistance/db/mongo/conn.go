package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	config Config
	client *mongo.Client
	db     *mongo.Database
}

func New(config Config) *MongoDB {
	ctx := context.Background()

	connectionURI := fmt.Sprintf("mongodb://%s:%s/", config.Host, config.Port) // for local machine
	//connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Username, config.Password, config.Host, config.Port) // for server machine

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))

	if err != nil {
		log.Fatalf("can't open mongo db: %v", err)
	}

	//err = client.Ping(Ctx, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return &MongoDB{config: config, client: client}
}

func (m *MongoDB) Conn() {

}
