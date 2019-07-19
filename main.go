package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"function/cmd"
	"function/internal"
)

//-------------------------------------------
// Internal Functions
//-------------------------------------------
func application(ctx context.Context, payload internal.Event) error {
	log.Println("CREATE SDK CLIENTS HERE...")

	err := cmd.MainLogic(payload)

	return handleError(err)
}

func handleError(err error) error {
	log.Println("MAIN ERROR HANDLING CODE GOES HERE...")
	return err
}

//-------------------------------------------
// Lambda entrypoint
//-------------------------------------------
func main() {
	lambda.Start(application)
}
