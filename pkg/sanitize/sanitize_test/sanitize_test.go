package sanitize_test

import (
	"github.com/saeedjhn/go-backend-clean-arch/pkg/sanitize"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:generate go test -v -race -count=1 ./...

func TestSanitize_SetPolicy_sanitizeStrictPolicy(t *testing.T) {
	t.Parallel()

	s := sanitize.New()
	s = s.SetPolicy(sanitize.StrictPolicy)
	assert.NotNil(t, s)
}

func TestSanitize_SetPolicy_UGCPolicy(t *testing.T) {
	t.Parallel()

	s := sanitize.New()
	s = s.SetPolicy(sanitize.UGCPolicy)
	assert.NotNil(t, s)
}

func TestSanitize_SetPolicy_StripTagsPolicy(t *testing.T) {
	t.Parallel()

	s := sanitize.New()
	s = s.SetPolicy(sanitize.StripTagsPolicy)
	assert.NotNil(t, s)
}

func TestSanitize_SetPolicy_DefaultPolicy(t *testing.T) {
	t.Parallel()

	s := sanitize.New()
	s = s.SetPolicy("unknown_policy")
	assert.NotNil(t, s)
}

func TestSanitize_String_HTMLSanitization(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := `<script>alert("XSS")</script>`
	output := s.String(input)
	assert.Equal(t, "", output)
}

func TestSanitize_String_JavaScriptRemoval(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := `javascript:alert("XSS")`
	output := s.String(input)
	assert.Equal(t, "", output)
}

func TestSanitize_String_JavaScriptRemovalLoop(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Simple JavaScript", `javascript:alert("XSS")`, ""},
		{"Mixed Case JavaScript", `JavaScript:alert("XSS")`, ""},
		{"Extra Spaces", `javascript    :alert("XSS")`, ""},
		{"Backslash t", "javascript\t:alert(\"XSS\")", ""},
		{"Backslash n", "javascript\n:alert(\"XSS\")", ""},
		{"No JavaScript", `safe text`, `safe text`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			output := s.String(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestSanitize_Any_StringInput(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := `<script>alert("XSS")</script>`
	output, err := s.Any(input)
	assert.NoError(t, err)
	assert.Equal(t, "", output)
}

func TestSanitize_Any_IntInput(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := 42
	output, err := s.Any(input)
	assert.NoError(t, err)
	assert.Equal(t, 42, output)
}

func TestSanitize_Any_SliceInput(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := []interface{}{`<script>alert("XSS")</script>`, 42}
	output, err := s.Any(input)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"", 42}, output)
}

func TestSanitize_Any_MapInput(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := map[string]interface{}{"key1": `<script>alert("XSS")</script>`, "key2": 42}
	output, err := s.Any(input)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"key1": "", "key2": 42}, output)
}

func TestSanitize_Any_StructInput(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	type TestStruct struct {
		Field1 string
		Field2 int
	}
	input := TestStruct{Field1: `<script>alert("XSS")</script>`, Field2: 42}
	output, err := s.Any(input)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"Field1": "", "Field2": 42}, output)
}

func TestSanitize_StructToMap_ValidStruct(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	type TestStruct struct {
		Field1 string
		Field2 int
	}
	input := TestStruct{Field1: `<script>alert("XSS")</script>`, Field2: 42}
	output, err := s.StructToMap(input)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"Field1": "", "Field2": 42}, output)
}

func TestSanitize_Struct_ValidPointer(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	type TestStruct struct {
		Field1 string
		Field2 int
	}
	input := &TestStruct{Field1: `<script>alert("XSS")</script>`, Field2: 42}
	err := s.Struct(input)
	assert.NoError(t, err)
	assert.Equal(t, &TestStruct{Field1: "", Field2: 42}, input)
}

func TestSanitize_Struct_InvalidPointer(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	type TestStruct struct {
		Field1 string
		Field2 int
	}
	input := TestStruct{Field1: `<script>alert("XSS")</script>`, Field2: 42}
	err := s.Struct(input)
	assert.Error(t, err)
}

func TestSanitize_Array_ValidArray(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := []interface{}{`<script>alert("XSS")</script>`, 42}
	output, err := s.Array(input)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"", 42}, output)
}

func TestSanitize_Map_ValidMap(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := map[string]interface{}{"key1": `<script>alert("XSS")</script>`, "key2": 42}
	output, err := s.Map(input)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"key1": "", "key2": 42}, output)
}

func TestSanitize_Recursively_UnsupportedType(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	input := make(chan int)
	_, err := s.Recursively(input)
	assert.Error(t, err)
}

func TestSanitize_Structure_ValidStruct(t *testing.T) {
	t.Parallel()

	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)
	type TestStruct struct {
		Field1 string
		Field2 int
	}
	input := TestStruct{Field1: `<script>alert("XSS")</script>`, Field2: 42}
	output, err := s.Structure(input)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"Field1": "", "Field2": 42}, output)
}

func TestSanitize_IsPointer_ValidPointer(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Field1 string
	}
	input := &TestStruct{Field1: "value"}
	assert.True(t, sanitize.IsPointer(input))
}

func TestSanitize_IsPointer_InvalidPointer(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Field1 string
	}
	input := TestStruct{Field1: "value"}
	assert.False(t, sanitize.IsPointer(input))
}
