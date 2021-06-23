package key

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func toLower(pattern string) string {
	return strings.ToLower(pattern)
}

func dropHead(pattern string) string {
	return pattern[1:]
}

func TestNewConverterPipeline(t *testing.T) {
	expected := 0
	actual := len(NewConverterPipeline().Hooks)
	assert.Equal(t, expected, actual, "hooks must be empty")
}

func buildSampleConverterPipeline() *ConverterPipeline {
	return NewConverterPipeline().AddHook(toLower).AddHook(dropHead)
}

func TestConverterPipeline_AddHook(t *testing.T) {
	expected := 2
	actual := len(NewConverterPipeline().AddHook(toLower).AddHook(dropHead).Hooks)
	assert.Equal(t, expected, actual, "hooks are not registered")
}

func TestConverterPipeline_Convert(t *testing.T) {
	expected := "abc"
	actual := buildSampleConverterPipeline().Convert("AaBC")
	assert.Equal(t, expected, actual, "conversion is invalid")
}
