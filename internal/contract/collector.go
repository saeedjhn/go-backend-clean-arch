package contract

type Collector interface {
	IncrementCounter(name string, value int, labels map[string]string)
	RecordGauge(name string, value float64, labels map[string]string)
	ObserveHistogram(name string, value float64, labels map[string]string)
}
