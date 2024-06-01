package bcrypt

import (
	"log"
	"testing"
)

func TestBcrypt(t *testing.T) {
	log.Println(Generate("pass", 32))
}
