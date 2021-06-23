package key

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func notEmpty(pattern string) error {
	if pattern == "" {
		return errors.New("pattern must not be empty")
	}
	return nil
}

func noSpacing(pattern string) error {
	for idx, char := range pattern {
		if char == ' ' {
			return fmt.Errorf(
				`pattern: "%s"  must not include spaces, but had spacing on index: %d`,
				pattern,
				idx,
			)
		}
	}
	return nil
}

func buildSampleValidatorPipeline() *ValidatorPipeline {
	return NewValidatorPipeline().AddHook(noSpacing).AddHook(notEmpty)
}

func TestNewValidatorPipeline(t *testing.T) {
	expected := 0
	actual := len(NewValidatorPipeline().Hooks)
	assert.Equal(t, expected, actual, "validator hooks must be empty when initialized")
}

func TestValidatorPipeline_AddHook(t *testing.T) {
	expected := 2
	actual := len(NewValidatorPipeline().AddHook(notEmpty).AddHook(noSpacing).Hooks)
	assert.Equal(t, expected, actual, "hooks are not added correctly")
}

func TestValidatorPipeline_Validate_CaseOneError(t *testing.T) {
	pipeline := buildSampleValidatorPipeline()
	assert.Error(
		t,
		pipeline.Validate("a k"),
		"validation did not catch the error",
	)
}

func TestValidatorPipeline_Validate_CaseTwoError(t *testing.T) {
	actual := buildSampleValidatorPipeline()
	assert.Error(t, actual.Validate(""), "validation did not catch the error")
}

func TestValidatorPipeline_Validate_NoError(t *testing.T) {
	actual := buildSampleValidatorPipeline()
	assert.NoError(t, actual.Validate("aaa"), "validation should have caught any errors")
}
