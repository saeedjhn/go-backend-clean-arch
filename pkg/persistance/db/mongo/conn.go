package mongo

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	config Config
	client *mongo.Client
	db     *mongo.Database
}

func New(config Config) *Mongo {
	ctx := context.Background()

	// connectionURI := fmt.Sprintf("mongodb://%s:%s/", config.Host, config.Port) // for local machine
	connectionURI := net.JoinHostPort(config.Host, config.Port)
	// connectionURI :=
	// fmt.Sprintf(
	// "mongodb://%s:%s@%s:%s", config.Username, config.Password, config.Host, config.Port,
	// ) // for server machine

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))

	db := client.Database(config.Database)

	if err != nil {
		log.Fatalf("can't open mongo db: %v", err)
	}

	// err = client.Ping(Ctx, nil) if err != nil { log.Fatal(err) }

	return &Mongo{config: config, client: client, db: db}
}

func (m *Mongo) Client() *mongo.Client {
	return m.client
}

func (m *Mongo) Database() *mongo.Database {
	return m.db
}
