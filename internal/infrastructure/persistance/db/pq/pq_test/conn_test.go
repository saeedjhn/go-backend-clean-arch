package pq_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
)

func TestConn(t *testing.T) {
	cfg := pq.Config{
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

	pq := pq.New(cfg)
	if err := pq.ConnectTo(); err != nil {
		t.Fatal(err)
	}

	t.Log(pq.Conn())
}
