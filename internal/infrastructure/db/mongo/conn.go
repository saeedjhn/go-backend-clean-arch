package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	cfg    Config
	client *mongo.Client
	db     *mongo.Database
}

func New(cfg Config) *MongoDB {
	ctx := context.Background()

	connectionURI := fmt.Sprintf("mongodb://%s:%s/", cfg.Host, cfg.Port) // for local machine
	//connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port) // for server machine

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))

	if err != nil {
		log.Fatalf("can't open mongo db: %v", err)
	}

	//err = client.Ping(Ctx, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return &MongoDB{cfg: cfg, client: client}
}

func (m *MongoDB) Conn() {

}
