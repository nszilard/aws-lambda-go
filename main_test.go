package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"function/internal"
)

//-------------------------------------------
// Variables
//-------------------------------------------
var (
	ctx          context.Context
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
func TestSuccess(t *testing.T) {
	defer resetMocks()
	err := application(ctx, eventSuccess)

	assert.Nil(t, err)
}

func TestFailure(t *testing.T) {
	defer resetMocks()
	err := application(ctx, eventFailure)

	assert.NotNil(t, err)
}

//-------------------------------------------
// Helper functions
//-------------------------------------------
func resetMocks() {}
