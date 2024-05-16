package snowflake

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupSnowflake(t *testing.T) {
	// Test the SetupSnowflake function to ensure it returns a valid generator without error.
	// This test assumes that the environment allows for successful generator creation.
	generator, err := SetupSnowflake()
	assert.NoError(t, err)
	assert.NotNil(t, generator)

	// Optionally, test if the generator can produce an ID to verify its functionality.
	if generator != nil {
		id, err := generator.NextID()
		assert.NoError(t, err)
		assert.NotZero(t, id)
	}
}
