package key

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewKeyPanics() {
	_ = New("", buildSampleConverterPipeline(), buildSampleValidatorPipeline())
}

func buildSampleKey() *Key {
	return New("abc", buildSampleConverterPipeline(), buildSampleValidatorPipeline())
}

func buildSampleSelector() *Key {
	return New("abc/*", buildSampleConverterPipeline(), buildSampleValidatorPipeline())
}

func TestNew_Panic(t *testing.T) {
	assert.Panics(t, NewKeyPanics, "constructor must panic when validation error occurs")
}

func TestNew_NoPanic(t *testing.T) {
	key := Key("bc")
	expected := &key
	actual := buildSampleKey()

	assert.Equal(t, expected, actual, "constructor not working correctly")
}

func TestKey_IsRaw_True(t *testing.T) {
	assert.True(t, buildSampleKey().IsRaw())
}

func TestKey_IsRaw_False(t *testing.T) {
	assert.False(t, buildSampleSelector().IsRaw())
}

func TestKey_IsSelector_True(t *testing.T) {
	assert.True(t, buildSampleSelector().IsSelector())
}

func TestKey_IsSelector_False(t *testing.T) {
	assert.False(t, buildSampleKey().IsSelector())
}

func TestKey_Size(t *testing.T) {
	expected := 4
	actual := buildSampleSelector().Size()
	assert.Equal(t, expected, actual)
}
