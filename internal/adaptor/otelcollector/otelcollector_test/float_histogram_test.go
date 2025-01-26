package oterlcollector_test

import (
	"context"
	"testing"
)

func TestMain_RandomAPIResponseAndRecordMetric_Success(t *testing.T) {
	for i := range 10 {
		rAPI := randomAPIResponse()

		if err := _otelClient.FloatHistogram(
			context.Background(),
			"float_histogram_api-response-duration",
			rAPI,
			"Execution time of api response duration",
		); err != nil {
			t.Errorf("Error api resonse duration at iteration %d: %v", i, err)
		}

		// time.Sleep(5 * time.Second)
	}
	// time.Sleep(5 * time.Minute)
}
