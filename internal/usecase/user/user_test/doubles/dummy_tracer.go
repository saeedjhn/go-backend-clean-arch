package doubles

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
)

type DummySpan struct {
}

func (d DummySpan) End(_ ...bool) {
}

func (d DummySpan) SetAttributes(_ map[string]interface{}) {
}

func (d DummySpan) SetAttribute(_, _ string) {
}

func (d DummySpan) AddEvent(_ string, _ ...map[string]interface{}) {
}

func (d DummySpan) SetName(_ string) {
}

func (d DummySpan) SetStatus(_ uint32, _ string) {
}

func (d DummySpan) RecordError(_ error, _ ...map[string]interface{}) {
}

type DummyTracer struct {
}

func NewDummyTracer() *DummyTracer {
	return &DummyTracer{}
}

func (d DummyTracer) Configure() error {
	return nil
}

func (d DummyTracer) Span(_ context.Context, _ string) (context.Context, contract.Span) {
	return context.Background(), &DummySpan{}
}

func (d DummyTracer) Shutdown(_ context.Context) error {
	return nil
}
