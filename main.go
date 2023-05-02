package main

import (
	"context"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/nszilard/log"

	"github.com/nszilard/aws-lambda-go/internal/cmd"
	"github.com/nszilard/aws-lambda-go/internal/models"
)

// -------------------------------------------
// Internal Functions
// -------------------------------------------
func application(ctx context.Context, payload models.Event) (*models.Response, error) {
	res, err := cmd.Entrypoint(payload)
	if err != nil {
		return nil, handleError(err)
	}

	return &models.Response{Value: res}, nil
}

func handleError(err error) error {
	log.Errorf("lambda invokation failed: %v", err)
	return err
}

// -------------------------------------------
// Lambda entrypoint
// -------------------------------------------
func main() {
	runtime.Start(application)
}
