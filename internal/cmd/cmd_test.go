package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nszilard/aws-lambda-go/internal/models"
)

func TestSuccess(t *testing.T) {
	pl := models.Event{
		Name:  "something",
		Count: 10,
	}

	_, err := Entrypoint(pl)

	assert.Nil(t, err)
}

func TestFailure(t *testing.T) {
	pl := models.Event{
		Count: 10,
	}

	_, err := Entrypoint(pl)

	assert.NotNil(t, err)
}
