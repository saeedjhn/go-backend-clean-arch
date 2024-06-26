package pq

import (
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	cfg := Config{
		Host:            "localhost", // usage docker: hostName(container name)
		Port:            "5001",      // usage docker: postName(5342)
		Username:        "admin",
		Password:        "123456",
		Database:        "backend_db",
		SSLMode:         "disable",
		MaxIdleConns:    2,
		MaxOpenConns:    15,
		ConnMaxLiftTime: 5,
	}

	postgresqlConn := New(cfg)

	log.Fatal(postgresqlConn.Conn())
}
