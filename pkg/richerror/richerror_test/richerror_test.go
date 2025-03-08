package richerror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/saeedjhn/go-domain-driven-design/pkg/richerror"
	"github.com/stretchr/testify/assert"
)

func TestRichError(t *testing.T) {
	t.Parallel()

	t.Run("New_WithOp_ReturnsInstanceWithOp", func(t *testing.T) {
		t.Parallel()

		op := richerror.Op("test_operation")
		err := richerror.New(op)

		assert.Equal(t, op, err.Op())
	})

	t.Run("WithMessage_SetMessage_ReturnsErrorWithMessage", func(t *testing.T) {
		t.Parallel()

		err := richerror.New("test_op").WithMessage("test message")

		assert.Equal(t, "test message", err.Message())
	})

	t.Run("WithKind_SetKind_ReturnsErrorWithKind", func(t *testing.T) {
		t.Parallel()

		err := richerror.New("test_op").WithKind(richerror.Kind(1))

		assert.Equal(t, richerror.Kind(1), err.Kind())
	})

	t.Run("WithErr_WrapError_ReturnsErrorWithWrappedError", func(t *testing.T) {
		t.Parallel()

		wrapped := errors.New("wrapped error")
		err := richerror.New("test_op").WithErr(wrapped)

		assert.Equal(t, wrapped, err.WrappedError())
	})

	t.Run("WithMeta_AddMetadata_ReturnsErrorWithMetadata", func(t *testing.T) {
		t.Parallel()

		err := richerror.New("test_op").WithMeta(map[string]interface{}{"key": "value"})

		meta := err.Meta()

		assert.Equal(t, "value", meta["key"])
	})

	t.Run("Meta_WithWrappedError_MergesMetadata", func(t *testing.T) {
		t.Parallel()

		wrapped := richerror.New("wrapped_op").WithMeta(map[string]interface{}{"key1": "value1"})
		err := richerror.New("test_op").WithMeta(map[string]interface{}{"key2": "value2"}).WithErr(wrapped)

		meta := err.Meta()

		assert.Equal(t, "value1", meta["key1"])
		assert.Equal(t, "value2", meta["key2"])
	})

	t.Run("Error_WithMessage_ReturnsMessage", func(t *testing.T) {
		t.Parallel()

		err := richerror.New("test_op").WithMessage("test message")

		assert.Equal(t, "test message", err.Error())
	})

	t.Run("Error_WithWrappedError_ReturnsWrappedErrorMessage", func(t *testing.T) {
		t.Parallel()

		wrapped := errors.New("wrapped error")
		err := richerror.New("test_op").WithErr(wrapped)

		assert.Equal(t, "wrapped error", err.Error())
	})

	t.Run("ToJSON_WithError_ReturnsSerializedJSON", func(t *testing.T) {
		t.Parallel()

		err := richerror.New("test_op").
			WithMessage("test message").
			WithKind(richerror.Kind(1))
		// WithMeta("key", "value")

		jsonString, jsonErr := err.ToJSON()

		require.NoError(t, jsonErr, "Expected no error during JSON serialization")
		assert.Contains(t, jsonString, "\"op\":\"test_op\"")
		assert.Contains(t, jsonString, "\"message\":\"test message\"")
		// assert.Contains(t, jsonString, "\"key\":\"value\"")
	})

	t.Run("Analysis_WithRichError_ReturnsExtractedRichError", func(t *testing.T) {
		t.Parallel()

		richErr := richerror.New("test_op").
			WithMessage("test message").
			WithMeta(map[string]interface{}{"foo": "bar1"})

		result := richerror.Analysis(richErr)

		assert.Equal(t, richErr.Op(), result.Op())
		assert.Equal(t, richErr.Message(), result.Message())
		assert.Equal(t, richErr.Meta(), result.Meta())
	})

	t.Run("Analysis_WithNonRichError_ReturnsEmptyRichError", func(t *testing.T) {
		t.Parallel()

		nonRichErr := errors.New("non-rich error")

		result := richerror.Analysis(nonRichErr)

		assert.Equal(t, richerror.Op(""), result.Op())
		assert.Equal(t, "", result.Message())
	})
}
