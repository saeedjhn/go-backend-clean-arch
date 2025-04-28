package oterlcollector_test

import (
	"context"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMain_ExecuteQueryAndRecordMetric_Success(t *testing.T) {
	for i := range 10 {
		durExecQuery := time.Duration(math.Abs(rand.NormFloat64()*500)) * time.Microsecond
		// log.Println(durExecQuery.Microseconds())

		err := _otelClient.IntHistogram(
			context.Background(),
			"int_histogram_sql-query-execution-time",
			int(durExecQuery.Milliseconds()),
			"Execution time of SQL queries",
		)
		if err != nil {
			t.Errorf("Error query exection at iteration %d: %v", i, err)
		}

		// time.Sleep(5 * time.Second)
	}
	// time.Sleep(5 * time.Minute)
}

func TestMain_RandomNumberAndRecordMetricWithBucketBoundariesCustom_Success(t *testing.T) {
	for i := range 10 {
		rn := randomNumber()

		_otelClient.WithBucketBoundaries([]float64{5.0, 10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 100.0})
		if err := _otelClient.IntHistogram(
			context.Background(),
			"int_histogram_sql-query-execution-time-with-bucket-boundaries",
			// int(durExecQuery.Milliseconds()),
			rn,
			"Execution generate random number",
		); err != nil {
			t.Errorf("Error generate random number at iteration %d: %v", i, err)
		}

		// time.Sleep(5 * time.Second)
	}
	// time.Sleep(5 * time.Minute)
}
