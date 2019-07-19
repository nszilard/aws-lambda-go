package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"function/internal"
)

//-------------------------------------------
// Variables
//-------------------------------------------
var (
	eventFailure = internal.Event{}
	eventSuccess = internal.Event{
		Name:   "something",
		Region: "eu-west-1",
	}
)

//-------------------------------------------
// Mocks
//-------------------------------------------

//-------------------------------------------
// Tests
//-------------------------------------------
func TestMainLogicError(t *testing.T) {
	defer resetMocks()
	err := MainLogic(eventFailure)

	assert.NotNil(t, err)
}

func TestFailure(t *testing.T) {
	err := MainLogic(eventSuccess)

	assert.Nil(t, err)
}

//-------------------------------------------
// Helper functions
//-------------------------------------------
func resetMocks() {}
