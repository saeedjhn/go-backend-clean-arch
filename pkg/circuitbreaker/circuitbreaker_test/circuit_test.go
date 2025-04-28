package circuitbreaker_test

import (
	"errors"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/circuitbreaker"
)

var counter = 1 //nolint:gochecknoglobals // nothing

func TestRun(t *testing.T) {
	manager := circuitbreaker.New[string](5*time.Second, 3) // timeout=5s, threshold=3 failures

	for i := range 10 {
		t.Logf("Request #%d:\n", i+1)

		result, err := manager.Execute("apiBreaker", unstableFunction)
		if err != nil {
			t.Log("❌ Error:", err)
		} else {
			t.Log("✅ Result:", result)
		}

		state, _ := manager.State("apiBreaker")
		t.Log("📊 State:", state.CurrentState, "| Consecutive Errors:", state.ConsecutiveErrors)
		t.Log(strings.Repeat("-", 40))

		time.Sleep(1 * time.Second)
	}
}

func unstableFunction() (string, error) {
	log.Println("unstableFunction invoked!", counter)
	counter++

	if rand.Intn(10) < 7 {
		return "", errors.New("🔥 simulated error occurred")
	}

	return "✅ successful response", nil
}
