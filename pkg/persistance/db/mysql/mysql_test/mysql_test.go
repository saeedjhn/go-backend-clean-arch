package mysql_test

import (
	"testing"

	mysql2 "github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

func TestConn(t *testing.T) {
	cfg := mysql2.Config{
		Host:            "localhost", // usage docker: hostName(container name)
		Port:            "5001",      // usage docker: postName(5342)
		Username:        "admin",
		Password:        "123456",
		Database:        "simorgh_db",
		SSLMode:         "disable",
		MaxIdleConns:    2,
		MaxOpenConns:    15,
		ConnMaxLiftTime: 5,
	}

	mysql := mysql2.New(cfg)
	if err := mysql.ConnectTo(); err != nil {
		t.Fatal(err)
	}

	t.Log(mysql.Conn())
}
